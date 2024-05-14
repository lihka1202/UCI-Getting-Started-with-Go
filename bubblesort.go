package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	var numbers []int

	fmt.Println("Enter up to 10 integers, separated by spaces:")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	numStrs := strings.Fields(input)

	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Invalid input:", err)
			return
		}
		numbers = append(numbers, num)
	}

	if len(numbers) > 10 {
		fmt.Println("Too many integers entered. Please enter up to 10 integers.")
		return
	}

	BubbleSort(numbers)

	fmt.Println("Sorted numbers:", numbers)
}

func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
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
