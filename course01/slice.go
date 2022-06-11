package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s string
	n := make([]int, 3)
	fmt.Printf("Press X to quit the loop\n")

	for {
		fmt.Print("\nEnter integer: ")
		fmt.Scan(&s)

		if s == "x" || s == "X" {
			break
		} else {
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Invalid input: %q\n", s)
				continue
			}
			n = append(n, i)
		}

		ret := n[3:]
		sort.Ints(ret)
		fmt.Printf("Sorted slice: %v\n", ret)
	}
}
