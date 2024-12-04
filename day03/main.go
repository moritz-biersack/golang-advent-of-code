package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func part1(input string) int {
	sum := sumMultiplies(input)
	return sum

}

func sumMultiplies(input string) int {
	r, _ := regexp.Compile(`(?:mul\()(\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		a, errA := strconv.Atoi(match[1])
		b, errB := strconv.Atoi(match[2])
		if errA != nil || errB != nil {
			panic("Couldn't convert numbers")
		}
		sum += a * b
	}
	return sum
}

func part2(input string) int {
	r := regexp.MustCompile(`(?s)(?:do\(\)|^)(.*?)(?:don't\(\)|$)`)
	matches := r.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		relevantPart := match[1]
		sum += sumMultiplies(relevantPart)
	}

	return sum
}