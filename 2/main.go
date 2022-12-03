package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// input, err := os.ReadFile("input")
	input, err := os.ReadFile("example-input")
	if err != nil {
		log.Fatal(err)
	}

	roundsString := strings.Split(string(input), "\n")

	rounds := [][]int{}

	for _, round := range roundsString {
		if round == "" {
			continue
		}

		ss := strings.Split(round, " ")

		round := []int{}
		for _, val := range ss {
			if val == "A" || val == "X" {
				round = append(round, 1)
			}
			if val == "B" || val == "Y" {
				round = append(round, 2)
			}
			if val == "C" || val == "Z" {
				round = append(round, 3)
			}
		}

		rounds = append(rounds, round)
	}

	round1(rounds)
	round2(rounds)
}

func round1(rounds [][]int) {

	points := 0
	for _, round := range rounds {
		points += round[1]

		// 1 > 3
		// 2 > 1
		// 3 > 2

		if round[1] == round[0] {
			points += 3
		}

		if round[1]-1 == round[0] {
			points += 6
		}

		if round[1] == 1 && round[0] == 3 {
			points += 6
		}
	}

	log.Println("Ponts:", points)
}

func round2(rounds [][]int) {
	win := map[int]int{
		1: 3,
		2: 1,
		3: 2,
	}
	lose := map[int]int{
		3: 1,
		1: 2,
		2: 3,
	}

	points := 0
	for _, round := range rounds {

		// 1 > 3
		// 2 > 1
		// 3 > 2

		if round[1] == 1 {
			points += lose[round[1]]
		}
		if round[1] == 2 {
			points += round[1]
			points += 3
		}
		if round[1] == 3 {
			points += win[round[1]]
			points += 6
		}
	}

	log.Println("Points:", points)
}
