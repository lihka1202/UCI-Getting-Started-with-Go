package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Define Cow, Bird, and Snake types
type Cow struct{}
type Bird struct{}
type Snake struct{}

// Implement the Animal interface for Cow
func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Implement the Animal interface for Bird
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Implement the Animal interface for Snake
func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		if len(parts) != 3 {
			fmt.Println("Invalid command. Please enter a command in the format 'newanimal name type' or 'query name info'.")
			continue
		}

		command, name, arg := parts[0], parts[1], parts[2]

		switch command {
		case "newanimal":
			switch arg {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Println("Unknown animal type. Please use 'cow', 'bird', or 'snake'.")
				continue
			}
			fmt.Println("Created it!")

		case "query":
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch arg {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unknown query. Please use 'eat', 'move', or 'speak'.")
			}

		default:
			fmt.Println("Unknown command. Please use 'newanimal' or 'query'.")
		}
	}
}
