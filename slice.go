package main

import "fmt"

func printSlice(input []int32) {
	for i := 0; i < len(input); i++ {
		fmt.Print(input[i], " ")
	}
}

func main() {
	count := 0
	sli := make([]int32, 3)
	for {
		if count > 2 {
			fmt.Println("\nPlease enter a value to add to the slice: ")
		} else {
			fmt.Println("\nPlease enter a value to add to slice (note that the initialized values of the slice are 0): ")
		}

		// Add the value to the slice somehow
		var input int32
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Entered an X for the input, exiting the program")
			return
		}

		if count > 2 {
			sli = append(sli, input)
		} else {
			sli[count] = input
			count += 1
		}

		//! Print the slice here
		printSlice(sli)
	}

}
