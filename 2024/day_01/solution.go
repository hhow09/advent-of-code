package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1

func main() {
	totalDist := part1()
	fmt.Println("Part 1:", totalDist)

	similarityScore := part2()
	fmt.Println("Part 2:", similarityScore)
}

func part1() int {
	left, right, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance += Abs(right[i] - left[i])
	}
	return totalDistance
}

func part2() int {
	left, right, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	rightCounter := make(map[int]int)
	for _, r := range right {
		rightCounter[r]++
	}
	simlarityScore := 0
	for _, num := range left {
		simlarityScore += num * rightCounter[num]
	}
	return simlarityScore
}

func parseInput(fileName string) ([]int, []int, error) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}
	lines := strings.Split(string(f), "\n")
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))
	for _, line := range lines {
		// parse line
		pair := strings.Split(line, "   ")
		if len(pair) != 2 {
			continue
		}
		l, err := strconv.Atoi(pair[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid input: %s", line)
		}
		r, err := strconv.Atoi(pair[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid input: %s", line)
		}
		left = append(left, l)
		right = append(right, r)
	}
	return left, right, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
