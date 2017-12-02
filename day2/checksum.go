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
	Advent of Code Day 2 - http://adventofcode.com/2017/day/2
*/

func main() {
	// get spreadsheet from file
	spreadsheet := getSpreadsheet()

	// calculate checksum for spreadsheet
	checksum1 := calcChecksumPart1(spreadsheet)
	checksum2 := calcChecksumPart2(spreadsheet)
	fmt.Println("Spreadsheet checksum for part 1 is:", checksum1)
	fmt.Println("Spreadsheet checksum for part 2 is:", checksum2)
}

// helper function to parse text file containing spreadsheet of numbers
func getSpreadsheet() [][]int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var fileLines []string

	// read file into array of lines
	for scanner.Scan() {
		row := scanner.Text()
		fileLines = append(fileLines, row)
	}
	spreadsheet := make([][]int, len(fileLines))

	// loop through each line in file
	for i, line := range fileLines {
		var numbersRow []int
		numArray := strings.Split(line, " ")

		// convert string to int and put in spreadsheet
		for _, str := range numArray {
			num, _ := strconv.Atoi(str)
			numbersRow = append(numbersRow, num)
		}
		spreadsheet[i] = append(spreadsheet[i], numbersRow...)
	}
	return spreadsheet
}

/*
	Part 1 Rules:
	The spreadsheet consists of rows of apparently-random numbers. To make sure the recovery process is on the right track, they need you to calculate the spreadsheet's checksum. For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.

	For example, given the following spreadsheet:

	5 1 9 5
	7 5 3
	2 4 6 8
	The first row's largest and smallest values are 9 and 1, and their difference is 8.
	The second row's largest and smallest values are 7 and 3, and their difference is 4.
	The third row's difference is 6.
	In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.
*/
func calcChecksumPart1(spreadsheet [][]int) int {
	differences := make([]int, len(spreadsheet))

	// calculate differences for each row
	for rowNumber, row := range spreadsheet {
		min := row[0]
		max := row[0]
		for colNumber, val := range row {
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
			// last number in row, update differences
			if colNumber == len(row)-1 {
				differences[rowNumber] = max - min
			}
		}
	}

	// calculate checksum based on row differences
	checksum := 0
	for _, difference := range differences {
		checksum += difference
	}
	return checksum
}

/*
	Part 2 Rules:
	It sounds like the goal is to find the only two numbers in each row where one evenly divides the other - that is, where the result of the division operation is a whole number. They would like you to find those numbers on each line, divide them, and add up each line's result.

	For example, given the following spreadsheet:

	5 9 2 8
	9 4 7 3
	3 8 6 5
	In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
	In the second row, the two numbers are 9 and 3; the result is 3.
	In the third row, the result is 2.
	In this example, the sum of the results would be 4 + 3 + 2 = 9.
*/
func calcChecksumPart2(spreadsheet [][]int) int {
	return 0
}
