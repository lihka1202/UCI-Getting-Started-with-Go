package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
*Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

Submit your source code for the program,
“makejson.go”.
*/
func main() {
	var name string
	var address string

	// Read the name in
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.Replace(name, "\r\n", "", -1)

	// Read the address in
	fmt.Print("Enter your address: ")
	newReader := bufio.NewReader(os.Stdin)
	address, err = newReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	address = strings.Replace(address, "\r\n", "", -1)

	// Convert to JSON
	hashmap := map[string]string{"name": name, "address": address}
	barr, err := json.Marshal(hashmap)

	//! To print it nicely
	var prettyJSON bytes.Buffer
	jsonErr := json.Indent(&prettyJSON, barr, "", "\t")
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	fmt.Println(strings.Replace(string(prettyJSON.Bytes()), "\r\n", "", -1))
}
