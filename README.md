# Hashculate - File Hash Calculator

[![CI](https://github.com/stl3/hashculate/workflows/CI/badge.svg)](https://github.com/stl3/hashculate/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/stl3/hashculate)](https://goreportcard.com/report/github.com/stl3/hashculate)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org)

A fast, cross-platform command-line file hash calculator written in Go. This tool provides functionality with enhanced performance and scriptability.

## Quick Start

```bash
# Download and build
git clone https://github.com/stl3/hashculate.git
cd hashculate
go build -o hashculate .

# Calculate MD5 hash
./hashculate myfile.txt

# Calculate SHA-256 hash
./hashculate -a sha256 myfile.txt
```

## Features

- **Multiple Hash Algorithms**: Supports MD5, SHA-1, SHA-256, and SHA-512
- **Chunked Processing**: Processes large files in configurable chunks (default 4MB)
- **Progress Tracking**: Real-time progress bar during hash calculation
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Command-Line Interface**: Easy-to-use CLI with flexible options

## Installation

### Build from Source

```bash
git clone https://github.com/stl3/hashculate.git
cd hashculate
go build -o hashculate .
```

### Prerequisites

- Go 1.18 or later
- No external dependencies (uses only Go standard library)

### Direct Run

```bash
go run main.go [options] <file>
```

## Usage

### Basic Usage

```bash
# Calculate MD5 hash (default)
./hashculate myfile.txt

# Calculate SHA-256 hash
./hashculate -algorithm sha256 myfile.txt
./hashculate -a sha256 myfile.txt  # Short form
```

### Advanced Options

```bash
# Use SHA-512 with 8MB chunks and no progress bar
./hashculate -a sha512 -c 8 -p=false largefile.bin

# Show help
./hashculate -help
./hashculate -h
```

### Command-Line Options

| Option | Short | Default | Description |
|--------|-------|---------|-------------|
| `-algorithm` | `-a` | `md5` | Hash algorithm (md5, sha1, sha256, sha512) |
| `-chunk-size` | `-c` | `4` | Chunk size in MB for processing large files |
| `-progress` | `-p` | `true` | Show progress bar during calculation |
| `-help` | `-h` | `false` | Show help message |

## Supported Hash Algorithms

- **MD5**: 128-bit hash (fast, but cryptographically broken)
- **SHA-1**: 160-bit hash (deprecated for security)
- **SHA-256**: 256-bit hash (recommended for most uses)
- **SHA-512**: 512-bit hash (highest security)

## Examples

### Calculate different hashes of the same file

```bash
./hashculate -a md5 document.pdf
./hashculate -a sha1 document.pdf
./hashculate -a sha256 document.pdf
./hashculate -a sha512 document.pdf
```

### Process large files efficiently

```bash
# Use larger chunks for better performance on large files
./hashculate -a sha256 -c 16 large-video.mp4

# Disable progress bar for scripting
./hashculate -a sha256 -p=false data.bin
```

## Output Format

The tool provides detailed output including:

- File name and size
- Selected hash algorithm
- Calculated hash value
- Human-readable description

Example output:
```
Calculating SHA-256 hash for: example.txt
Chunk size: 4 MB

Progress: [==================================================] 100%

Hash calculation complete!
===================================================
File: example.txt
Size: 1024 kb (kilobytes)
Algorithm: SHA-256
Hash: a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3
===================================================

Description:
"example.txt", with size of 1024 kb (kilobytes), and file hash using the hashing algorithm SHA-256 has the value : a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3.
```
## Performance

- **Memory Efficient**: Uses constant memory regardless of file size
- **Fast Processing**: Optimized chunked reading for large files
- **Configurable**: Adjust chunk size based on available memory and performance needs

## Requirements

- Go 1.18 or later
- No external dependencies

## Testing

Run the built-in tests to verify functionality:

```bash
# Test with a sample file
echo "Hello, World!" > test.txt
./hashculate test.txt
./hashculate -a sha256 test.txt
```

## Troubleshooting

### Common Issues

1. **Permission denied**: Ensure the executable has proper permissions
   ```bash
   chmod +x hashculate  # On Unix-like systems
   ```

2. **File not found**: Verify the file path is correct and the file exists

3. **Out of memory**: For very large files, try reducing chunk size:
   ```bash
   ./hashculate -c 2 largefile.bin  # Use 2MB chunks instead of 4MB
   ```

## Security Notes

- **MD5 and SHA-1**: These algorithms are cryptographically broken and should not be used for security purposes
- **SHA-256/SHA-512**: Recommended for integrity verification and security applications
- **Local Processing**: All calculations are performed locally; no data is transmitted over the network

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Changelog

### v1.0.0
- Initial release
- Support for MD5, SHA-1, SHA-256, SHA-512 algorithms
- Chunked file processing
- Progress tracking
- Cross-platform support
