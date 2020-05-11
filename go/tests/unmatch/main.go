package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
)

// Returns the element of the slice that doesn't have a correspondant pair
func unmatch(arr []int) int {
	var quant int
	for _, el := range arr {
		quant = 0
		for _, v := range arr {
			if v == el {
				quant++
			}
		}
		if quant%2 != 0 {
			return el
		}
	}
	return -1
}

func main() {
	arg1 := []int{1, 1, 2, 3, 4, 3, 4}
	arg2 := []int{1, 1, 2, 4, 3, 4, 2, 3, 4}
	arg3 := []int{1, 2, 1, 1, 4, 5, 5, 4, 1, 7}
	arg4 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arg5 := []int{0, 20, 91, 23, 10, 34}
	arg6 := []int{1, 1, 2, 2, 3, 4, 3, 4, 5, 5, 8, 9, 8, 9}

	randInt1 := lib.RandIntBetween(-100, 100)
	randInt2 := lib.RandIntBetween(-1000, 1000)
	randInt3 := lib.RandIntBetween(-10, 10)

	arg7 := []int{randInt1, randInt2, randInt1, randInt2, randInt1 + randInt3, randInt1 + randInt3}
	arg8 := []int{randInt1, randInt2, randInt1, randInt2, randInt1 + randInt3, randInt2 - randInt3}

	args := [][]int{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8}

	for _, v := range args {
		lib.Challenge("Unmatch", student.Unmatch, unmatch, v)
	}
}
