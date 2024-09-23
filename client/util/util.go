package util

import (
	"crypto/sha256"
	"fmt"
	"os"
	"sync"
)

// verifyChecksum verifies if the checksum matches the received data
func VerifyChecksum(data []byte, expectedChecksum string) bool {
	checksum := sha256.Sum256(data)
	return fmt.Sprintf("%x", checksum) == expectedChecksum
}

func CreateFileDescriptors(filePath string, num int) ([]*os.File, []sync.Mutex, error) {
	files := make([]*os.File, num)
	mutexes := make([]sync.Mutex, num)
	for i := 0; i < num; i++ {
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return nil, nil, err
		}
		files[i] = file
	}
	return files, mutexes, nil
}
