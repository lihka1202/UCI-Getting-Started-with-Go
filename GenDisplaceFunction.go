package main

import (
	"fmt"
)

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + (s0)
	}

	return fn
}

func getInput(input string) float64 {
	var tempHolder float64
	fmt.Print("Please enter the ", input, ": ")
	_, err := fmt.Scan(&tempHolder)
	if err != nil {
		fmt.Println("Error reading input")
	}
	return tempHolder

}

func main() {
	var a, v0, s0, t float64

	a = getInput("acceleration")
	v0 = getInput("initial velocity")
	s0 = getInput("initial displacement")
	t = getInput("time")

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Println(fn(t))
}
