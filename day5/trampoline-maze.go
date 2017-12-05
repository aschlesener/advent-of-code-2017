package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Author: Amy Schlesener - github.com/aschlesener
	Advent of Code Day 5 - http://adventofcode.com/2017/day/5
*/

func main() {
	// get maze instructions from file
	instructions := getInstructions()

	// calculate how many steps it takes to escape the maze
	numStepsPart1 := countNumStepsPart1(instructions)
	instructions = getInstructions()
	numStepsPart2 := countNumStepsPart2(instructions)
	fmt.Println("Number of steps to escape the maze for part 1 is:", numStepsPart1)
	fmt.Println("Number of steps to escape the maze for part 2 is:", numStepsPart2)
}

// helper function to parse text file containing list of integer instructions
func getInstructions() []int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var instructions []int

	// read file into array of ints
	for scanner.Scan() {
		instruction, _ := strconv.Atoi(scanner.Text())
		instructions = append(instructions, instruction)
	}

	return instructions
}

/*
	Part 1 Rules:
	An urgent interrupt arrives from the CPU: it's trapped in a maze of jump instructions, and it would like assistance from any programs with spare cycles to help find the exit.

	The message includes a list of the offsets for each jump. Jumps are relative: -1 moves to the previous instruction, and 2 skips the next one. Start at the first instruction in the list. The goal is to follow the jumps until one leads outside the list.

	In addition, these instructions are a little strange; after each jump, the offset of that instruction increases by 1. So, if you come across an offset of 3, you would move three instructions forward, but change it to a 4 for the next time it is encountered.

	For example, consider the following list of jump offsets:

	0
	3
	0
	1
	-3
	Positive jumps ("forward") move downward; negative jumps move upward. For legibility in this example, these offset values will be written all on one line, with the current instruction marked in parentheses. The following steps would be taken before an exit is found:

	(0) 3  0  1  -3  - before we have taken any steps.
	(1) 3  0  1  -3  - jump with offset 0 (that is, don't jump at all). Fortunately, the instruction is then incremented to 1.
	2 (3) 0  1  -3  - step forward because of the instruction we just modified. The first instruction is incremented again, now to 2.
	2  4  0  1 (-3) - jump all the way to the end; leave a 4 behind.
	2 (4) 0  1  -2  - go back to where we just were; increment -3 to -2.
	2  5  0  1  -2  - jump 4 steps forward, escaping the maze.
	In this example, the exit is reached in 5 steps.

	How many steps does it take to reach the exit?
*/
func countNumStepsPart1(instructions []int) int {
	numSteps := 0
	nextPosition := 0
	for nextPosition < len(instructions) {
		currPosition := nextPosition
		nextPosition = nextPosition + instructions[nextPosition]
		instructions[currPosition]++
		numSteps++
	}
	return numSteps
}

/*
	Part 2 Rules:
	Now, the jumps are even stranger: after each jump, if the offset was three or more, instead decrease it by 1. Otherwise, increase it by 1 as before.

	Using this rule with the above example, the process now takes 10 steps, and the offset values after finding the exit are left as 2 3 2 3 -1.

	How many steps does it now take to reach the exit?
*/
func countNumStepsPart2(instructions []int) int {
	numSteps := 0
	nextPosition := 0
	for nextPosition < len(instructions) {
		currPosition := nextPosition
		offset := instructions[nextPosition]
		nextPosition = nextPosition + instructions[nextPosition]
		if offset > 2 {
			instructions[currPosition]--
		} else {
			instructions[currPosition]++
		}
		numSteps++
	}
	return numSteps
}
