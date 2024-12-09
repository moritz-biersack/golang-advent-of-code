package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var part int
	var inputpath string
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.StringVar(&inputpath, "input", "example.txt", "file path of input")
	flag.Parse()
	fmt.Println("Running part", part)

	content, err := os.ReadFile(inputpath)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	if part == 1 {
		sum := part1(string(content))
		fmt.Println("Sum:", sum)
	} else {
		sum := part2(string(content))
		fmt.Println("Sum:", sum)
	}
}

func getSortFunc(rules map[int][]int) func(a, b int) int {
	return func(a, b int) int {
		rules, included := rules[b]
		if included && slices.Contains(rules, a) {
			return 1
		}
		return -1
	}
}

func part1(input string) int {
	instructions := strings.Split(input, "\n\n")
	orderings := strings.Split(instructions[0], "\n")
	updates := strings.Split(instructions[1], "\n")
	rulesPerPage := getRulesPerPage(orderings)
	sortFunc := getSortFunc(rulesPerPage)
	sum := 0
	for _, update := range updates {
		pageNumbers := getPageNumbers(update)
		if slices.IsSortedFunc(pageNumbers, sortFunc) {
			middle := len(pageNumbers) / 2
			sum += pageNumbers[middle]
		}

	}
	
	return sum
}

func getRulesPerPage(orderings []string) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, ord := range orderings {
		parts := strings.Split(ord, "|")
		first, err1 := strconv.Atoi(parts[0])
		second, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			panic("Could not convert rule numbers")
		}
		ruleMap[first] = append(ruleMap[first], second)
	}
	return ruleMap
}

func getPageNumbers(update string) []int {
	pageNumberStings := strings.Split(update, ",")
	var pageNumbers []int

	for _, page := range pageNumberStings {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			panic("Error converting update")
		}
		pageNumbers = append(pageNumbers, pageInt)
	}

	return pageNumbers
}

func part2(input string) int {
	instructions := strings.Split(input, "\n\n")
	orderings := strings.Split(instructions[0], "\n")
	updates := strings.Split(instructions[1], "\n")
	rulesPerPage := getRulesPerPage(orderings)
	sortFunc := getSortFunc(rulesPerPage)
	sum := 0
	for _, update := range updates {
		pageNumbers := getPageNumbers(update)
		if slices.IsSortedFunc(pageNumbers, sortFunc) {
			continue
		}
		slices.SortFunc(pageNumbers, sortFunc)
		sum += pageNumbers[len(pageNumbers)/2]
	}
	
	return sum
}