package main

import "math"

// original quantity
var quantitySnapshot int

func calculatePacksNeeded(quantity int, packs []int, res []int) []int {
	if len(res) == 0 {
		quantitySnapshot = quantity
	}

	remaining := quantity

	if contains(quantity, packs) {
		res = append(res, quantity)
	} else {
		for _, pack := range packs {
			// if pack fits the remaining quantity, add it to res (result) and recursively call the calculate function
			if remaining >= pack {
				res = append(res, pack)
				return calculatePacksNeeded(remaining - pack, packs, res)
			}
		}

		// if the remaining quantity is less than the smallest pack, add the smallest pack to res (result)
		// else if there's still some quantity, then recursively call the calculate function.
		// Note: packs are already sorted by descending.
		if remaining < packs[len(packs) - 1] {
			res = append(res, packs[len(packs) - 1])
		} else if remaining > 0 {
			return calculatePacksNeeded(remaining, packs, res)
		}
	}

	sumRes := sum(res)
	nearest := getNearest(quantitySnapshot, packs) // getting the nearest pack number to the original quantity

	// if there's a pack which is equals to the sum of res, then use the pack
	// else if the nearest pack to the quantity is the best choice, then use it
	if contains(sumRes, packs) {
		res = []int{sumRes}
	} else if nearest < sumRes && nearest > quantitySnapshot {
		res = []int{nearest}
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
