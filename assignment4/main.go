package main

import (
	"fmt"
	"strings"
	"unicode"
)

func findMiddleIndex(s string, c rune) int {
	
	lowercaseStr := strings.ToLower(s)
	lowercaseChar := unicode.ToLower(c)


	count := strings.Count(lowercaseStr, string(lowercaseChar))

	if count == 0 {
	
		return -1
	}
	middleIndex := -1
	occurrenceIndex := 0
	for i, ch := range lowercaseStr {
		if unicode.ToLower(ch) == lowercaseChar {
			occurrenceIndex++
			if occurrenceIndex == count/2+1 {
				middleIndex = i
				break
			}
		}
	}

	return middleIndex
}

func main() {
	str := "This is a ample string"
	character := 'S'

	middleIndex := findMiddleIndex(str, character)
	if middleIndex == -1 {
		fmt.Printf("Character '%c' not found in the string.\n", character)
	} else {
		fmt.Printf("Middle index of '%c' in the string: %d\n", character, middleIndex)
	}
}
