package main

import "fmt"

//Reverse digits of an integer.
//The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.
//
//Example1: x = 123, return 321
//Example2: x = -123, return -321

func reverse(x int) int {

	var division = 10
	var result int
	for x != 0 {
		result = result*10 + (x % division)
		x /= division
	}
	if x > 2147483647 || result > 2147483647 || result < -2147483648 || x < -2147483648 {
		return 0
	}
	return result
}

func main() {
	x := 123
	fmt.Println(reverse(x))
}
