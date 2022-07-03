package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		ret := 1/2.0*a*t*t + v*t + s
		return ret
	}
	return fn
}

func ParseInput(scanner *bufio.Scanner, n int) []float64 {
	var ret []float64
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")

	if len(s) != n {
		fmt.Fprintln(os.Stderr, "\nInput values error")
		os.Exit(1)
	}

	for _, item := range s {
		x, err := strconv.ParseFloat(item, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\nInput data error")
			os.Exit(1)
		}
		ret = append(ret, x)
	}

	return ret
}

// 1, 2, 3, 2; 9.000
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter three initial values for a, v, s:")
	d1 := ParseInput(scanner, 3)

	fmt.Println("\nEnter a value for time:")
	d2 := ParseInput(scanner, 1)

	g := GenDisplaceFn(d1[0], d1[1], d1[2])
	fmt.Printf("\nOutput: %.3f\n", g(d2[0]))
}
