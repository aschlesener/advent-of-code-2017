package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
	Author: Amy Schlesener - github.com/aschlesener
	Advent of Code Day 6 - http://adventofcode.com/2017/day/6
*/

func main() {
	// get memory banks from file
	banks := getBanks()

	// calculate how many redistribution cycles it takes to find an existing configuration
	numCyclesPart1 := countNumCyclesPart1(banks)
	numCyclesPart2 := countNumCyclesPart2(banks)
	fmt.Println("Number of cycles to redistribute for part 1 is:", numCyclesPart1)
	fmt.Println("Number of cycles to redistribute for part 2 is:", numCyclesPart2)
}

// helper function to parse text file containing list of integer instructions
func getBanks() []int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var row string

	// read file into line
	for scanner.Scan() {
		row = scanner.Text()
	}

	// loop through line in file to get list of space-delimited ints
	blockStrs := strings.Split(row, " ")
	banks := make([]int, len(blockStrs))
	for i, block := range blockStrs {
		banks[i], _ = strconv.Atoi(block)
	}

	return banks
}

/*
	Part 1 Rules:
	A debugger program here is having an issue: it is trying to repair a memory reallocation routine, but it keeps getting stuck in an infinite loop.

	In this area, there are sixteen memory banks; each memory bank can hold any number of blocks. The goal of the reallocation routine is to balance the blocks between the memory banks.

	The reallocation routine operates in cycles. In each cycle, it finds the memory bank with the most blocks (ties won by the lowest-numbered memory bank) and redistributes those blocks among the banks. To do this, it removes all of the blocks from the selected bank, then moves to the next (by index) memory bank and inserts one of the blocks. It continues doing this until it runs out of blocks; if it reaches the last memory bank, it wraps around to the first one.

	The debugger would like to know how many redistributions can be done before a blocks-in-banks configuration is produced that has been seen before.

	For example, imagine a scenario with only four memory banks:

	The banks start with 0, 2, 7, and 0 blocks. The third bank has the most blocks, so it is chosen for redistribution.
	Starting with the next bank (the fourth bank) and then continuing to the first bank, the second bank, and so on, the 7 blocks are spread out over the memory banks. The fourth, first, and second banks get two blocks each, and the third bank gets one back. The final result looks like this: 2 4 1 2.
	Next, the second bank is chosen because it contains the most blocks (four). Because there are four memory banks, each gets one block. The result is: 3 1 2 3.
	Now, there is a tie between the first and fourth memory banks, both of which have three blocks. The first bank wins the tie, and its three blocks are distributed evenly over the other three banks, leaving it with none: 0 2 3 4.
	The fourth bank is chosen, and its four blocks are distributed such that each of the four banks receives one: 1 3 4 1.
	The third bank is chosen, and the same thing happens: 2 4 1 2.
	At this point, we've reached a state we've seen before: 2 4 1 2 was already seen. The infinite loop is detected after the fifth block redistribution cycle, and so the answer in this example is 5.

	Given the initial block counts in your puzzle input, how many redistribution cycles must be completed before a configuration is produced that has been seen before?
*/
func countNumCyclesPart1(banks []int) int {
	// TODO: this algorithm can be optimized to be faster
	numCycles := 0
	bankPosition := 0
	configurationSeen := false
	configurations := make([][]int, 0)
	// keep redistributing blocks until we reach a configuration we've seen before (indicates infinite loop)
	for !configurationSeen {
		numCycles++
		max := 0
		bankPosition = 0
		// get bank with highest number of blocks
		for i, blocks := range banks {
			if blocks > max {
				bankPosition = i
				max = blocks
			}
		}
		numBlocksRemaining := banks[bankPosition]

		// remove all blocks from highest bank
		banks[bankPosition] = 0

		// redistribute blocks among other banks
		for numBlocksRemaining > 0 {
			// increment position, wrapping around to beginning if needed
			if bankPosition == len(banks)-1 {
				bankPosition = 0
			} else {
				bankPosition++
			}
			banks[bankPosition]++
			numBlocksRemaining--
		}
		// store this configuration if it doesn't exist, otherwise quit
		for _, config := range configurations {
			if reflect.DeepEqual(config, banks) {
				configurationSeen = true
				break
			}
		}
		banksCopy := make([]int, len(banks))
		copy(banksCopy, banks)
		configurations = append(configurations, banksCopy)
	}

	return numCycles
}

/*
	Part 2 Rules:
	Out of curiosity, the debugger would also like to know the size of the loop: starting from a state that has already been seen, how many block redistribution cycles must be performed before that same state is seen again?

	In the example above, 2 4 1 2 is seen again after four cycles, and so the answer in that example would be 4.

	How many cycles are in the infinite loop that arises from the configuration in your puzzle input?
*/
func countNumCyclesPart2(banks []int) int {
	numCycles := 0
	bankPosition := 0
	configurationSeen := false
	configurationSeenTwice := false
	configurations := make([][]int, 0)
	seenConfig := make([]int, 1)
	// keep redistributing blocks until we reach a configuration we've seen before (indicates infinite loop)
	for !configurationSeenTwice {
		if configurationSeen {
			numCycles++
		}
		max := 0
		bankPosition = 0
		// get bank with highest number of blocks
		for i, blocks := range banks {
			if blocks > max {
				bankPosition = i
				max = blocks
			}
		}
		numBlocksRemaining := banks[bankPosition]

		// remove all blocks from highest bank
		banks[bankPosition] = 0

		// redistribute blocks among other banks
		for numBlocksRemaining > 0 {
			// increment position, wrapping around to beginning if needed
			if bankPosition == len(banks)-1 {
				bankPosition = 0
			} else {
				bankPosition++
			}
			banks[bankPosition]++
			numBlocksRemaining--
		}
		// store this configuration if it doesn't exist, otherwise quit
		if !configurationSeen {
			banksCopy := make([]int, len(banks))
			copy(banksCopy, banks)
			for _, config := range configurations {
				if reflect.DeepEqual(config, banks) {
					// store already-seen config
					configurationSeen = true
					seenConfig = banksCopy
				}
			}
			configurations = append(configurations, banksCopy)
		} else {
			if reflect.DeepEqual(banks, seenConfig) {
				configurationSeenTwice = true
			}
		}
	}

	return numCycles
}
