package main

import "fmt"

func main() {

	res := ArrayDiff([]int{1, 2, 3}, []int{1, 2})
	fmt.Printf("%#v\n", res)

	res = ArrayDiff([]int{1, 2, 2, 2, 3}, []int{2})
	fmt.Printf("%#v\n", res)

	res = ArrayDiff([]int{1, 2, 2}, []int{1})
	fmt.Printf("%#v\n", res)

}

func ArrayDiff(a, b []int) []int {

	res := make([]int, 0)

out:
	for _, x := range a {
		for _, y := range b {
			if x == y {
				continue out
			}
		}
		res = append(res, x)
	}

	return res
}
