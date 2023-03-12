package main

import (
	"fmt"
	"math"
)

func evenOrOdd(s []string) string {
	var sum int
	arr := []byte{}
	for i := 0; i < len(s); i++ {
		temp := []byte(s[i])
		arr = append(arr, temp...) // remember dots for appending slice into another slice
	}

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		sum += int(math.Pow(float64(arr[i]), 2))
	}

	if sum%2 != 0 {
		return "ODD"
	} else {
		return "EVEN"
	}
}

func main() {
	str := []string{"hi", "there"}
	fmt.Println(evenOrOdd(str))
}
