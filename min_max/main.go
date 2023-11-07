package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%#v\n", MinMax([]int{2334454, 5}))
	fmt.Printf("%#v\n", MinMax([]int{-2334454, -5}))
	fmt.Printf("%#v\n", MinMax([]int{1, 1}))
	fmt.Printf("%#v\n", MinMax([]int{7, 33, 99, 1, 5, 2, 3, 2}))
}

func MinMax(arr []int) [2]int {
	min, max := arr[0], arr[0]
	for _, v := range arr[1:] {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return [2]int{min, max}
}
