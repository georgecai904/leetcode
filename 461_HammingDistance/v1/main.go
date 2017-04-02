package main

import (
	"fmt"
)

//The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

//Given two integers x and y, calculate the Hamming distance.

func hammingDistance(x int, y int) int {
	count := 0
	for{
		if x % 2 != y % 2{
			count++
		}
		y /= 2
		x /= 2
		if y == 0 && x == 0{
			break
		}
	}
	return count
}

func main() {
	x := 1
	y := 3
	fmt.Println(hammingDistance(x, y))
}
