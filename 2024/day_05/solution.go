package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Day 5: Print Queue
// https://adventofcode.com/2024/day/5

const inputFile = "input.txt"

func main() {
	total, err := part1(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", total) // 5064

	total, err = part2(inputFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2:", total) // 5152
}

func part1(input string) (int, error) {
	rules, orders, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}
	middlePageSum := 0
	for _, order := range orders {
		if isOrderValid(rules, order) {
			middlePageSum += order[len(order)/2]
		}
	}
	return middlePageSum, nil
}

func isOrderValid(rules map[int]map[int]bool, order []int) bool {
	visited := make(map[int]bool) // to keep track of the visited numbers
	isValid := true
	for _, num := range order {
		if _, ok := rules[num]; ok {
			for shouldNotExistBefore := range rules[num] {
				if visited[shouldNotExistBefore] {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		visited[num] = true
	}
	return isValid
}

func part2(input string) (int, error) {
	rules, orders, err := parseInput(input)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse input: %v", err)
	}
	middlePageSum := 0
	for _, order := range orders {
		if !isOrderValid(rules, order) {
			ordered := orderByRules(rules, order)
			middlePageSum += ordered[len(ordered)/2]
		}
	}
	return middlePageSum, nil
}

func orderByRules(rules map[int]map[int]bool, order []int) []int {
	res := make([]int, len(order))
	copy(res, order)
	continueLoop := true
	for continueLoop {
		for i := 1; i < len(res); i++ {
			if !rules[res[i-1]][res[i]] {
				continueLoop = true
				reflect.Swapper(res)(i-1, i)
				break
			} else {
				continueLoop = false
			}
		}
	}
	return res
}

// parseInput reads the input file and returns the rules and orders
// rules is a map where the list of numbers (value) should not happens BEFORE the key
func parseInput(fileName string) (map[int]map[int]bool, [][]int, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	var rules map[int]map[int]bool = make(map[int]map[int]bool)
	var orders [][]int
	part2 := false
	for _, line := range lines {
		if len(line) == 0 {
			part2 = true
			continue
		}
		if !part2 {
			l := strings.Split(line, "|")
			left, err := strconv.Atoi(l[0])
			if err != nil {
				return nil, nil, fmt.Errorf("couldn't parse rule: %v", err)
			}
			right, err := strconv.Atoi(l[1])
			if err != nil {
				return nil, nil, fmt.Errorf("couldn't parse rule: %v", err)
			}
			if _, ok := rules[left]; ok {
				rules[left][right] = true
			} else {
				rules[left] = map[int]bool{right: true}
			}
		} else {
			var o []int
			for _, num := range strings.Split(line, ",") {
				n, err := strconv.Atoi(num)
				if err != nil {
					return nil, nil, fmt.Errorf("couldn't parse number: %v", err)
				}
				o = append(o, n)
			}
			orders = append(orders, o)
		}
	}
	return rules, orders, nil
}
