package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	Author: Amy Schlesener - github.com/aschlesener
	Advent of Code Day 4 - http://adventofcode.com/2017/day/4
*/

func main() {
	// get number of square from file
	passphrases := getPassphrases()

	// calculate how many passphrases are valid
	numValidPart1 := countValidPassphrasesPart1(passphrases)
	fmt.Println("Number of valid passphrases for part 1 is:", numValidPart1)
}

// helper function to parse text file containing list of passphrases
func getPassphrases() [][]string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var fileLines []string

	// read file into array of lines
	for scanner.Scan() {
		row := scanner.Text()
		fileLines = append(fileLines, row)
	}
	passphrases := make([][]string, len(fileLines))

	// loop through each line in file to get list of space-delimited phrases
	for i, line := range fileLines {
		var row []string
		phrases := strings.Split(line, " ")

		for _, phrase := range phrases {
			row = append(row, phrase)
		}
		passphrases[i] = append(passphrases[i], row...)
	}
	return passphrases
}

/*
	Part 1 Rules:
	A new system policy has been put in place that requires all accounts to use a passphrase instead of simply a password. A passphrase consists of a series of words (lowercase letters) separated by spaces.

	To ensure security, a valid passphrase must contain no duplicate words.

	For example:

	aa bb cc dd ee is valid.
	aa bb cc dd aa is not valid - the word aa appears more than once.
	aa bb cc dd aaa is valid - aa and aaa count as different words.
	The system's full passphrase list is available as your puzzle input. How many passphrases are valid?
*/
func countValidPassphrasesPart1(passphrases [][]string) int {
	numValid := 0
	for _, passphraseSet := range passphrases {
		// for each passphrase set, use map to keep track of which words are already in set
		passMap := make(map[string]int)
		valid := true
		for _, passphrase := range passphraseSet {
			if passMap[passphrase] == 0 {
				passMap[passphrase] = 1
			} else {
				// duplicated words in passphrase, not valid
				valid = false
			}
		}
		if valid {
			numValid++
		}
	}
	return numValid
}
