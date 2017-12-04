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
	steps2 := calcStepsPart2(squareNumber)
	fmt.Println("Number of steps required for part 1 is:", steps1)
	fmt.Println("First value larger than input number for part 2 is:", steps2)
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
	_, centerCoord, numCoord := generateSpiralMatrix(matrixSize, num, false)
	// calculate Manhattan distance from the given number to square 1 (the center square)
	distance := int(math.Abs(float64(numCoord[0]-centerCoord)) + math.Abs(float64(numCoord[1]-centerCoord)))
	return distance
}

// function to generate a spiral matrix of a given size, where size is the final lower-right number in the matrix
// also returns the location of the center number and the number we want to calculate the distance for
func generateSpiralMatrix(size int, searchValue int, part2 bool) ([][]int, int, []int) {
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
			stop := traverseMatrixDirection("right", matrix, searchValue, &number, lastCoord, i, matrixHeight, numCoord, part2)
			if stop {
				break
			}
			if i != matrixHeight {
				stop := traverseMatrixDirection("up", matrix, searchValue, &number, lastCoord, i, matrixHeight, numCoord, part2)
				if stop {
					break
				}
			}
		} else {
			right = true
			stop := traverseMatrixDirection("left", matrix, searchValue, &number, lastCoord, i, matrixHeight, numCoord, part2)
			if stop {
				break
			}
			stop = traverseMatrixDirection("down", matrix, searchValue, &number, lastCoord, i, matrixHeight, numCoord, part2)
			if stop {
				break
			}
		}
	}

	if part2 {
		numCoord = lastCoord
	}
	return matrix, centerCoord, numCoord
}

// helper function to traverse the matrix in a given direction, filling out numbers along the way
func traverseMatrixDirection(direction string, matrix [][]int, searchValue int, number *int, lastCoord []int, i int, size int, numCoord []int, part2 bool) bool {
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

		if !part2 {
			matrix[x][y] = *number + 1
			*number++

			if *number == searchValue {
				numCoord[0] = lastCoord[0]
				numCoord[1] = lastCoord[1]
			}
		} else {
			matrix[x][y] = neighborSum(matrix, x, y)
			if matrix[x][y] > searchValue {
				// lastCoord contains the position of the first number that's greater than the input value, return
				return true
			}
		}
	}
	return false
}

/*
	Part2:
	As a stress test on the system, the programs here clear the grid and then store the value 1 in square 1. Then, in the same allocation order as shown above, they store the sum of the values in all adjacent squares, including diagonals.

	So, the first few squares' values are chosen as follows:

	Square 1 starts with the value 1.
	Square 2 has only one adjacent filled square (with value 1), so it also stores 1.
	Square 3 has both of the above squares as neighbors and stores the sum of their values, 2.
	Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4.
	Square 5 only has the first and fourth squares as neighbors, so it gets the value 5.
	Once a square is written, its value does not change. Therefore, the first few squares would receive the following values:

	147  142  133  122   59
	304    5    4    2   57
	330   10    1    1   54
	351   11   23   25   26
	362  747  806--->   ...
	What is the first value written that is larger than your puzzle input?
*/
func calcStepsPart2(num int) int {
	// TOOD: calculate matrix size based on input number (should be the next odd perfect square that is greater than the input number)
	matrixSize := 363609
	matrix, _, numCoord := generateSpiralMatrix(matrixSize, num, true)
	firstLargerValue := matrix[numCoord[0]][numCoord[1]]
	return firstLargerValue
}

// helper function to get the sum of all filled-out neighbors for a given coordinate
func neighborSum(matrix [][]int, x int, y int) int {
	matrixSize := len(matrix)
	sum := 0
	if y < matrixSize-1 {
		right := matrix[x][y+1]
		if right != 0 {
			sum += right
		}
	}
	if x > 0 && y < matrixSize-1 {
		rightUp := matrix[x-1][y+1]
		if rightUp != 0 {
			sum += rightUp
		}
	}
	if x > 0 {
		up := matrix[x-1][y]
		if up != 0 {
			sum += up
		}
	}
	if x > 0 && y > 0 {
		leftUp := matrix[x-1][y-1]
		if leftUp != 0 {
			sum += leftUp
		}
	}
	if y > 0 {
		left := matrix[x][y-1]
		if left != 0 {
			sum += left
		}
	}
	if y > 0 && x < matrixSize-1 {
		leftDown := matrix[x+1][y-1]
		if leftDown != 0 {
			sum += leftDown
		}
	}
	if x < matrixSize-1 {
		down := matrix[x+1][y]
		if down != 0 {
			sum += down
		}
	}
	if y < matrixSize-1 && x < matrixSize-1 {
		rightDown := matrix[x+1][y+1]
		if rightDown != 0 {
			sum += rightDown
		}
	}
	return sum
}
