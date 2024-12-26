package main

import (
	"fmt"
	"os"
	"strings"
)

// Day 4: Ceres Search
// https://adventofcode.com/2024/day/4

const inputFile = "input.txt"
const keyword = "XMAS"

var directs = map[int][2]int{
	0: {0, 1},
	1: {1, 0},
	2: {0, -1},
	3: {-1, 0},
	4: {1, 1},
	5: {1, -1},
	6: {-1, -1},
	7: {-1, 1},
}

func main() {
	total, err := part1(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", total) // 2493
	total2, err := part2(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2:", total2) // 1890
}

func part1(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}
	total := 0
	for i := range matrix {
		for j := range matrix[i] {
			for dir := range directs {
				res := search1(matrix, i, j, dir, 0)
				total += res
			}
		}
	}
	return total, nil
}

func search1(matrix [][]rune, i, j, dir, idx int) int {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return 0
	}

	if matrix[i][j] != []rune(keyword)[idx] {
		return 0
	}
	if idx == len(keyword)-1 {
		return 1
	}
	return search1(matrix, i+directs[dir][0], j+directs[dir][1], dir, idx+1)
}

func part2(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}
	total := 0
	for i := range matrix {
		for j := range matrix[i] {
			res := search2(matrix, i, j)
			total += res
		}
	}
	return total, nil
}

func search2(matrix [][]rune, i, j int) int {
	if i < 1 || j < 1 || i >= len(matrix)-1 || j >= len(matrix[0])-1 {
		return 0
	}
	if matrix[i][j] != 'A' {
		return 0
	}
	if !((matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S') || (matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M')) {
		return 0
	}
	if (matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S') || (matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M') {
		return 1
	}
	return 0
}

func parseInput(fileName string) ([][]rune, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	res := make([][]rune, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		res = append(res, []rune(line))
	}
	return res, nil
}
