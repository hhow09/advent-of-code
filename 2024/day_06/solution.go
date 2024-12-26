package main

import (
	"fmt"
	"os"
	"strings"
)

// Day 6: Guard Gallivant
// https://adventofcode.com/2024/day/6

const inputFile = "input.txt"

var dirs = map[int][2]int{
	0: {-1, 0},
	1: {0, 1},
	2: {1, 0},
	3: {0, -1},
}

func main() {
	total, err := part1(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", total) // 5531

	total, err = part2(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2:", total) // 251
}

func part1(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}

	steps := walk(matrix)
	return steps, nil
}

func startIndex(matrix [][]rune) (int, int) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func walk(matrix [][]rune) int {
	i, j := startIndex(matrix)
	if i == -1 {
		panic("couldn't find the starting point")
	}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	visited[i][j] = true

	count := 1
	dirIdx := 0
	for true {
		nxti, nxtj := i+dirs[dirIdx][0], j+dirs[dirIdx][1]
		if nxti < 0 || nxtj < 0 || nxti >= len(matrix) || nxtj >= len(matrix[0]) {
			return count
		}
		if matrix[nxti][nxtj] != '#' {
			i, j = nxti, nxtj
			if !visited[i][j] {
				count++
				visited[i][j] = true
			}
		} else {
			dirIdx = (dirIdx + 1) % 4 // turn right
		}
	}
	return count
}

func part2(input string) (int, error) {
	matrix, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}

	res := findObstacles(matrix)
	return res, nil
}

func findObstacles(matrix [][]rune) int {
	i, j := startIndex(matrix)
	if i == -1 {
		panic("couldn't find the starting point")
	}

	putObs := make([][]bool, len(matrix))
	for i := range putObs {
		putObs[i] = make([]bool, len(matrix[0]))
	}

	obsCount := 0
	dirIdx := 0
	for true {
		nxti, nxtj := i+dirs[dirIdx][0], j+dirs[dirIdx][1]
		if nxti < 0 || nxtj < 0 || nxti >= len(matrix) || nxtj >= len(matrix[0]) {
			return obsCount
		}
		if matrix[nxti][nxtj] != '#' {
			// fmt.Println("i, j!!", i, j, dirIdx)
			hasLoop := tryIfCauseLoop(matrix, i, j, nxti, nxtj, dirIdx)
			if hasLoop {
				obsCount++
			}
			i, j = nxti, nxtj
		} else {
			dirIdx = (dirIdx + 1) % 4 // turn right
		}
	}
	return obsCount
}

func tryIfCauseLoop(inputMatrix [][]rune, i, j, obsi, obsj, dirIdx int) bool {
	matrix := make([][]rune, len(inputMatrix))
	for i := range matrix {
		matrix[i] = make([]rune, len(inputMatrix[0]))
		copy(matrix[i], inputMatrix[i])
	}
	matrix[obsi][obsj] = '#'
	visited := make([][]map[int]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]map[int]bool, len(matrix[0]))
		for j := range visited[i] {
			visited[i][j] = make(map[int]bool)
		}
	}
	visited[i][j][dirIdx] = true

	for true {
		nxti, nxtj := i+dirs[dirIdx][0], j+dirs[dirIdx][1]
		if nxti < 0 || nxtj < 0 || nxti >= len(matrix) || nxtj >= len(matrix[0]) {
			return false
		}
		if matrix[nxti][nxtj] != '#' {
			i, j = nxti, nxtj
			if visited[i][j][dirIdx] {
				return true
			}
			visited[i][j][dirIdx] = true
		} else {
			dirIdx = (dirIdx + 1) % 4 // turn right
			visited[i][j][dirIdx] = true
		}
	}
	return false

}

func parseInput(input string) ([][]rune, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix, nil
}
