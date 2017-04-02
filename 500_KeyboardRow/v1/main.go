package main

import "fmt"

//Given a List of words, return the words that can be typed using
//letters of alphabet on only one row's of American keyboard like the image below.




func findWords(words []string) []string {
	var hashmap = map[rune]int{
		'a': 2, 'b': 1, 'c': 1, 'd': 2, 'e': 3, 'f': 2, 'g': 2, 'h': 2, 'i': 3, 'j': 2, 'k': 2, 'l': 2, 'm': 1,
		'n': 1, 'o': 3, 'p': 3, 'q': 3, 'r': 3, 's': 2, 't': 3, 'u': 3, 'v': 1, 'w': 3, 'x': 1, 'y': 3, 'z': 1,
		'A': 2, 'B': 1, 'C': 1, 'D': 2, 'E': 3, 'F': 2, 'G': 2, 'H': 2, 'I': 3, 'J': 2, 'K': 2, 'L': 2, 'M': 1,
		'N': 1, 'O': 3, 'P': 3, 'Q': 3, 'R': 3, 'S': 2, 'T': 3, 'U': 3, 'V': 1, 'W': 3, 'X': 1, 'Y': 3, 'Z': 1,
	}
	var result []string
	for _, val := range words{
		var row int
		suitable := true
		for _, char := range val{
			if row == 0{
				row = hashmap[char]
			}else if row != hashmap[char]{
				suitable = false
				break
			}
		}
		if suitable {
			result = append(result, val)
		}
	}
	return result
}

func main()  {
	input := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(input))
}