package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
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

func checkForward(subString string) int {
	if subString == "XMAS" {
		return 1
	}
	return 0
}

func checkBackward(subString string) int {
	reversed := make([]byte, len(subString))
	copy(reversed, subString)
	slices.Reverse(reversed)
	reversedString := string(reversed)
	
	return checkForward(reversedString)
}

func checkDown(lines []string, x int, y int) int {
	var subString = make([]byte, 4)
	for i := range 4 {
		subString[i] = lines[y+i][x]
	}
	count := 0
	count += checkForward(string(subString))
	count += checkBackward(string(subString))

	return count
}

func checkDownRight(lines []string, x int, y int) int {
	var subString = make([]byte, 4)
	for i := range 4 {
		subString[i] = lines[y+i][x+i]
	}
	count := 0
	count += checkForward(string(subString))
	count += checkBackward(string(subString))

	return count
}

func checkUpRight(lines []string, x int, y int) int {
	var subString = make([]byte, 4)
	for i := range 4 {
		subString[i] = lines[y-i][x+i]
	}
	count := 0
	count += checkForward(string(subString))
	count += checkBackward(string(subString))

	return count
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for y:=0; y<len(lines); y++ {
		line := lines[y]
		xLimit := len(line) - 3
		yLimit := len(lines) - 3
		for x:=0; x<len(line); x++ {
			if x < xLimit{
				subString := line[x:x+4]
				count += checkForward(subString)
				count += checkBackward(subString)
				if y >= 3 {
					count += checkUpRight(lines, x, y)
				}
				if y < yLimit {
					count += checkDownRight(lines, x, y)
				}
			}
			if y < yLimit {
				count += checkDown(lines, x, y)
			}
		}
	}

	return count
}


func part2(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for y:=1; y<len(lines)-1; y++ {
		line := lines[y]
		for x:=1; x<len(line)-1; x++ {
			if line[x] != 'A' {
				continue
			}
			diagonal1 := make([]byte, 3)
			diagonal2 := make([]byte, 3)
			for i := range 3 {
				diagonal1[i] = lines[y-1+i][x-1+i]
				diagonal2[i] = lines[y+1-i][x-1+i]
			}
			dia1Ok := string(diagonal1) == "MAS" || string(diagonal1) == "SAM"
			dia2Ok := string(diagonal2) == "MAS" || string(diagonal2) == "SAM"
			if dia1Ok && dia2Ok {
				count++
			}
		}
	}

	return count
}