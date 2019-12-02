package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	program := readIn("intcode.csv")

	fmt.Println("Part 1:")

	program[1] = 12
	program[2] = 2

	for i := 0; i < len(program); i += 4 {
		fmt.Println("Position:", i, "Handling opcode:", program[i])
		switch program[i] {
		case 1:
			{
				//add
				result := program[program[i+1]] + program[program[i+2]]
				program[program[i+3]] = result
				continue
			}
		case 2:
			{
				//mult
				result := program[program[i+1]] * program[program[i+2]]
				program[program[i+3]] = result
				continue
			}
		case 99:
			{
				//end
				break
			}
		default:
			{
				//error
				fmt.Println("Program error! Found unexpected opcode:", program[i])
				break
			}
		}
	}

	fmt.Println("Value at position 0:", program[0])
	part2()
}

func readIn(path string) []int {
	//open file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	programText, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var program []int

	for _, text := range programText {
		val, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		program = append(program, val)
	}

	return program
}
