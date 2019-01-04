package common

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"

	"golang.org/x/crypto/sha3"
)

// Difficulty - Amount of leading zeros needed for POW
const Difficulty = 1

// // Mutex - A mutex for block generation
// var Mutex = &sync.Mutex{}

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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()`~-_=+[{]}\\|;:'\".>/?,<"

// GetNonce - Return an array of random (read-able) bytes
func GetNonce(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	hash := Sha3(b)
	return []byte(fmt.Sprintf("%x", hash))
}
