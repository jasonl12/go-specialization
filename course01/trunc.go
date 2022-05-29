package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// go build -o trunc trunc.go
func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		var s string

		for {
			fmt.Print("\nEnter a floating point number: ")
			fmt.Scan(&s)

			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Println("Please enter floating point number.")
				continue
			}

			fmt.Printf("%d\n", int(f))
		}
	}()

	<-ch
	fmt.Println("\n\nProgram is going to exit...")
}
