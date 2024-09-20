package main

import (
	"fmt"
)

func main() {
	var sequence string
	fmt.Scan(&sequence) // Read the sequence of pressed keys

	var required string
	fmt.Scan(&required) // Read the required characters

	var k int
	fmt.Scan(&k) // Read the maximum length of the password

	requiredCount := make(map[rune]int)
	for _, char := range required {
		requiredCount[char]++
	}

	// Sliding window variables
	left := 0
	formed := 0
	windowCount := make(map[rune]int)
	var bestStart, bestEnd int
	var found bool

	// Sliding window to find the rightmost valid substring
	for right, char := range sequence {
		if _, exists := requiredCount[char]; exists {
			windowCount[char]++
			if windowCount[char] == requiredCount[char] {
				formed++
			}
		}

		// Try to contract the window until it ceases to be 'desirable'
		for left <= right && formed == len(requiredCount) {
			// Record the current window if it is a valid password candidate
			if right-left+1 <= k {
				found = true
				bestStart = left
				bestEnd = right
			}

			// Now we try to move left pointer to find a smaller window
			if _, exists := requiredCount[rune(sequence[left])]; exists {
				windowCount[rune(sequence[left])]--
				if windowCount[rune(sequence[left])] < requiredCount[rune(sequence[left])] {
					formed--
				}
			}
			left++
		}
	}

	// If we found a valid substring, print it
	if found {
		fmt.Println(sequence[bestStart : bestEnd+k/2])
	} else {
		fmt.Println("-1")
	}
}
