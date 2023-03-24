package main

import "fmt"

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
3,5,6,9,10,12,15,18,20,21
Find the sum of all the multiples of 3 or 5 below 1000.
*/

// func multiples(max int) int {

// 	// O(1)
// 	// sum 3 multiples and sum 5 multiples - 15 multiples(bc this already get's covered twice)
// 	// summation notation 1 + 2 + 3 + 4 ... n = n * (n-1)/2

// 	sum3 := 3 * ((max - 1) / 3 * ((max-1)/3 + 1) / 2)     // goes up from 1 - 333
// 	sum5 := 5 * ((max - 1) / 5 * ((max-1)/5 + 1) / 2)     // goes from 1 - 199
// 	sum15 := 15 * ((max - 1) / 15 * ((max-1)/15 + 1) / 2) // goes from 1 - 66

// 	return (sum3 + sum5 - sum15)
// }

// Generalize it more
func multiples(max int, multiples1 int, multiples2 int) int {

	// O(1)
	// sum 3 multiples and sum 5 multiples - 15 multiples(bc this already get's covered twice)
	// summation notation 1 + 2 + 3 + 4 ... n = n * (n-1)/2

	sum1 := multiples1 * ((max - 1) / multiples1 * ((max-1)/multiples1 + 1) / 2)
	sum2 := multiples2 * ((max - 1) / multiples2 * ((max-1)/multiples2 + 1) / 2)
	sumTot := multiples1 * multiples2 * ((max - 1) / (multiples1 * multiples2) * ((max-1)/(multiples1*multiples2) + 1) / 2) // goes from 1 - 66

	return (sum1 + sum2 - sumTot)
}

func main() {
	fmt.Println(multiples(1000, 3, 5))
}
