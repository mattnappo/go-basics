package util

import (
	"fmt"
	"testing"
)

func TestGetRandomChain(t *testing.T) {
	chain := GetRandomChain(3)
	fmt.Println(chain.String())
	chain = GetRandomChain(5)
	fmt.Println(chain.String())
}
