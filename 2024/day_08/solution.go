package main

import (
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2024/day/8/input

type loc struct {
	x, y int
}

const inputFile = `input.txt`

func main() {
	// Part 1
	totalAntinode, err := part1(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", totalAntinode) // 369

}

func part1(input string) (int, error) {
	m, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	fmt.Println(len(m), len(m[0]))
	groups := groupByFrequency(m) // group by frequency
	// each pair of antennas will create 2 antinode (within boundary)
	antinodes := make([][]bool, len(m))
	for i := range antinodes {
		antinodes[i] = make([]bool, len(m[0]))
	}
	for _, locs := range groups {
		if len(locs) < 2 {
			continue
		}
		for i := 0; i < len(locs); i++ {
			for j := i + 1; j < len(locs); j++ {
				a1, a2 := locs[i], locs[j]
				ns := locateAntinodes(a1, a2)
				n1, n2 := ns[0], ns[1]
				if n1.x >= 0 && n1.x < len(m) && n1.y >= 0 && n1.y < len(m[0]) {
					antinodes[n1.x][n1.y] = true
				}
				if n2.x >= 0 && n2.x < len(m) && n2.y >= 0 && n2.y < len(m[0]) {
					antinodes[n2.x][n2.y] = true
				}
			}
		}
	}

	count := 0
	for _, row := range antinodes {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}

	for _, row := range antinodes {
		fmt.Println(row)
	}
	return count, nil
}

func groupByFrequency(matrix [][]rune) map[rune][]loc {
	m := make(map[rune][]loc)
	for i, row := range matrix {
		for j, char := range row {
			if char != '.' && char != '#' {
				m[char] = append(m[char], loc{i, j})
			}
		}
	}
	return m
}

func locateAntinodes(a1, a2 loc) [2]loc {
	// (n1 + (a2))/2 = a1 => n1 = 2*a1 - a2
	// (n2 + (a1))/2 = a2 => n2 = 2*a2 - a1
	n1x := (a1.x+1)*2 - (a2.x + 1) - 1
	n1y := (a1.y+1)*2 - (a2.y + 1) - 1
	n2x := (a2.x+1)*2 - (a1.x + 1) - 1
	n2y := (a2.y+1)*2 - (a1.y + 1) - 1
	return [2]loc{{n1x, n1y}, {n2x, n2y}}
}

// parse input file and return as a matrix of runes
func parseInput(inputFile string) ([][]rune, error) {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}
	var matrix [][]rune
	for _, line := range strings.Split(string(content), "\n") {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, []rune(line))
	}
	return matrix, nil
}
