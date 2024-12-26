package main

import (
	"fmt"
	"os"
	"strconv"
)

// Day 3: Mull It Over
// https://adventofcode.com/2024/day/3

const inputFile = "input.txt"

func main() {
	total, err := part1(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", total) // 173517243
	total2, err := part2(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2:", total2) // 100450138

}

func part1(input string) (int, error) {
	input, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}
	total := 0
	for i := range input {
		if i+4 < len(input) && input[i:i+4] == "mul(" {
			res, reached := parseMul(input, i)
			i = reached
			total += res
		}
	}
	return total, nil
}

// part2 is the same as part1 but with turnon/turnoff logic
func part2(input string) (int, error) {
	input, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}
	enabled := true
	total := 0
	for i := range input {
		if i+4 < len(input) && input[i:i+4] == "do()" {
			enabled = true
			i += 3
		}
		if i+7 < len(input) && input[i:i+7] == "don't()" {
			enabled = false
			i += 7
		}
		if enabled && i+4 < len(input) && input[i:i+4] == "mul(" {
			res, reached := parseMul(input, i)
			i = reached
			total += res
		}
	}
	return total, nil
}

// parseMul parses the "mul(left, right)" string and returns the result of the multiplication and the index of next character
// offset is the index of the first character of the "mul(" string
func parseMul(input string, offset int) (int, int) {
	left, right, err := 0, 0, error(nil)
	comma := -1           // comma index
	reached := offset + 4 // reached index of j
	for j := offset + 4; j < min(len(input), offset+12); j++ {
		if input[j] == ',' {
			comma = j
			left, err = strconv.Atoi(input[offset+4 : j])
			if err != nil {
				break
			}
		}
		if comma > -1 && input[j] == ')' {
			right, err = strconv.Atoi(input[comma+1 : j])
			if err != nil {
				break
			}
			break
		}
		reached = j
	}
	return left * right, reached + 1

}

func parseInput(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
