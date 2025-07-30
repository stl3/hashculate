package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// HashAlgorithm represents the supported hash algorithms
type HashAlgorithm string

const (
	MD5    HashAlgorithm = "md5"
	SHA1   HashAlgorithm = "sha1"
	SHA256 HashAlgorithm = "sha256"
	SHA512 HashAlgorithm = "sha512"
)

// HashResult contains the result of a hash calculation
type HashResult struct {
	Algorithm   HashAlgorithm
	Hash        string
	Filename    string
	FileSize    int64
	ChunkSize   int64
	Description string
}

// HashCalculator handles file hash calculations
type HashCalculator struct {
	ChunkSize int64 // Default 4MB like the HTML version
}

// NewHashCalculator creates a new hash calculator with default chunk size
func NewHashCalculator() *HashCalculator {
	return &HashCalculator{
		ChunkSize: 4 * 1024 * 1024, // 4MB chunks
	}
}

// createHasher creates the appropriate hash.Hash based on algorithm
func (hc *HashCalculator) createHasher(algorithm HashAlgorithm) (hash.Hash, error) {
	switch algorithm {
	case MD5:
		return md5.New(), nil
	case SHA1:
		return sha1.New(), nil
	case SHA256:
		return sha256.New(), nil
	case SHA512:
		return sha512.New(), nil
	default:
		return nil, fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}
}

// formatBytes formats bytes in a human-readable format (similar to HTML version)
func formatBytes(bytes int64) string {
	if bytes == 0 {
		return "0 bytes"
	}

	// For small files, show bytes; for larger files, show KB
	if bytes < 1024 {
		return fmt.Sprintf("%d bytes", bytes)
	}

	// Convert to KB for consistency with HTML version
	kb := float64(bytes) / 1024
	return fmt.Sprintf("%.1f kb (kilobytes)", kb)
}

// getAlgorithmName returns the display name for the algorithm
func getAlgorithmName(algorithm HashAlgorithm) string {
	switch algorithm {
	case MD5:
		return "MD5"
	case SHA1:
		return "SHA-1"
	case SHA256:
		return "SHA-256"
	case SHA512:
		return "SHA-512"
	default:
		return "Unknown"
	}
}

// CalculateFileHash calculates the hash of a file using the specified algorithm
func (hc *HashCalculator) CalculateFileHash(filePath string, algorithm HashAlgorithm, progressCallback func(float64)) (*HashResult, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	// Create hasher
	hasher, err := hc.createHasher(algorithm)
	if err != nil {
		return nil, err
	}

	// Process file in chunks
	buffer := make([]byte, hc.ChunkSize)
	var totalRead int64 = 0
	fileSize := fileInfo.Size()

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		if bytesRead == 0 {
			break
		}

		// Update hash with chunk
		hasher.Write(buffer[:bytesRead])
		totalRead += int64(bytesRead)

		// Report progress
		if progressCallback != nil && fileSize > 0 {
			progress := float64(totalRead) / float64(fileSize)
			progressCallback(progress)
		}
	}

	// Finalize hash
	hashBytes := hasher.Sum(nil)
	hashHex := fmt.Sprintf("%x", hashBytes)

	// Create description similar to HTML version
	filename := filepath.Base(filePath)
	algorithmName := getAlgorithmName(algorithm)
	description := fmt.Sprintf("\"%s\", with size of %s, and file hash using the hashing algorithm %s has the value : %s.",
		filename, formatBytes(fileSize), algorithmName, hashHex)

	return &HashResult{
		Algorithm:   algorithm,
		Hash:        hashHex,
		Filename:    filename,
		FileSize:    fileSize,
		ChunkSize:   hc.ChunkSize,
		Description: description,
	}, nil
}

// String returns a string representation of the hash result
func (hr *HashResult) String() string {
	return fmt.Sprintf("File: %s\nAlgorithm: %s\nHash: %s\nSize: %s\n",
		hr.Filename, getAlgorithmName(hr.Algorithm), hr.Hash, formatBytes(hr.FileSize))
}

// parseAlgorithm parses algorithm string and returns HashAlgorithm
func parseAlgorithm(alg string) (HashAlgorithm, error) {
	switch strings.ToLower(alg) {
	case "md5":
		return MD5, nil
	case "sha1", "sha-1":
		return SHA1, nil
	case "sha256", "sha-256":
		return SHA256, nil
	case "sha512", "sha-512":
		return SHA512, nil
	default:
		return "", fmt.Errorf("unsupported algorithm: %s. Supported: md5, sha1, sha256, sha512", alg)
	}
}

// printUsage prints usage information
func printUsage() {
	fmt.Println("Hashculate - File Hash Calculator")
	fmt.Println("Usage: hashculate [options] <file>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -algorithm, -a  Hash algorithm (md5, sha1, sha256, sha512) [default: md5]")
	fmt.Println("  -chunk-size, -c Chunk size in MB for processing large files [default: 4]")
	fmt.Println("  -progress, -p   Show progress during calculation [default: true]")
	fmt.Println("  -help, -h       Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  hashculate myfile.txt")
	fmt.Println("  hashculate -algorithm sha256 myfile.txt")
	fmt.Println("  hashculate -a sha512 -c 8 largefile.bin")
}

// progressBar displays a simple progress bar
func progressBar(progress float64) {
	barWidth := 50
	filled := int(progress * float64(barWidth))
	bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth-filled)
	percentage := int(progress * 100)
	fmt.Printf("\rProgress: [%s] %d%%", bar, percentage)
	if progress >= 1.0 {
		fmt.Println()
	}
}

func main() {
	// Define command line flags
	var (
		algorithm     = flag.String("algorithm", "md5", "Hash algorithm (md5, sha1, sha256, sha512)")
		algShort      = flag.String("a", "md5", "Hash algorithm (short)")
		chunkSize     = flag.Int("chunk-size", 4, "Chunk size in MB")
		chunkShort    = flag.Int("c", 4, "Chunk size in MB (short)")
		showProgress  = flag.Bool("progress", true, "Show progress")
		progressShort = flag.Bool("p", true, "Show progress (short)")
		help          = flag.Bool("help", false, "Show help")
		helpShort     = flag.Bool("h", false, "Show help (short)")
	)

	flag.Parse()

	// Show help if requested
	if *help || *helpShort {
		printUsage()
		return
	}

	// Get file path from arguments
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Error: Please specify exactly one file to hash")
		fmt.Println()
		printUsage()
		os.Exit(1)
	}

	filePath := args[0]

	// Use short flags if provided, otherwise use long flags
	selectedAlgorithm := *algorithm
	if flag.Lookup("a").Value.String() != "md5" {
		selectedAlgorithm = *algShort
	}

	selectedChunkSize := *chunkSize
	if flag.Lookup("c").Value.String() != "4" {
		selectedChunkSize = *chunkShort
	}

	selectedProgress := *showProgress
	if flag.Lookup("p").Value.String() != "true" {
		selectedProgress = *progressShort
	}

	// Parse algorithm
	hashAlg, err := parseAlgorithm(selectedAlgorithm)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' does not exist\n", filePath)
		os.Exit(1)
	}

	// Create hash calculator with custom chunk size
	calculator := &HashCalculator{
		ChunkSize: int64(selectedChunkSize) * 1024 * 1024, // Convert MB to bytes
	}

	fmt.Printf("Calculating %s hash for: %s\n", getAlgorithmName(hashAlg), filePath)
	fmt.Printf("Chunk size: %d MB\n", selectedChunkSize)
	fmt.Println()

	// Define progress callback
	var progressCallback func(float64)
	if selectedProgress {
		progressCallback = progressBar
	}

	// Calculate hash
	result, err := calculator.CalculateFileHash(filePath, hashAlg, progressCallback)
	if err != nil {
		fmt.Printf("Error calculating hash: %v\n", err)
		os.Exit(1)
	}

	// Display results
	fmt.Println()
	fmt.Println("Hash calculation complete!")
	fmt.Println("=" + strings.Repeat("=", 50))
	fmt.Printf("File: %s\n", result.Filename)
	fmt.Printf("Size: %s\n", formatBytes(result.FileSize))
	fmt.Printf("Algorithm: %s\n", getAlgorithmName(result.Algorithm))
	fmt.Printf("Hash: %s\n", result.Hash)
	fmt.Println("=" + strings.Repeat("=", 50))
	fmt.Println()
	fmt.Println("Description:")
	fmt.Println(result.Description)
}
