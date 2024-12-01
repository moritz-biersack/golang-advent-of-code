package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)


func main() {
	fmt.Println("hello, advent of code!")
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

func splitLines(input string) (l []int, r []int, err error) {
	var left, right []int
	for _, line := range strings.Split(input, "\n") {
		var a, b int
		_, err := fmt.Sscanf(line, "%d %d", &a, &b)
		if err != nil {
			fmt.Println("Error when splitting line", err)
			return nil, nil, err
		}
		left = append(left, a)
		right = append(right, b)
	}
	return left, right, nil
}

func part1(input string) int {
	var sum int
	left, right, err := splitLines(input)
	if err != nil {
		return -1
	}
	slices.Sort(left)
	slices.Sort(right)
	for i, l := range left {
		r := right[i]
		diff := l - r
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum

}

func part2(input string) int {
	var sum int
	left, right, err := splitLines(input)
	if err != nil {
		return -1
	}
	for _, l := range left {
		var count int
		for _, r := range right {
			if r == l {
				count++
			}
		}
		score := l * count
		sum += score
	}
	return sum
}