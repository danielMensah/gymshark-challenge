package main

import (
	"fmt"
	"math"
)

func main() {
	packs := []int{250, 500, 1000, 2000, 5000}
	var res []int
	quantity := 12001

	result := order(quantity, packs, res)

	fmt.Printf("%v\n", result)
}


func order (quantity int, packs []int, res []int) []int {
	if contains(quantity, packs) {
		res = append(res, quantity)
	} else {
		a := getNearest(quantity, packs)
		b := quantity - a

		res = append(res, a)

		if b > 0 {
			return order(b, packs, res)
		}
	}

	sumRes := sum(res)

	if contains(sumRes, packs) {
		res = []int{sumRes}
	}

	return res

}

func getNearest(quantity int, packs []int) int {
	var nearest = packs[0]

	for _, pack := range packs {
		a := math.Abs(float64(quantity - pack))
		b := math.Abs(float64(quantity - nearest))

		if a < b {
			nearest = pack
		}
	}

	return nearest
}

func contains(item int, arr []int) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}