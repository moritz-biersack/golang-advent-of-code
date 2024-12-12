package main

import (
	"flag"
	"fmt"
	"os"
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

func isOutOfBounds(y, x, h, w int) bool {
	return y < 0 || y >= h || x < 0 || x >= w
}

func part1(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	antennas := make([][]rune, len(lines))
	antinodes := make([][]rune, len(lines))
	h := len(lines)
	var w int
	for y, line := range lines {
		rs := []rune(line)
		w = len(rs)
		antennas[y] = make([]rune, w)
		antinodes[y] = make([]rune, w)
		copy(antennas[y], rs)
		copy(antinodes[y], rs)
	}
	for y := range h {
		for x := range w {
			currentChar := antennas[y][x]
			if currentChar == '.' {
				continue
			}
			for cy := range h {
				for cx := range w {
					if cy == y && cx == x {
						continue
					}
					if antennas[cy][cx] != currentChar {
						continue
					}
					dy := cy - y
					dx := cx - x
					ccy := cy + dy
					ccx := cx + dx
					if !isOutOfBounds(ccy, ccx, w, h) {
						if antinodes[ccy][ccx] != '#' {
							antinodes[ccy][ccx] = '#' 
							sum++
						}
					}
				}
			}
		}
	}
	for y := range h {
		for x := range w {
			fmt.Print(string(antinodes[y][x]))
		}
		fmt.Println()
	}
	return sum
}


func part2(input string) int {
	return -1
}
