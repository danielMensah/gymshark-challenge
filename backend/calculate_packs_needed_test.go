package main

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

var packs = []int{250, 500, 1000, 2000, 5000}

func init() {
	sort.Slice(packs, func(i, j int) bool {
		return packs[i] > packs[j]
	})
}

func TestOrder1(t *testing.T) {
	var res []int
	expectedResult := "250"

	actualResult := calculatePacksNeeded(1, packs, res)

	if expectedResult != convert(actualResult) {
		t.Errorf("Order 1 failed, expected %v, got %v", expectedResult, actualResult)
	}
}

func TestOrder250(t *testing.T) {
	var res []int
	expectedResult := "250"

	actualResult := calculatePacksNeeded(250, packs, res)

	if expectedResult != convert(actualResult) {
		t.Errorf("Order 250 failed, expected %v, got %v", expectedResult, actualResult)
	}
}

func TestOrder251(t *testing.T) {
	var res []int
	expectedResult := "500"

	actualResult := calculatePacksNeeded(251, packs, res)

	if expectedResult != convert(actualResult) {
		t.Errorf("Order 251 failed, expected %v, got %v", expectedResult, actualResult)
	}
}

func TestOrder501(t *testing.T) {
	var res []int
	expectedResult := "500, 250"

	actualResult := calculatePacksNeeded(501, packs, res)

	if expectedResult != convert(actualResult) {
		t.Errorf("Order 501 failed, expected %v, got %v", expectedResult, actualResult)
	}
}

func TestOrder12001(t *testing.T) {
	var res []int
	expectedResult := "5000, 5000, 2000, 250"

	actualResult := calculatePacksNeeded(12001, packs, res)

	if expectedResult != convert(actualResult) {
		t.Errorf("Order 12001 failed, expected %v, got %v", expectedResult, actualResult)
	}
}

func TestOrder160(t *testing.T) {
	var res []int
	newPacks := packs
	newPacks = append(newPacks, 50)
	expectedResult := "50, 50, 50, 50"

	actualResult := calculatePacksNeeded(160, newPacks, res)

	if expectedResult != convert(actualResult) {
		t.Errorf("Order 160 failed, expected %v but got %v", expectedResult, actualResult)
	}
}
func TestOrder239(t *testing.T) {
	var res []int
	newPacks := packs
	newPacks = append(newPacks, 110)
	expectedResult := "250"

	actualResult := calculatePacksNeeded(239, newPacks, res)

	if expectedResult != convert(actualResult) {
		t.Errorf("Order 239 failed, expected %v but got %v", expectedResult, actualResult)
	}
}

func convert(array []int) string {
	var num []string
	for _, i := range array {
		num = append(num, strconv.Itoa(i))
	}

	return strings.Join(num, ", ")
}
