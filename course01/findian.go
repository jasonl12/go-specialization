package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const offset = 32

type data struct {
	count int
	lower rune
	items []string
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		var input string

		for {
			var d data
			hash := map[rune]bool{
				97:  false,
				105: false,
				110: false,
			}

			fmt.Print("\nEnter a string: ")
			fmt.Scan(&input)

			for _, c := range input {
				if c < 96 {
					d.lower = c + offset
				} else {
					d.lower = c
				}

				if value, ok := hash[d.lower]; ok && !value {
					hash[d.lower] = true
					d.count++
					d.items = append(d.items, string(d.lower))
				}
			}

			if d.count == 3 {
				fmt.Println("Found!")
			} else if d.count == 0 {
				fmt.Println("Not Found!")
			} else {
				fmt.Printf("Only contains: %v\n", d.items)
			}
		}
	}()

	<-ch
	fmt.Println("\n\nProgram is going to exit...")
}
