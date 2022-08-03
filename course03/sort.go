package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// The goal of this activity is to explore the use of threads by creating a
// program for sorting integers that uses four goroutines to create four
// sub-arrays and then merge the arrays into a single array.

// Write a program to sort an array of integers. The program should partition
// the array into 4 parts, each of which is sorted by a different goroutine.
// Each partition should be of approximately equal size. Then the main goroutine
// should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers. Each
// goroutine which sorts 1/4 of the array should print the subarray that it
// will sort. When sorting is complete, the main goroutine should print the
// entire sorted list.

const parts = 4

func worker(input []int, ch chan<- []int) {
	if len(input) == 1 {
		ch <- input
	} else {
		s := input
		sort.Ints(s)
		ch <- s
	}
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func main() {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input a series of integers, then enter 'EOF' for sorting:")
	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "EOF" {
			break
		} else {
			items := strings.Split(text, " ")
			for _, item := range items {
				if i, err := strconv.Atoi(item); err == nil {
					nums = append(nums, i)
				}
			}
		}
	}

	if len(nums) == 0 {
		fmt.Println("\nNo input integers.")
		os.Exit(0)
	} else {
		fmt.Printf("\ninput integers: %v\n\n", nums)

		if len(nums) <= parts {
			sort.Ints(nums)
			fmt.Println(nums)
		} else {
			length := len(nums) / parts
			ch := make(chan []int, parts)

			for i := 0; i < parts; i++ {
				left := i * length
				right := (i + 1) * length

				if i == 3 {
					sub := nums[left:]
					go worker(sub, ch)
					break
				}
				sub := nums[left:right]
				go worker(sub, ch)
			}

			a, b, c, d := <-ch, <-ch, <-ch, <-ch
			fmt.Println("goroutine return:", a)
			fmt.Println("goroutine return:", b)
			fmt.Println("goroutine return:", c)
			fmt.Println("goroutine return:", d)

			r := merge(merge(a, b), merge(c, d))
			fmt.Printf("\nsort: %v\n", r)
		}
	}
}
