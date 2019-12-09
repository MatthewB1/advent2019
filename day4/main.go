package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lowBound = 234208
var highBound = 765869

func main() {

	validPasswords := 0

	for i := lowBound; i <= highBound; i++ {
		intSlice := toIntSlice(i)
		if !ascending(intSlice) {
			continue
		}
		if !hasAdjacentRepeat(intSlice) {
			continue
		}
		validPasswords++
	}
	fmt.Println("Valid passwords:", validPasswords)

	part2()
}

func toIntSlice(in int) []int {
	numberString := strconv.Itoa(in)
	stringSlice := strings.Split(numberString, "")
	var intSlice = []int{}

	for _, val := range stringSlice {
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func ascending(in []int) bool {
	for i, val := range in {
		if i > 0 {
			if val < in[i-1] {
				return false
			}
		}
	}
	return true
}

func hasAdjacentRepeat(in []int) bool {
	for i, val := range in {
		if i > 0 {
			if val == in[i-1] {
				return true
			}
		}
	}
	return false
}

func part2() {

	validPasswords := 0

	for i := lowBound; i <= highBound; i++ {
		intSlice := toIntSlice(i)
		if !ascending(intSlice) {
			continue
		}
		if !hasMaxTwoAdjacentRepeat(intSlice) {
			continue
		}
		validPasswords++
	}
	fmt.Println("Valid passwords:", validPasswords)
}

func hasMaxTwoAdjacentRepeat(in []int) bool {
	for i := 1; i < len(in); i++ {
		if in[i] == in[i-1] {
			//check for larger group
			if i+1 < len(in) {
				if in[i] == in[i+1] {
					continue
				}
			}
			if i-2 >= 0 {
				if in[i] == in[i-2] {
					continue
				}
			}
			return true
		}
	}
	return false
}
