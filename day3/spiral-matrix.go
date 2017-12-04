package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

/*
	Author: Amy Schlesener - github.com/aschlesener
	Advent of Code Day 3 - http://adventofcode.com/2017/day/3
*/

func main() {
	// get number of square from file
	squareNumber := getNumber()

	// calculate steps required to reach given number
	steps1 := calcStepsPart1(squareNumber)
	fmt.Println("Number of steps required for part 1 is:", steps1)
}

// helper function to parse text file containing spreadsheet of numbers
func getNumber() int {
	bytes, _ := ioutil.ReadFile("input.txt")
	num, _ := strconv.Atoi(string(bytes))
	return num
}

/*
	Part 1 Rules:
	You come across an experimental new kind of memory stored on an infinite two-dimensional grid.

	Each square on the grid is allocated in a spiral pattern starting at a location marked 1 and then counting up while spiraling outward. For example, the first few squares are allocated like this:

	17  16  15  14  13
	18   5   4   3  12
	19   6   1   2  11
	20   7   8   9  10
	21  22  23---> ...
	While this is very space-efficient (no squares are skipped), requested data must be carried back to square 1 (the location of the only access port for this memory system) by programs that can only move up, down, left, or right. They always take the shortest path: the Manhattan Distance between the location of the data and square 1.

	For example:

	Data from square 1 is carried 0 steps, since it's at the access port.
	Data from square 12 is carried 3 steps, such as: down, left, left.
	Data from square 23 is carried only 2 steps: up twice.
	Data from square 1024 must be carried 31 steps.
	How many steps are required to carry the data from the square identified in your puzzle input all the way to the access port?
*/
func calcStepsPart1(num int) int {
	// TOOD: calculate matrix size based on input number (should be the next odd perfect square that is greater than the input number)
	matrixSize := 363609
	_, centerCoord, numCoord := generateSpiralMatrix(matrixSize, num)
	// calculate Manhattan distance from the given number to square 1 (the center square)
	distance := int(math.Abs(float64(numCoord[0]-centerCoord)) + math.Abs(float64(numCoord[1]-centerCoord)))
	return distance
}

// function to generate a spiral matrix of a given size, where size is the final lower-right number in the matrix
// also returns the location of the center number and the number we want to calculate the distance for
func generateSpiralMatrix(size int, searchValue int) ([][]int, int, []int) {
	// TODO: check for valid size input (must be an odd perfect square number)
	// spiral sequence is such that we repeat going a) right and up, then b) left and down
	right := true
	// matrix height and width will be the square root of the size
	matrixHeight := int(math.Sqrt(float64(size)))
	matrix := make([][]int, matrixHeight)
	for i := 0; i < matrixHeight; i++ {
		matrix[i] = make([]int, matrixHeight)
	}
	// e.g. if matrix is size 5, center coordinates will be (2, 2) and value is 1
	centerCoord := matrixHeight / 2
	matrix[centerCoord][centerCoord] = 1
	numCoord := make([]int, 2)
	number := 1
	// last examined coordinate used in algorithm when setting number
	lastCoord := make([]int, 2)
	lastCoord[0] = centerCoord
	lastCoord[1] = centerCoord

	// for i = height of matrix, traverse around matrix in a spiral, filling out numbers along the way
	for i := 1; i <= matrixHeight; i++ {
		if right {
			right = false
			traverseMatrixDirection("right", matrix, searchValue, &number, lastCoord, i, matrixHeight, numCoord)
			if i != matrixHeight {
				traverseMatrixDirection("up", matrix, searchValue, &number, lastCoord, i, matrixHeight, numCoord)
			}
		} else {
			right = true
			traverseMatrixDirection("left", matrix, searchValue, &number, lastCoord, i, matrixHeight, numCoord)
			traverseMatrixDirection("down", matrix, searchValue, &number, lastCoord, i, matrixHeight, numCoord)
		}
	}

	return matrix, centerCoord, numCoord
}

// helper function to traverse the matrix in a given direction, filling out numbers along the way
func traverseMatrixDirection(direction string, matrix [][]int, searchValue int, number *int, lastCoord []int, i int, size int, numCoord []int) {
	x := 0
	y := 0
	if i == size {
		// end of matrix - stop one early
		i = i - 1
	}

	for j := 0; j < i; j++ {
		if direction == "right" {
			x = lastCoord[0]
			y = lastCoord[1] + 1
			lastCoord[1] = y
		} else if direction == "up" {
			x = lastCoord[0] - 1
			y = lastCoord[1]
			lastCoord[0] = x
		} else if direction == "left" {
			x = lastCoord[0]
			y = lastCoord[1] - 1
			lastCoord[1] = y
		} else if direction == "down" {
			x = lastCoord[0] + 1
			y = lastCoord[1]
			lastCoord[0] = x
		}

		matrix[x][y] = *number + 1
		*number++

		if *number == searchValue {
			numCoord[0] = lastCoord[0]
			numCoord[1] = lastCoord[1]
		}
	}
}
