package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day 2: Red-Nosed Reports
// https://adventofcode.com/2024/day/2

const inputFile = "input.txt"

func main() {
	safeReports := part1()
	fmt.Println("Part 1:", safeReports)

	safeReportsTolerate := part2()
	fmt.Println("Part 2:", safeReportsTolerate)
}

func part1() int {
	reports, err := parseInput(inputFile)
	if err != nil {
		panic(err)
	}
	safeReports := 0
	for _, report := range reports {
		if checkReportSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

func checkReportSafe(report []int) bool {
	// either ascending or descending
	asc := true
	if report[0] > report[1] {
		asc = false
	}
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		dist := diff
		if !asc {
			dist = -diff
		}
		if dist > 3 || dist < 1 {
			return false
		}
	}
	return true
}

func part2() int {
	reports, err := parseInput(inputFile)
	if err != nil {
		panic(err)
	}
	safeReports := 0
	for _, report := range reports {
		if checkReportSafeWithTolerate(report) {
			safeReports++
		}
	}
	return safeReports
}

func checkReportSafeWithTolerate(report []int) bool {
	safeStrict := checkReportSafe(report)
	if safeStrict {
		return true
	}
	// brute force: try to remove one element iteratively
	for i := 0; i < len(report); i++ {
		copied := make([]int, len(report))
		copy(copied, report)
		removed := append(copied[:i], copied[i+1:]...)
		safe := checkReportSafe(removed)
		if safe {
			return true
		}
	}
	return false
}

func parseInput(fileName string) ([][]int, error) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(f), "\n")
	res := make([][]int, 0, len(lines))
	for _, line := range lines {
		// parse line
		nums := strings.Split(line, " ")
		if len(nums) == 1 {
			continue
		}
		lineNums := make([]int, 0, len(nums))
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("invalid input: %s", line)
			}
			lineNums = append(lineNums, n)
		}
		res = append(res, lineNums)
	}
	return res, nil
}
