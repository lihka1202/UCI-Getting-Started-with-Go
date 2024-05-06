package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func findian(input string) string {
	input = strings.ToLower(input)

	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		return "Found!"
	}
	return "Not Found!"
}

func main() {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	line = strings.TrimSpace(line)
	fmt.Println(findian(line))
	fmt.Println("\nHere are some other cases where this code works:")
	fmt.Println("ian: ", findian("ian"))
	fmt.Println("Ian: ", findian("Ian"))
	fmt.Println("iuiygaygn: ", findian("iuiygaygn"))
	fmt.Println("“I d skd a efju N: ", findian("I d skd a efju N"))
	fmt.Println("ihhhhhn: ", findian("ihhhhhn"))
	fmt.Println("ina: ", findian("ina"))
	fmt.Println("xian: ", findian("xian"))

}
