package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxLayerSize = 25 * 6

func main() {
	//open file
	file, err := os.Open("image.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	//read in
	scanner := bufio.NewScanner(file)
	var text string

	for scanner.Scan() {
		text = scanner.Text()
	}

	//split to chars
	chars := strings.Split(text, "")

	var layers [][]int
	var layer []int

	//build layers
	for _, char := range chars {
		charint, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		layer = append(layer, charint)

		if len(layer) == maxLayerSize {
			layers = append(layers, layer)
			layer = nil
		}
	}

	layerIndex := layerWithFewestZeroes(layers)

	fmt.Println(onesByTwos(layers[layerIndex]))

}

func countZeroes(layer []int) int {
	zeroes := 0

	for _, val := range layer {
		if val == 0 {
			zeroes++
		}
	}
	return zeroes
}

func layerWithFewestZeroes(layers [][]int) int {
	index := 0

	for i, layer := range layers {
		if countZeroes(layer) < countZeroes(layers[index]) {
			index = i
		}
	}

	return index
}

func onesByTwos(layer []int) int {
	ones := 0
	twos := 0

	for _, val := range layer {
		if val == 1 {
			ones++
		}
		if val == 2 {
			twos++
		}
	}

	return ones * twos
}
