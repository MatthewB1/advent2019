package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func readInputs(path string) [][]string {
	//open file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	//read in
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	inputs := [][]string{}

	//iterate over
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		inputs = append(inputs, input)
	}
	return inputs
}
