package main

import (
	"log"
	"math"
	"os"
	"strings"
)

var priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func toPriority(s string) int {
	return strings.Index(priorities, s) + 1
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

type Rucksacks [][][]int

func main() {
	// input, err := os.ReadFile("input")
	input, err := os.ReadFile("example-input")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(input), "\n")

	round1(data)
	round2(data)
}

func checkPairs(pairs [][]int) int {
	for _, priority := range pairs[0] {
		if contains(pairs[1], priority) {
			return priority
		}
	}

	return 0
}

func round1(data []string) {
	rucksacks := Rucksacks{}
	errorsSum := 0

	for a, line := range data {
		if line == "" {
			continue
		}
		rucksacks = append(rucksacks, make([][]int, 2))

		for i, item := range line {
			if len(line)/2 > i {
				rucksacks[a][0] = append(rucksacks[a][0], toPriority(string(item)))
			} else {
				rucksacks[a][1] = append(rucksacks[a][1], toPriority(string(item)))
			}
		}

		errorsSum += checkPairs(rucksacks[a])
	}

	log.Println(errorsSum)
}

func round2(lines []string) {
	rucksacks := Rucksacks{}

	for a, line := range lines {
		if line == "" {
			continue
		}
		if (a+1)%3 == 1 {
			rucksacks = append(rucksacks, make([][]int, 3))
		}

		group := int(math.Ceil(float64(a+1)/3.0)) - 1
		elf := (a + 1) % 3

		for _, item := range line {
			rucksacks[group][elf] = append(rucksacks[group][elf], toPriority(string(item)))
		}
	}

	sum := 0
	for _, group := range rucksacks {
		for _, priority := range group[0] {
			if contains(group[1], priority) && contains(group[2], priority) {
				sum += priority
				break
			}
		}
	}

	log.Println(sum)
}
