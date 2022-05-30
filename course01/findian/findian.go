package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func find(input string) string {
	var ret string

	s := strings.ToLower(input)
	s = strings.TrimSuffix(s, "\n")
	last := len(s) - 1

	// i, a, n; [105 97 110], 3
	// fmt.Printf("%v, %d\n", []byte(s), len(s))

	i_index := strings.Index(s, "i")
	a_index := strings.Index(s, "a")
	n_index := strings.Index(s, "n")

	if i_index == 0 && n_index == last && a_index != -1 {
		ret = fmt.Sprintln("Found!")
	} else {
		ret = fmt.Sprintln("Not Found!")
	}

	return ret
}

// go build -o findian findian.go
func main() {
	fmt.Print("\nEnter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Print(find(input))
}
