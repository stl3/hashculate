package main

import (
	"os"
	"testing"
)

func TestHashCalculator(t *testing.T) {
	// Create a temporary test file
	testContent := "Hello, World!\nThis is a test file for hash calculation."
	testFile := "test_hash_calc.txt"

	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	calculator := NewHashCalculator()

	// Test MD5
	result, err := calculator.CalculateFileHash(testFile, MD5, nil)
	if err != nil {
		t.Fatalf("MD5 calculation failed: %v", err)
	}
	if result.Hash == "" {
		t.Error("MD5 hash is empty")
	}
	if result.Algorithm != MD5 {
		t.Errorf("Expected algorithm MD5, got %s", result.Algorithm)
	}

	// Test SHA256
	result, err = calculator.CalculateFileHash(testFile, SHA256, nil)
	if err != nil {
		t.Fatalf("SHA256 calculation failed: %v", err)
	}
	if result.Hash == "" {
		t.Error("SHA256 hash is empty")
	}
	if len(result.Hash) != 64 { // SHA256 produces 64 hex characters
		t.Errorf("Expected SHA256 hash length 64, got %d", len(result.Hash))
	}

	// Test file info
	if result.Filename != testFile {
		t.Errorf("Expected filename %s, got %s", testFile, result.Filename)
	}
	if result.FileSize != int64(len(testContent)) {
		t.Errorf("Expected file size %d, got %d", len(testContent), result.FileSize)
	}
}

func TestParseAlgorithm(t *testing.T) {
	tests := []struct {
		input    string
		expected HashAlgorithm
		hasError bool
	}{
		{"md5", MD5, false},
		{"MD5", MD5, false},
		{"sha1", SHA1, false},
		{"SHA-1", SHA1, false},
		{"sha256", SHA256, false},
		{"SHA-256", SHA256, false},
		{"sha512", SHA512, false},
		{"SHA-512", SHA512, false},
		{"invalid", "", true},
	}

	for _, test := range tests {
		result, err := parseAlgorithm(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for input %s, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %s: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("For input %s, expected %s, got %s", test.input, test.expected, result)
			}
		}
	}
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "0 bytes"},
		{100, "100 bytes"},
		{1023, "1023 bytes"},
		{1024, "1.0 kb (kilobytes)"},
		{2048, "2.0 kb (kilobytes)"},
		{1536, "1.5 kb (kilobytes)"},
	}

	for _, test := range tests {
		result := formatBytes(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %s, got %s", test.input, test.expected, result)
		}
	}
}

func TestNonExistentFile(t *testing.T) {
	calculator := NewHashCalculator()
	_, err := calculator.CalculateFileHash("nonexistent_file.txt", MD5, nil)
	if err == nil {
		t.Error("Expected error for non-existent file, but got none")
	}
}
