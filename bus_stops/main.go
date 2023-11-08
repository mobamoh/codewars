package main

import "fmt"

func main() {

	fmt.Printf("%#v\n", Number([][2]int{{10, 0}, {3, 5}, {5, 8}}))
	fmt.Printf("%#v\n", Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}))
	fmt.Printf("%#v\n", Number([][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}))
	fmt.Printf("%#v\n", Number([][2]int{{0, 0}}))

}

func Number(busStops [][2]int) int {
	peopleInTheBus := 0
	for _, stops := range busStops {
		peopleInTheBus += stops[0]
		peopleInTheBus -= stops[1]
	}
	return peopleInTheBus
}
