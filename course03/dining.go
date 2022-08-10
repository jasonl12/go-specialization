package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

// sharing chopsticks with adjacent pair, each philosopher is numbered
type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

var wg sync.WaitGroup

// buffered channel for 2 eat concurrency
func host(ch chan bool) {
	ch <- true
	ch <- true
}

func (p Philo) eat(ch chan bool) {
	// Each philosopher should eat only 3 times
	for i := 0; i < 3; i++ {
		// must get permission from host which executes in its own goroutine
		<-ch
		p.leftCS.Lock()
		p.rightCS.Lock()

		// after it has obtained the lock, it prints
		fmt.Println("starting to eat", p.number)
		time.Sleep(3 * time.Second)

		// before it has released the lock, it prints
		fmt.Println("finishing eating", p.number)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		ch <- true
	}

	defer wg.Done()
}

func randbool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		// pick up the chopsticks in any order, not the lowest-numbered first
		if randbool() {
			philos[i] = &Philo{
				CSticks[i],
				CSticks[(i+1)%5],
				i + 1,
			}
		} else {
			philos[i] = &Philo{
				CSticks[(i+1)%5],
				CSticks[i],
				i + 1, // Each philosopher is numbered, 1 through 5
			}
		}
	}

	// The host allows no more than 2 philosophers to eat concurrently
	ch := make(chan bool, 2)
	go host(ch)

	// Start the dining
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(ch)
	}

	wg.Wait()
}
