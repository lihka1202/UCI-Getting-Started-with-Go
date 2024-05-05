package main

import "fmt"

func main() {
	var input float32
	fmt.Println("Please input float number of your choice")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input")
	}
	changedNum := int32(input)

	fmt.Printf("The changed number is: %d\n", changedNum)

}
