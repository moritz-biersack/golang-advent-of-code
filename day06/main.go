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

type Direction int
const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
	EMPTY
)
var DirectionMap = map[rune]Direction{
	'^': UP,
	'>': RIGHT,
	'v': DOWN,
	'<': LEFT,
}

type Guard struct {
	x int
	y int
	dir Direction
}

func (g *Guard) turnUp() {
	g.dir = UP
}

func (g *Guard) turnRight() {
	g.dir = RIGHT
}

func (g *Guard) turnDown() {
	g.dir = DOWN
}

func (g *Guard) turnLeft() {
	g.dir = LEFT
}

func (g *Guard) move() {
	switch g.dir {
	case UP:
		g.y--
	case RIGHT:
		g.x++
	case DOWN:
		g.y++
	case LEFT:
		g.x--
	}
}

func part1(input string) int {
	// nop
	height, width, obstacles, seen, guard := mapSetup(input)
	runGuard(guard, height, width, seen, obstacles)
	sum := 0
	for y, row := range seen {
		for x, s := range row {
			if s != EMPTY {
				sum++
			} else if obstacles[y][x] {
			} else {
			}
		}
	}
	return sum
}

func mapSetup(input string) (int, int, [][]bool, [][]Direction, Guard) {
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	obstacles := make([][]bool, height)
	seen := make([][]Direction, height)
	var guard Guard

	for y, line := range lines {
		obstacles[y] = make([]bool, width)
		seen[y] = make([]Direction, width)
		for x, char := range line {
			seen[y][x] = EMPTY
			switch char {
			case '#':
				obstacles[y][x] = true
			case '^':
				guard = Guard{y: y, x: x, dir: UP}
			case '.':

			default:
				panic("Unrecognized character: " + string(char))

			}
		}
	}
	return height, width, obstacles, seen, guard
}

type GuardEndState int
const (
	OFFMAP = iota
	LOOPED
)

func runGuard(guard Guard, height int, width int, seen [][]Direction, obstacles [][]bool) GuardEndState{
	started := false
	for {
		if guard.x >= height || guard.x < 0 || guard.y >= width || guard.y < 0 {
			return OFFMAP
		}
		x, y := guard.x, guard.y
		switch guard.dir {
		case UP:
			if obstacles[max(y-1, 0)][x] {
				guard.turnRight()
			} else {
				guard.turnUp()
			}
		case RIGHT:
			if obstacles[y][min(x+1, width-1)] {
				guard.turnDown()
			} else {
				guard.turnRight()
			}
		case DOWN:
			if obstacles[min(y+1, height-1)][x] {
				guard.turnLeft()
			} else {
				guard.turnDown()
			}
		case LEFT:
			if obstacles[y][max(x-1, 0)] {
				guard.turnUp()
			} else {
				guard.turnLeft()
			}
		}
		if started && (guard.dir == seen[y][x]) {
			return LOOPED
		}
		seen[y][x] = guard.dir
		guard.move()
		started = true
	}
}

func part2(input string) int {
	height, width, obstacles, initiallySeen, guard := mapSetup(input)
	startY, startX := guard.y, guard.x
	runGuard(guard, height, width, initiallySeen, obstacles)
	count := 0
	var seen [][]Direction
	for y := range height {
		for x := range width {
			height, width, obstacles, seen, guard = mapSetup(input)
			if initiallySeen[y][x] == EMPTY || obstacles[y][x] || (x == startX && y == startY) {
				continue
			}
			if obstacles[y][x] || (x == startX && y == startY) {
				continue
			}
			obstacles[y][x] = true
			result := runGuard(guard, height, width, seen, obstacles)
			if result == LOOPED {
				count++
				fmt.Println("Looped:", count)
			}
		}
	}
	return count
}