package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

/*
	Author: Amy Schlesener - github.com/aschlesener
	Advent of Code Day 11 - http://adventofcode.com/2017/day/11
*/

type point struct {
	x int
	y int
	z int
}

func main() {
	// get directions
	directions := getDirections()

	// calulate shortest number of steps required
	stepsPart1 := calcSteps(directions, false)
	fmt.Println("Shortest steps needed for part 1 is:", stepsPart1)
	stepsPart2 := calcSteps(directions, true)
	fmt.Println("Longest steps taken for part 2 is:", stepsPart2)
}

// helper function to parse text file containing comma-separated list of string directions
func getDirections() []string {
	bytes, _ := ioutil.ReadFile("input.txt")
	line := string(bytes)
	directions := strings.Split(line, ",")
	return directions
}

/*
	Part 1 Rules:
	The hexagons ("hexes") in this grid are aligned such that adjacent hexes can be found to the north, northeast, southeast, south, southwest, and northwest:

	\ n  /
	nw +--+ ne
	/    \
	-+      +-
	\    /
	sw +--+ se
	/ s  \
	You have the path the child process took. Starting where he started, you need to determine the fewest number of steps required to reach him. (A "step" means to move from the hex you are in to any adjacent hex.)

	For example:

	ne,ne,ne is 3 steps away.
	ne,ne,sw,sw is 0 steps away (back where you started).
	ne,ne,s,s is 2 steps away (se,se).
	se,sw,se,sw,sw is 3 steps away (s,s,sw).

	Part 2 Rules:
	How many steps away is the furthest he ever got from his starting position?
*/
func calcSteps(directions []string, longest bool) int {
	var start point
	var end point
	longestDistance := 0

	// traverse through directions to find end point
	for _, direction := range directions {
		switch direction {
		case "n":
			{
				end.y++
				end.z--
			}
		case "ne":
			{
				end.x++
				end.z--
			}
		case "se":
			{
				end.x++
				end.y--
			}
		case "s":
			{
				end.y--
				end.z++
			}
		case "sw":
			{
				end.x--
				end.z++
			}
		case "nw":
			{
				end.x--
				end.y++
			}
		}

		// keep track of longest distance
		distance := int(math.Abs(float64(start.x-end.x))+math.Abs(float64(start.y-end.y))+math.Abs(float64(start.z-end.z))) / 2
		if distance > longestDistance {
			longestDistance = distance
		}
	}

	// calculate shortest distance between start and end points (Manhattan distance)
	shortestDistance := int(math.Abs(float64(start.x-end.x))+math.Abs(float64(start.y-end.y))+math.Abs(float64(start.z-end.z))) / 2

	if longest {
		return longestDistance
	}
	return shortestDistance
}
