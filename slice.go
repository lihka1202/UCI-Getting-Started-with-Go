package main

import (
	"fmt"
	"slices"
)

func printSlice(input []int32) {
	for i := 0; i < len(input); i++ {
		fmt.Print(input[i], " ")
	}
}

func main() {
	sli := make([]int32, 3)
	for {
		fmt.Println("\nPlease enter a value to add to the slice: ")
		// Add the value to the slice somehow
		var input int32
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Entered an X for the input, exiting the program")
			return
		}

		sli = append(sli, input)
		slices.Sort(sli)

		//! Print the slice here
		printSlice(sli)
	}

}
