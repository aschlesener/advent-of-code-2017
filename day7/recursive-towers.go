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
	Advent of Code Day 7 - http://adventofcode.com/2017/day/7
*/

type Tower struct {
	weight             int
	name               string
	towersHoldingNames []string
}

func main() {
	// get towers
	towers := getTowers()

	// calculate name of bottom tower
	towerNamePart1 := calcNamePart1(towers)
	fmt.Println("Name of bottom tower for part 1 is:", towerNamePart1)

}

// helper function to parse text file containing list of towers
func getTowers() []Tower {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var fileLines []string

	// read file into array of lines
	for scanner.Scan() {
		row := scanner.Text()
		fileLines = append(fileLines, row)
	}
	towers := make([]Tower, len(fileLines))

	// loop through each line in file
	for i, line := range fileLines {
		var tower Tower
		var towerList []string
		if strings.Contains(line, "->") {
			// tower contains list of towers that it points to
			towerListStr := strings.Replace(strings.Split(line, "->")[1], " ", "", -1)
			towerList = strings.Split(towerListStr, ",")
			line = strings.Split(line, "->")[0]
		}

		// parse tower name and weight
		towerParts := strings.Split(line, " ")
		tower.name = towerParts[0]
		tower.towersHoldingNames = towerList
		tower.weight, _ = strconv.Atoi(strings.Replace(strings.Replace(towerParts[1], "(", "", 1), ")", "", 1))
		towers[i] = tower
	}
	return towers
}

/*
	Part 1 Rules:
	Wandering further through the circuits of the computer, you come upon a tower of programs that have gotten themselves into a bit of trouble. A recursive algorithm has gotten out of hand, and now they're balanced precariously in a large tower.

	One program at the bottom supports the entire tower. It's holding a large disc, and on the disc are balanced several more sub-towers. At the bottom of these sub-towers, standing on the bottom disc, are other programs, each holding their own disc, and so on. At the very tops of these sub-sub-sub-...-towers, many programs stand simply keeping the disc below them balanced but with no disc of their own.

	You offer to help, but first you need to understand the structure of these towers. You ask each program to yell out their name, their weight, and (if they're holding a disc) the names of the programs immediately above them balancing on that disc. You write this information down (your puzzle input). Unfortunately, in their panic, they don't do this in an orderly fashion; by the time you're done, you're not sure which program gave which information.

	For example, if your list is the following:

	pbga (66)
	xhth (57)
	ebii (61)
	havc (66)
	ktlj (57)
	fwft (72) -> ktlj, cntj, xhth
	qoyq (66)
	padx (45) -> pbga, havc, qoyq
	tknk (41) -> ugml, padx, fwft
	jptl (61)
	ugml (68) -> gyxo, ebii, jptl
	gyxo (61)
	cntj (57)
	...then you would be able to recreate the structure of the towers that looks like this:

					gyxo
				/
			ugml - ebii
		/      \
		|         jptl
		|
		|         pbga
		/        /
	tknk --- padx - havc
		\        \
		|         qoyq
		|
		|         ktlj
		\      /
			fwft - cntj
				\
					xhth
	In this example, tknk is at the bottom of the tower (the bottom program), and is holding up ugml, padx, and fwft. Those programs are, in turn, holding up other programs; in this example, none of those programs are holding up any other programs, and are all the tops of their own towers. (The actual tower balancing in front of you is much larger.)

	Before you're ready to help them, you need to make sure your information is correct. What is the name of the bottom program?
*/
func calcNamePart1(towers []Tower) string {
	// find bottom tower - the only tower that a) is holding towers and b) does not have any other tower pointing to it
	for _, tower := range towers {
		bottomTower := true
		if len(tower.towersHoldingNames) > 0 {
			for _, otherTower := range towers {
				if otherTower.name != tower.name {
					for _, holdingTowerName := range otherTower.towersHoldingNames {
						if holdingTowerName == tower.name {
							// the tower name is pointed to by another tower so it's not the bottom
							bottomTower = false
						}
					}
				}
			}
			if bottomTower {
				return tower.name
			}
		} else {
			bottomTower = false
		}

	}
	return "No bottom tower name found"
}
