package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const word = "XMAS"
var directions = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
}

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
	count := 0
	for y:=0; y<len(lines); y++ {
		line := lines[y]
		for x:=0; x<len(line); x++ {
			if line[x] != 'X' {
				continue
			}
			for _, dir := range directions {
				dy, dx := dir[0], dir[1]
				charsMatching := true
				for i := range len(word) {
					cy := y + dy*i
					cx := x + dx*i
					if cy < 0 || cx < 0 || cy >= len(lines) || cx >= len(line) {
						charsMatching = false
						break
					}
					c := lines[cy][cx]
					if c != word[i] {
						charsMatching = false
						break
					}
				}
				if charsMatching {
					count++
				}
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