//Given a string, find the length of the longest substring without repeating characters.
package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	temp := ""
	max := 0
	for _, c := range s {
		char := string(c)
		if strings.Contains(temp, char) {
			if len(temp) > max {
				max = len(temp)
			}
			temp = strings.Split(temp, char)[1] + char

		} else {
			temp += char
		}
	}
	if len(temp) > max {
		return len(temp)
	} else {
		return max
	}
}

func main() {
	val := "dvdf"
	fmt.Println(lengthOfLongestSubstring(val))

}
