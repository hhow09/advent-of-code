package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Day 7: Bridge Repair
// https://adventofcode.com/2024/day/7

const inputFile = "input.txt"

const (
	add    = 1
	mul    = 2
	concat = 3
)

type equation struct {
	result int
	nums   []int
}

func main() {
	total, err := part1(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", total) // 3312271365652

	total2, err := part2(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2:", total2) // 509463489296712
}

func part1(input string) (int, error) {
	eqs, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}
	res := 0
	for _, eq := range eqs {
		if verifyEquation(eq, []int{add, mul}) {
			res += eq.result
		}
	}
	return res, nil
}

func part2(input string) (int, error) {
	eqs, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}
	res := 0
	for _, eq := range eqs {
		if verifyEquation(eq, []int{add, mul, concat}) {
			res += eq.result
		}
	}
	return res, nil
}

func verifyEquation(eq equation, operators []int) bool {
	oppermus := allCombinations(operators, len(eq.nums)-1)
	for _, ops := range oppermus {
		if evalEquationOpPossible(eq, ops) {
			return true
		}
	}
	return false
}

// allCombinations returns all possible combinations of the given operators with the given length i.
// ref: https://docs.python.org/3/library/itertools.html#itertools.combinations
func allCombinations(its []int, i int) [][]int {
	if i == 0 {
		return [][]int{{}}
	}
	perms := allCombinations(its, i-1)
	var res [][]int
	for _, perm := range perms {
		for _, it := range its {
			copied := make([]int, len(perm))
			copy(copied, perm)
			res = append(res, append(copied, it))
		}
	}
	return res
}

func evalEquationOpPossible(eq equation, ops []int) bool {
	// Operators are always evaluated left-to-right, not according to precedence rules.
	// Furthermore, numbers in the equations cannot be rearranged.
	// Glancing into the jungle, you can see elephants holding two different types of operators: add (+) and multiply (*).
	curr := eq.nums[0]
	for i, num := range eq.nums[1:] {
		switch ops[i] {
		case add:
			curr += num
		case mul:
			curr *= num
		case concat:
			curr, _ = strconv.Atoi(fmt.Sprintf("%d%d", curr, num))
		}
	}

	return curr == eq.result
}

func parseInput(input string) ([]equation, error) {
	content, err := os.ReadFile(input)
	if err != nil {
		return nil, fmt.Errorf("couldn't read file: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	eqs := make([]equation, 0, len(lines))
	for _, line := range lines {
		a := strings.Split(line, ":")
		res, err := strconv.Atoi(a[0])
		if err != nil {
			return nil, fmt.Errorf("couldn't parse %q: %v", a[0], err)
		}
		numsStr := strings.Split(strings.TrimSpace(a[1]), " ")
		nums, err := parseNumSlice(numsStr)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse %q: %v", numsStr, err)
		}
		eqs = append(eqs, equation{result: res, nums: nums})
	}
	return eqs, nil
}

func parseNumSlice(nums []string) ([]int, error) {
	var res []int
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse %q: %v", num, err)
		}
		res = append(res, n)
	}
	return res, nil
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
