package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	var sliceHolder []person

	//! Get the filename from the user
	var fileName string
	fmt.Println("Please enter the name of the file: ")
	_, err := fmt.Scan(&fileName)

	if err != nil {
		fmt.Println(err)
	}

	//! Check if .txt is in the file name or not
	if !strings.Contains(fileName, ".txt") {
		fileName = fileName + ".txt"
	}

	//! Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		firstName, lastName := strings.Split(fileScanner.Text(), " ")[0], strings.Split(fileScanner.Text(), " ")[1]
		holder := person{firstName, lastName}
		sliceHolder = append(sliceHolder, holder)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}

	//! Go through the slices and get the values
	for elem := range sliceHolder {
		fmt.Println("firstName: "+sliceHolder[elem].fname, "\t", "lastName: "+sliceHolder[elem].lname)
	}

}
