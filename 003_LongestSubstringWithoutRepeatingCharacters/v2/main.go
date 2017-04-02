package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var start, end, maxLength int
	var temp string
	for i, c := range s {
		var char = string(c)
		end = i
		if strings.Contains(temp, char) {
			start = strings.LastIndex(s[:i], char) + 1
		}

		temp = s[start : end+1]
		if len(temp) > maxLength {
			maxLength = len(temp)
		}
	}
	return maxLength
}
func main() {
	val := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(val))
}
