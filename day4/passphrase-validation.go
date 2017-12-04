package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	numValidPart2 := countValidPassphrasesPart2(passphrases)
	fmt.Println("Number of valid passphrases for part 1 is:", numValidPart1)
	fmt.Println("Number of valid passphrases for part 2 is:", numValidPart2)
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

/*
	Part 2 Rules:
	For added security, yet another system policy has been put in place. Now, a valid passphrase must contain no two words that are anagrams of each other - that is, a passphrase is invalid if any word's letters can be rearranged to form any other word in the passphrase.

	For example:

	abcde fghij is a valid passphrase.
	abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
	a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
	iiii oiii ooii oooi oooo is valid.
	oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.

	Under this new system policy, how many passphrases are valid?
*/
func countValidPassphrasesPart2(passphrases [][]string) int {
	numValid := 0
	for _, passphraseSet := range passphrases {
		// for each passphrase set, use map to keep track of which sorted words are already in set
		passMap := make(map[string]int)
		valid := true
		for _, passphrase := range passphraseSet {
			sortedPassphrase := sortStringByCharacter(passphrase)
			if passMap[sortedPassphrase] == 0 {
				passMap[sortedPassphrase] = 1
			} else {
				// duplicated anagrams in passphrase, not valid
				valid = false
			}
		}
		if valid {
			numValid++
		}
	}
	return numValid
}

// helper funtions for sorting charaters within a string
// https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
