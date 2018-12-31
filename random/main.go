package main

import (
	"fmt"
	"math"

	"github.com/xoreo/go-basics/testpkg"
)

func main() {
	fmt.Println("hell, world!")
	fmt.Println(fmt.Sprintf("test: %s", "despacito"))
	fmt.Println(math.Pow(2, 2))

	num1 := 17
	num2 := 5
	fmt.Println(num1 + num2)

	array := [5]string{"test", "test", "test", "test", "test"}
	fmt.Println(array)

	slice := []string{"test", "test"}
	slice = append(slice, "test")

	testInstance := testpkg.Test()

	stuff := make(map[string]string)
	stuff["one"] = "one value"
	stuff["two"] = "two value"
	stuff["three"] = "three value"

	fmt.Println(stuff)
	fmt.Println(testInstance)

	for _, sliceMember := range slice {
		fmt.Println(sliceMember)
	}
}
