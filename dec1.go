package main

import "strconv"

func dec1() {
	lines, _ := readLines("dec1.txt")

	inv := readInventory(lines)

	dec1_1(inv)
	dec1_2(inv)
}

type inventory struct {
	totalCals int
	cals      []int
}

func dec1_1(inv []inventory) {
	println("dec1_1", inv[maxIndex(inv)].totalCals)
}

func dec1_2(inv []inventory) {
	var maxs []inventory

	for i := 0; i < 3; i++ {
		maxIndex := maxIndex(inv)
		maxs = append(maxs, inv[maxIndex])
		inv = append(inv[:maxIndex], inv[maxIndex+1:]...)
	}

	var total int
	for _, i := range maxs {
		total = total + i.totalCals
	}

	println("dec1_2", total)
}

func maxIndex(inv []inventory) int {
	var maxCals int
	var maxIndex int

	for i, x := range inv {
		if x.totalCals > maxCals {
			maxIndex = i
			maxCals = x.totalCals
		}
	}
	return maxIndex
}

func readInventory(lines []string) []inventory {
	var inv []inventory
	var i inventory
	for _, line := range lines {
		if line == "" {
			inv = append(inv, i)
			i = inventory{}
		} else {
			cals, _ := strconv.Atoi(line)
			i.totalCals = i.totalCals + cals
			i.cals = append(i.cals, cals)
		}
	}
	return inv
}
