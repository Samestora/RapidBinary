package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// CalculateSHA256 reads a file in chunks and returns the hash
func CalculateSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha256.New()
	// Use a 1MB buffer for speed, especially on limited RAM systems
	if _, err := io.CopyBuffer(h, file, make([]byte, 1024*1024)); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
