package common

import "testing"

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
