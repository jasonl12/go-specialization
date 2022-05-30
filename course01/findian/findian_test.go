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

// The program should print "Found!" for the following example entered strings,
// "ian", "Ian", "iuiygaygn", "I d skd a efju N". The program should print
// "Not Found!" for the following strings, "ihhhhhn", "ina", "xian".
