package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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

		fuel := int(mass/3) - 2

		total += fuel
	}

	fmt.Println("Total fuel required:", total)
	part2()
}
