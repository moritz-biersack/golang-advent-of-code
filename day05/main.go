package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
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
	instructions := strings.Split(input, "\n\n")
	orderings := strings.Split(instructions[0], "\n")
	updates := strings.Split(instructions[1], "\n")
	rulesPerPage := getRulesPerPage(orderings)
	sum := 0
	for _, update := range updates {
		allGood := true
		pageNumbers := getPageNumbers(update)
		for i, page := range pageNumbers {
			good := true
			pageRules, included := rulesPerPage[page]
			if !included {
				continue
			}
			for _, r := range pageRules {
				rIndex := slices.Index(pageNumbers, r)
				if rIndex < i && rIndex != -1 {
					good = false
					break
				}
			}
			if !good {
				allGood = false
				break
			}
		}
		if allGood {
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

func getPageNumbersLinked(update string) LinkedList {
	pageNumberStrings := strings.Split(update, ",")
	var list LinkedList
	

	for _, page := range pageNumberStrings {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			panic("Error converting update")
		}
		list.Add(pageInt)
	}

	return list
}

type Node struct {
    value int
    next  *Node
}

type LinkedList struct {
    head *Node
	len int
}

func (l *LinkedList) Add(n int) {
	l.len++
	newNode := &Node{value: n}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Get(i int) *Node {
	current := l.head
	for range i {
		if current.next == nil {
			panic("Out of index")
		}
		current = current.next
	}
	return current
}

func (l *LinkedList) Print() {
	current := l.head
	fmt.Print(current.value, "|")
	for current.next != nil {
		current = current.next
		fmt.Print(current.value, "|")
	}
	fmt.Println()
}

func part2(input string) int {
	instructions := strings.Split(input, "\n\n")
	orderings := strings.Split(instructions[0], "\n")
	updates := strings.Split(instructions[1], "\n")
	rulesPerPage := getRulesPerPage(orderings)
	sum := 0
	for _, update := range updates {
		allGood := true
		pageNumbers := getPageNumbersLinked(update)
		for i:=1; i<pageNumbers.len; i++ {
			page := pageNumbers.Get(i)
			pageRules, included := rulesPerPage[page.value]
			if !included {
				continue
			}
			current := pageNumbers.head
			var prev *Node
			for j:=0; j<i; j++ {
				if slices.Contains(pageRules, current.value) {
					allGood = false
					if prev == nil {
						pageNumbers.head = current.next
					} else {
						prev.next = current.next
						prev = prev.next
					}
					temp := page.next
					page.next = current
					current.next = temp
					if prev == nil {
					current = pageNumbers.head
					} else {
						current = prev.next
					}
					i = j
				} else {
					prev = current
					current = current.next
				}
			}
		}
		if !allGood {
			middle := pageNumbers.len / 2
			sum += pageNumbers.Get(middle).value
		}

	}
	
	return sum
}