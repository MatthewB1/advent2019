package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part2() {
	//open file
	file, err := os.Open("modules.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	//read in
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0

	//iterate over
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		totalFuel := 0
		input := mass

		for {

			fuel := calcFuel(input)

			if int(fuel) <= 0 {
				break
			}

			totalFuel += fuel
			input = fuel
		}

		total += totalFuel
	}

	fmt.Println("Total fuel required (considering weight of fuel):", total)
}

func calcFuel(input int) int {
	return int(input/3) - 2
}
