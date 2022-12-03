package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)
func sum(arr []int64) int64 {
	var sum int64
	for _, v := range arr {
		sum += v
	}

	return sum
}

func main () {
    input, err := os.ReadFile("input")
    // input, err := os.ReadFile("example-input")
    if err != nil {
        log.Fatal(err)
    }

	elfs := strings.Split(string(input), "\n\n")

	records := [][]int64{}

	for i, v := range elfs {
		cals := strings.Split(v, "\n")

		records = append(records, []int64{})

		for _, cal := range cals {
			if cal != "" {
				parsedValue, err := strconv.ParseInt(cal, 10, 0)
				if err != nil {
					log.Fatalln("Fuck")
				}
				records[i] = append(records[i], parsedValue)
			}
		}
	}

	sort.Slice(records, func(i, j int) bool {
		return sum(records[i]) > sum(records[j])
	})

	log.Println(records[0])
	log.Println(sum(records[0]))

	var top3 int64

	for i, cals := range records {
		if i < 3 {
			top3 += sum(cals)
		}
	}


	log.Println("Sum top 3:")
	log.Println(top3)
}
