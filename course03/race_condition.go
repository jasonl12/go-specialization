package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Explanation of race conditions:
// Race conditions are the outcomes of two different goroutines reading and
// writing to the same shared data at the same time, resulting an unexpected
// output.

// How the race condition can occur:
// x is a share variable with value 0 in main
// "goroutine 1" prints x with expecting "x = 0" in console
// "goroutine 2" increases x value and causes the race condition in this program
// "goroutine 1" prints "x = 0" or "x = 1" depends on non-deterministic ordering

/*
go run race_condition.go
goroutine 2: x = 1
goroutine 1: x = 1
main: x = 1

go run --race race_condition.go
goroutine 1: x = 0
==================
WARNING: DATA RACE
Write at 0x00c00013a008 by goroutine 8:
  main.main.func2()
      /home/jason/code/go/tmp/race_condition.go:63 +0x46

Previous read at 0x00c00013a008 by goroutine 7:
  main.main.func1()
      /home/jason/code/go/tmp/race_condition.go:58 +0x3a

Goroutine 8 (running) created at:
  main.main()
      /home/jason/code/go/tmp/race_condition.go:62 +0x124

Goroutine 7 (finished) created at:
  main.main()
      /home/jason/code/go/tmp/race_condition.go:57 +0xbd
==================
goroutine 2: x = 1
main: x = 1
Found 1 data race(s)
exit status 66
*/

func main() {
	x := 0
	wg.Add(2)

	go func() {
		fmt.Printf("goroutine 1: x = %d\n", x)
		wg.Done()
	}()

	go func() {
		x += 1
		fmt.Printf("goroutine 2: x = %d\n", x)
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("main: x = %d\n", x)
}
