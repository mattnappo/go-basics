package common

import (
	"fmt"
	"testing"
)

func TestSha3(t *testing.T) {
	str := "I'm a string"
	bytes := []byte(str)
	hash := Sha3(bytes)
	t.Log(hash)
	t.Log("success")
}
func TestCreateDirIfDoesNotExist(t *testing.T) {
	err := CreateDirIfDoesNotExist("test/test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetNonce(t *testing.T) {
	t1 := GetNonce(0)
	fmt.Printf("0: %s\n", t1)

	t1 = GetNonce(25)
	fmt.Printf("25: %s\n", t1)

	t1 = GetNonce(50)
	fmt.Printf("50: %s\n", t1)

	t1 = GetNonce(75)
	fmt.Printf("75: %s\n", t1)

	t1 = GetNonce(100)
	fmt.Printf("100: %s\n", t1)
}
