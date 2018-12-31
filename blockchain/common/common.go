package common

import (
	"os"
	"path/filepath"

	"golang.org/x/crypto/sha3"
)

// Sha3 - Hash input using sha3
func Sha3(b []byte) []byte {
	hash := sha3.New256()
	hash.Write(b)
	return hash.Sum(nil)
}

// CreateDirIfDoesNotExist - Create a directory if it does not already exist
func CreateDirIfDoesNotExist(dir string) error {
	dir = filepath.FromSlash(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
