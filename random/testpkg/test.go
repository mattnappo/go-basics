package testpkg

// Example - test
type Example struct {
	x int
	y int
}

// Test - test
func Test() *Example {
	test := Example{x: 10, y: 10} // Int test

	return &test
}
