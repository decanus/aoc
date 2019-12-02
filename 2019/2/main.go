package main

import "fmt"

func main() {

	input := []int{
		1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,10,19,1,19,6,23,2,23,13,27,1,27,5,31,2,31,10,35,1,9,35,39,1,39,9,43,2,9,43,47,1,5,47,51,2,13,51,55,1,55,9,59,2,6,59,63,1,63,5,67,1,10,67,71,1,71,10,75,2,75,13,79,2,79,13,83,1,5,83,87,1,87,6,91,2,91,13,95,1,5,95,99,1,99,2,103,1,103,6,0,99,2,14,0,0,
	}

	fmt.Println(run(input, 12, 2)[0])
	fmt.Print(runUntil(input, 19690720))
}

func runUntil(input []int, val int) int {
	for noun := 0; noun <= 99; noun++  {
		for verb := 0; verb <= 99; verb++ {
			result := run(input, noun, verb)
			if result[0] == val {
				return 100*noun + verb
			}
		}
	}

	return 0
}

func run(scratch []int, noun int, verb int) []int {
	current := 0

	input := make([]int, len(scratch))
	copy(input, scratch)

	input[1] = noun
	input[2] = verb
	output := input

	for {
		if input[current] == 99 {
			break
		}

		if current+3 >= len(input) {
			break
		}

		if input[current] == 1 {
			output[input[current+3]] = input[input[current+1]] + input[input[current+2]]
		} else if input[current] == 2 {
			output[input[current+3]] = input[input[current+1]] * input[input[current+2]]
		}

		current += 4
	}

	return output
}