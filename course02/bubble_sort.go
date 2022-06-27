package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func BubbleSort(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}

func main() {
	nums := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter to exit...")

	for {
		fmt.Print("Enter integer: ")
		scanner.Scan()
		text := scanner.Text()

		if len(text) != 0 {
			intVar, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, intVar)
		} else {
			break
		}
	}

	fmt.Printf("\nInput slice: %v\n", nums)
	BubbleSort(nums)
	fmt.Printf("Bubble sort: %v\n", nums)
}
