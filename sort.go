package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Function to merge two sorted arrays
func merge(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)
	return result
}

// Function to read integers from user input
func readInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a series of integers separated by spaces:")
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Fields(input)
	arr := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Invalid input. Please enter integers only.")
			os.Exit(1)
		}
		arr[i] = num
	}
	return arr
}

// Function to sort a partition
func sortPartition(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting partition:", arr)
	sort.Ints(arr)
	fmt.Println("Sorted partition:", arr)
}

func main() {
	arr := readInput()
	length := len(arr)
	partSize := (length + 3) / 4 // To handle the case when length is not exactly divisible by 4

	var wg sync.WaitGroup
	sortedParts := make([][]int, 4)

	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if end > length {
			end = length
		}
		part := arr[start:end]
		sortedParts[i] = part
		wg.Add(1)
		go sortPartition(part, &wg)
	}

	wg.Wait()

	// Merge sorted partitions
	merged1 := merge(sortedParts[0], sortedParts[1])
	merged2 := merge(sortedParts[2], sortedParts[3])
	finalSorted := merge(merged1, merged2)

	fmt.Println("Final sorted array:", finalSorted)
}
