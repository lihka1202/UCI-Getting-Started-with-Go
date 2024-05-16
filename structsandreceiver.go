package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func main() {
	cow := Animal{
		food:       "Grass",
		locomotion: "Walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "Worms",
		locomotion: "Fly",
		noise:      "Peep",
	}

	snake := Animal{
		food:       "Mice",
		locomotion: "Slither",
		noise:      "hsss",
	}

	end := false

	for end == false {
		fmt.Println("> ")
		var input string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}

		combiString := strings.Fields(input)
		if len(combiString) != 2 {
			fmt.Println("> Invalid input")
			fmt.Println("Ending this session. Bye!")
			break
		}
		ani, action := combiString[0], combiString[1]
		if strings.ToLower(action) == "eat" && strings.ToLower(ani) == "cow" {
			cow.Eat()
		} else if strings.ToLower(action) == "eat" && strings.ToLower(ani) == "bird" {
			bird.Eat()

		} else if strings.ToLower(action) == "eat" && strings.ToLower(ani) == "snake" {
			snake.Eat()

		} else if strings.ToLower(action) == "move" && strings.ToLower(ani) == "cow" {
			cow.Move()

		} else if strings.ToLower(action) == "move" && strings.ToLower(ani) == "bird" {
			bird.Move()

		} else if strings.ToLower(action) == "move" && strings.ToLower(ani) == "snake" {
			snake.Move()

		} else if strings.ToLower(action) == "speak" && strings.ToLower(ani) == "cow" {
			cow.Speak()

		} else if strings.ToLower(action) == "speak" && strings.ToLower(ani) == "bird" {
			bird.Speak()

		} else if strings.ToLower(action) == "speak" && strings.ToLower(ani) == "snake" {
			snake.Speak()

		} else {
			fmt.Println("> Invalid input")
			fmt.Println("Ending this session. Bye!")
			end = true
		}
	}

}
