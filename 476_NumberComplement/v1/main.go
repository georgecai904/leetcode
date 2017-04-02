package main

import (
	"fmt"
	"math"
)

//Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.
//
//Note:
//The given integer is guaranteed to fit within the range of a 32-bit signed integer.
//You could assume no leading zero bit in the integer’s binary representation.

func findComplement(num int) int {
	var power float64
	result := 0
	for {
		if num % 2 == 0{
			result += int(math.Pow(2, power))
		}
		power += 1
		num /= 2
		if num == 0{
			break
		}
	}
	return result
}

func main() {
	fmt.Println(findComplement(9))
}