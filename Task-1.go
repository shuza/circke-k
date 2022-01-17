package main

import (
	"strconv"
	"strings"
)

//	if there is an empty parts that means the syntax is not correct
func testValidity(input string) bool {
	parts := strings.Split(input, "-")
	for _, v := range parts {
		if len(v) == 0 {
			//	syntax not correct
			return false
		}
		if v[0] >= '0' && v[0] <= '9' {
			if !isValidNumber(v) {
				return false
			}
		} else {
			if !isValidWord(v) {
				return false
			}
		}
	}
	return true
}

//	it will  parse the string and return average of all number
func averageNumber(input string) int {
	parts := strings.Split(input, "-")
	sum := 0
	count := 0
	for _, v := range parts {
		num, err := strconv.Atoi(v)
		if err == nil {
			sum = sum + num
			count++
		}
	}
	avg := sum / count
	return avg
}

//	it will parse input string and join all valid world with space separate
func wholeStory(input string) string {
	parts := strings.Split(input, "-")
	result := ""
	for _, v := range parts {
		if isValidWord(v) {
			if len(result) > 0 {
				result = result + " "
			}
			result = result + v
		}
	}
	return result
}

func storyStats(input string) (string, string, []string) {
	parts := strings.Split(input, "-")
	words := make([]string, 0)
	sumOfLength := 0
	for _, v := range parts {
		if isValidWord(v) {
			words = append(words, v)
			sumOfLength = sumOfLength + len(v)
		}
	}
	//	no valid word
	if sumOfLength == 0 {
		return "", "", []string{}
	}
	avgLength := sumOfLength / len(words)
	if sumOfLength%avgLength > 0 {
		avgLength++
	}
	if len(words) == 1 {
		return words[0], words[0], words
	}
	shortest, largest := words[0], words[0]
	avgWord := make([]string, 0)
	for _, v := range words {
		if len(v) < len(shortest) {
			shortest = v
		}
		if len(v) > len(largest) {
			largest = v
		}
		if len(v) == avgLength {
			avgWord = append(avgWord, v)
		}
	}

	return shortest, largest, avgWord
}

//	an input will be valid number if it only contains 0-9
func isValidNumber(input string) bool {
	for _, d := range input {
		if d < '0' || d > '9' {
			return false
		}
	}
	return true
}

//	an input will be valid if it only contains alphabits
func isValidWord(input string) bool {
	for _, c := range input {
		if c < 'A' || c > 'z' {
			return false
		}
	}
	return true
}
