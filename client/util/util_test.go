package util

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

// TestVerifyChecksum tests the VerifyChecksum function.
func TestVerifyChecksum(t *testing.T) {
	data := []byte("Hello, World!")
	expectedChecksum := fmt.Sprintf("%x", sha256.Sum256(data))

	// Test case 1: Correct checksum
	if !VerifyChecksum(data, expectedChecksum) {
		t.Errorf("Expected checksum %s to be valid", expectedChecksum)
	}

	// Test case 2: Incorrect checksum
	if VerifyChecksum(data, "invalidchecksum") {
		t.Errorf("Expected checksum to be invalid")
	}

	// Test case 3: Empty data and checksum
	if VerifyChecksum([]byte{}, "") {
		t.Errorf("Expected empty checksum to be invalid")
	}
}

// TestCreateFileDescriptors tests the CreateFileDescriptors function.
func TestCreateFileDescriptors(t *testing.T) {
	// Create a temporary file to test file descriptor creation.
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up the file after test.

	// Test case 1: Correct number of file descriptors.
	numDescriptors := 3
	files, mutexes, err := CreateFileDescriptors(tempFile.Name(), numDescriptors)
	if err != nil {
		t.Fatalf("Failed to create file descriptors: %v", err)
	}
	if len(files) != numDescriptors {
		t.Errorf("Expected %d file descriptors, got %d", numDescriptors, len(files))
	}
	if len(mutexes) != numDescriptors {
		t.Errorf("Expected %d mutexes, got %d", numDescriptors, len(mutexes))
	}

	// Ensure that all file descriptors are open.
	for _, file := range files {
		if file == nil {
			t.Errorf("File descriptor is nil")
		} else {
			file.Close() // Close file to release resources.
		}
	}

	// Test case 2: Invalid file path.
	_, _, err = CreateFileDescriptors("/invalid/path/to/file", numDescriptors)
	if err == nil {
		t.Error("Expected an error for invalid file path, got nil")
	}
}
