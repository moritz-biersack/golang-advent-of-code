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

var directions = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func walkMap(topo [][]int, y, x int, visited [][]bool, trailEnds map[[2]int]int) {
	visited[y][x] = true
	if topo[y][x] == 9 {
		trailEnds[[2]int{y, x}]++
		return
	}
	for _, d := range directions {
		dy, dx := d[0], d[1]
		cy, cx := y+dy, x+dx
		if cy < 0 || cx < 0 || cy >= len(topo) || cx >= len(topo[0]) || visited[cy][cx] {
			continue
		}
		currentHeight := topo[y][x]
		if topo[cy][cx]-currentHeight != 1 {
			continue
		}
		visitedNew := make([][]bool, len(visited))
		for y, row := range visited {
			visitedNew[y] = make([]bool, len(row))
			copy(visitedNew[y], visited[y])
		}
		walkMap(topo, cy, cx, visitedNew, trailEnds)
	}
}

func part1(input string) int {
	topo := prepareTopo(input)
	sum := 0
	for y := range len(topo) {
		for x, n := range topo[y] {
			if n == 0 {
				visited := makeVisited(topo)
				trailEnds := make(map[[2]int]int)
				walkMap(topo, y, x, visited, trailEnds)
				sum += len(trailEnds)
			}
		}
	}

	return sum
}

func part2(input string) int {
	topo := prepareTopo(input)
	sum := 0
	for y := range len(topo) {
		for x, n := range topo[y] {
			if n == 0 {
				visited := makeVisited(topo)
				trailEnds := make(map[[2]int]int)
				walkMap(topo, y, x, visited, trailEnds)
				for _, v := range trailEnds {
					sum += v
				}
			}
		}
	}

	return sum
}

func makeVisited(topo [][]int) [][]bool {
	visited := make([][]bool, len(topo))
	for y := range visited {
		visited[y] = make([]bool, len(topo[y]))
	}
	return visited
}

func prepareTopo(input string) [][]int {
	lines := strings.Split(input, "\n")
	topo := make([][]int, len(lines))
	for y, line := range lines {
		for _, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				n = -1
			}
			topo[y] = append(topo[y], n)
		}
	}
	return topo
}
