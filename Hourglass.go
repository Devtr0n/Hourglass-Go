package main

import (
	"fmt"
)

func hourGlassSum(arr [][]int) int {

	// initialize the sum and largestSum, sum is used to temporarily store the hourglass sum of each calculation, largestSum is used to store the largest sum of all calculations
	sum := 0
	largestSum := 0

	// traverse the non-edge numbers of the array (that is, the number that can become the center of the hourglass)
	for j := 1; j < len(arr)-1; j++ {
		for i := 1; i < len(arr)-1; i++ {

			// sum of current hourglass numbers
			sum = arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i][j] + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]

			// determine the largest hourglass sum, a clever trick
			largestSum = max(largestSum, sum)
		}
	}

	return largestSum
}

func main() {

	// initialize a two-dimensional array
	arr := [][]int{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0}}

	fmt.Println(hourGlassSum(arr)) // output should be 19

	// initialize another two-dimensional array
	arr = [][]int{
		{1, 2, 3, 0, 0},
		{0, 0, 0, 0, 0},
		{2, 1, 4, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0}}

	fmt.Println(hourGlassSum(arr)) // output should be 13

}

// returns the larger of x or y
func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
