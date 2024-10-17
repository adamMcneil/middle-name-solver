package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func genOrderArray(n int) []int {
	if n <= 0 {
		return nil
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func genRandomArray(n int) []int {
	if n <= 0 {
		return nil
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return arr
}

func shiftRightExceptFixed(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr // No need to shift if the array has 1 or fewer elements
	}

	result := make([]int, n)
	copy(result, arr) // Initialize result with the original array

	fixed := make([]bool, n)
	for i := 0; i < n; i++ {
		if arr[i] == i {
			fixed[i] = true
		}
	}

	for i := 0; i < n; i++ {
		if fixed[i] {
			continue
		}

		// Find the next index to shift the value to
		nextIdx := (i + 1) % n
		for fixed[nextIdx] {
			nextIdx = (nextIdx + 1) % n // Skip over fixed indices
		}

		result[nextIdx] = arr[i]
	}

	return result
}

func printArray(arr []int) {
	for i, x := range arr {
		if i == x {
			fmt.Printf("\033[32m%d\t\033[0m", x)
		} else {
			fmt.Printf("\033[31m%d\t\033[0m", x)
		}
	}

	fmt.Printf("| %d", numCorrect(arr))
	fmt.Println()
}

func isSorted(arr []int) bool {
	for i, x := range arr {
		if i != x {
			return false
		}
	}
	return true
}

func numCorrect(arr []int) int {
	count := 0
	for i, x := range arr {
		if i == x {
			count++
		}
	}
	return count
}

func compute(n int) int {
	x := genRandomArray(n)
	printArray(x)
	count := 0
	for !isSorted(x) {
		x = shiftRightExceptFixed(x)
		printArray(x)
		count++
	}
	fmt.Println(count)
	return count
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	rounds, _ := strconv.Atoi(os.Args[2])
	total := 0
	for i := 0; i < rounds; i++ {
		total += compute(n)
	}
	fmt.Printf("Average number of round: %f", float32(total)/float32(rounds))
}
