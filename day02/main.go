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

func part1(input string) int {
	lines := strings.Split(input, "\n")
	safe_count := 0
	for _, line := range lines {
		is_safe := processLine(line)
		safe_count += is_safe
	}
	return safe_count

}

func processLine(line string) int {
	levels := strings.Split(line, " ")
	previous, err := strconv.Atoi(levels[0])
	if err != nil {
		panic(err)
	}
	second, err := strconv.Atoi(levels[1])
	if err != nil {
		panic(err)
	}
	is_safe := 1
	increasing := previous < second
	for i := 1; i < len(levels); i++ {
		level_str := levels[i]
		level, err := strconv.Atoi(level_str)
		if err != nil {
			panic(err)
		}
		currently_increasing := previous < level
		if increasing != currently_increasing {
			is_safe = 0
			break
		}
		diff := previous - level
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			is_safe = 0
			break
		}
		previous = level
	}
	return is_safe
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	safe_count := 0
	for _, line := range lines {
		is_safe := processLine(line)
		if is_safe == 1 {
			safe_count++
			continue
		}
		// brute force it by removing levels
		levels := strings.Split(line, " ")
		for i:=0; i<len(levels); i++ {
			c := make([]string, len(levels))
			copy(c, levels)
			removedOne := append(c[:i], c[i+1:]...)
			newLine := strings.Join(removedOne, " ")
			is_safe := processLine(newLine)
			if is_safe == 1 {
				safe_count++
				break
			}
		}

	}
	return safe_count
}