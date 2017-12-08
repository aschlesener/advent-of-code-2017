package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	Author: Amy Schlesener - github.com/aschlesener
	Advent of Code Day 8 - http://adventofcode.com/2017/day/8
*/

func main() {
	// get instructions
	instructions := getInstructions()

	// calculate max register value
	maxRegisterPart1 := calcMaxRegister(instructions, false)
	fmt.Println("Largest value in any register for part 1 is:", maxRegisterPart1)
	maxRegisterPart2 := calcMaxRegister(instructions, true)
	fmt.Println("Largest value ever held in any register for part 2 is:", maxRegisterPart2)
}

// helper function to parse text file containing list of string instructions
func getInstructions() []string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var instructions []string

	// read file into array of lines
	for scanner.Scan() {
		row := scanner.Text()
		instructions = append(instructions, row)
	}
	return instructions
}

/*
	Part 1 Rules:
	You receive a signal directly from the CPU. Because of your recent assistance with jump instructions, it would like you to compute the result of a series of unusual register instructions.

	Each instruction consists of several parts: the register to modify, whether to increase or decrease that register's value, the amount by which to increase or decrease it, and a condition. If the condition fails, skip the instruction without modifying the register. The registers all start at 0. The instructions look like this:

	b inc 5 if a > 1
	a inc 1 if b < 5
	c dec -10 if a >= 1
	c inc -20 if c == 10
	These instructions would be processed as follows:

	Because a starts at 0, it is not greater than 1, and so b is not modified.
	a is increased by 1 (to 1) because b is less than 5 (it is 0).
	c is decreased by -10 (to 10) because a is now greater than or equal to 1 (it is 1).
	c is increased by -20 (to -10) because c is equal to 10.
	After this process, the largest value in any register is 1.

	You might also encounter <= (less than or equal to) or != (not equal to). However, the CPU doesn't have the bandwidth to tell you what all the registers are named, and leaves that to you to determine.

	What is the largest value in any register after completing the instructions in your puzzle input?

	Part 2 Rules:
	To be safe, the CPU also needs to know the highest value held in any register during this process so that it can decide how much memory to allocate to these operations. For example, in the above instructions, the highest value ever held was 10 (in register c after the third instruction was evaluated).
*/
func calcMaxRegister(instructions []string, isPart2 bool) int {
	// store registers in dictionary for fast lookup by name
	dict := make(map[string]int)
	max := -10000
	// loop through instructions
	for _, instruction := range instructions {
		instructionSet := strings.Split(instruction, " ")
		// parse if condition
		registerToCheck := instructionSet[4]
		registerToCheckValue, registerToCheckExists := dict[registerToCheck]
		if !registerToCheckExists {
			dict[registerToCheck] = 0
		}
		b, _ := strconv.Atoi(instructionSet[6])
		if compare(registerToCheckValue, b, instructionSet[5]) {
			// if condition is true, inc/dec register
			register := instructionSet[0]
			_, registerExists := dict[register]
			if !registerExists {
				dict[register] = 0
			}
			valueToChangeBy, _ := strconv.Atoi(instructionSet[2])
			if instructionSet[1] == "inc" {
				dict[register] += valueToChangeBy
			} else if instructionSet[1] == "dec" {
				dict[register] -= valueToChangeBy
			}
			if isPart2 {
				if dict[register] > max {
					max = dict[register]
				}
			}
		}
	}
	if !isPart2 {
		// loop through map to find greatest value
		for _, v := range dict {
			if v > max {
				max = v
			}
		}
	}
	return max
}

// helper func to compare two values given a string comparison token
func compare(a int, b int, comparisonToken string) bool {
	if comparisonToken == ">" {
		return a > b
	} else if comparisonToken == "<" {
		return a < b
	} else if comparisonToken == ">=" {
		return a >= b
	} else if comparisonToken == "<=" {
		return a <= b
	} else if comparisonToken == "==" {
		return a == b
	} else if comparisonToken == "!=" {
		return a != b
	}
	return a == b
}
