package main

import (
	"fmt"
	"io/ioutil"
)

/*
	Author: Amy Schlesener - github.com/aschlesener
	Advent of Code Day 1 - http://adventofcode.com/2017/day/1
*/

func main() {
	// get captcha from input file
	captcha := getCaptcha()

	// calculate results
	resultPart1 := solveCaptchaPart1(captcha)
	resultPart2 := solveCaptchaPart2(captcha)
	fmt.Println("Captcha result for part 1 is:", resultPart1)
	fmt.Println("Captcha result for part 2 is:", resultPart2)
}

// helper function to parse text file containing captcha string
func getCaptcha() string {
	bytes, _ := ioutil.ReadFile("input.txt")
	return string(bytes)
}

/*
	Part 1 Rules:
	The captcha requires you to review a sequence of digits (your puzzle input) and find the sum of all digits that match the next digit in the list. The list is circular, so the digit after the last digit is the first digit in the list.

	For example:

	1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
	1111 produces 4 because each digit (all 1) matches the next.
	1234 produces 0 because no digit matches the next.
	91212129 produces 9 because the only digit that matches the next one is the last digit, 9.
*/
func solveCaptchaPart1(input string) int {
	sum := 0
	ptr2 := 1

	if len(input) < 2 {
		return 0
	}

	for ptr1, char := range input {
		if ptr1 == len(input)-1 {
			// at end of string, check last number against first
			ptr2 = 0
		}
		if char == rune(input[ptr2]) {
			// numbers match, add number to total sum
			sum += int(char - '0')
		}
		ptr2++
	}

	return sum
}

/*
	Part 2 Rules:
	Now, instead of considering the next digit, it wants you to consider the digit halfway around the circular list. That is, if your list contains 10 items, only include a digit in your sum if the digit 10/2 = 5 steps forward matches it. Fortunately, your list has an even number of elements.

	For example:

	1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
	1221 produces 0, because every comparison is between a 1 and a 2.
	123425 produces 4, because both 2s match each other, but no other digit has a match.
	123123 produces 12.
	12131415 produces 4.
*/
func solveCaptchaPart2(input string) int {
	if len(input) < 2 {
		return 0
	}

	sum := 0
	// initialize second pointer to halfway position
	halfway := len(input) / 2
	ptr2 := halfway

	for ptr1, char := range input {
		if ptr1 == halfway {
			// at end of string, wrap around list to find next comparison point
			ptr2 = 0
		}
		if char == rune(input[ptr2]) {
			// numbers match, add number to total sum
			sum += int(char - '0')
		}
		ptr2++
	}

	return sum
}
