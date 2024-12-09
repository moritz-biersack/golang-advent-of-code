package main

import (
	"flag"
	"fmt"
	"os"
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

func getPermutations(n int, current []string, result *[][]string) {
	if len(current) == n {
		c := make([]string, len(current))
		copy(c, current)
		*result = append(*result, c)
		return
	}
	getPermutations(n, append(current, "*"), result)
	getPermutations(n, append(current, "+"), result)

}
func part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		test_value, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("Couldn't convert test value")
		}
		var numbers []int
		for _, n := range strings.Split(parts[1], " ") {
			if n == "" {
				continue
			}
			number, err := strconv.Atoi(n)
			if err != nil {
				panic("Couldn't convert number")
			}
			numbers = append(numbers, number)
		}
		permutations := [][]string{}
		getPermutations(len(numbers)-1, []string{}, &permutations)
		for _, p := range permutations {
			total := 0
			for i, operation := range p {
				switch operation {
				case "*":
					if i == 0 {
						total = numbers[i] * numbers[i+1]
					} else {
						total = total * numbers[i+1]
					}
				case "+":
					if i == 0 {
						total = numbers[i] + numbers[i+1]
					} else {
						total = total + numbers[i+1]
					}
				default:
					panic("Invalid operation!")
				}
			}
			if total == test_value {
				sum += total
				break
			}
		}
	}
	return sum
}


func part2(input string) int {
	return -1
}