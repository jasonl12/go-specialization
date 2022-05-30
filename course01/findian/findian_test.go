package main

import (
	"fmt"
	"testing"
)

// go test -v ./
func TestFindian(t *testing.T) {
	var Tests = []struct {
		input  string
		expect string
	}{
		{"ian", "Found!"},
		{"Ian", "Found!"},
		{"iuiygaygn", "Found!"},
		{"I d skd a efju N", "Found!"},
		{"I A N", "Found!"},
		{"ihhhhhn", "Not Found!"},
		{"ina", "Not Found!"},
		{"xian", "Not Found!"},
		{"xyzfoobar", "Not Found!"},
		{"1234", "Not Found!"},
	}

	for _, tt := range Tests {
		fmt.Printf("input: %v; expect: %v\noutput: %v\n", tt.input, tt.expect,
			find(tt.input))
	}
}
