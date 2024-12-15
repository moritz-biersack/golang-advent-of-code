package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
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
	disk, _ := createDisk(input)
	for i := range len(disk) {
		if disk[i] != "." {
			continue
		}
		j := len(disk) - 1
		for range len(disk) {
			if disk[j] != "." {
				break
			}
			j--
		}
		if j <= i {
			break
		}
		disk[i], disk[j] = disk[j], disk[i]
	}
	sum := 0
	for i, num := range disk {
		n, _ := strconv.Atoi(num)
		sum += i * n
	}
	return sum
}

func createDisk(input string) ([]string, []int) {
	blocks := []int{}
	freeSpace := []int{}
	size := 0
	for i, r := range input {
		num, err := strconv.Atoi(string(r))
		if err != nil {
			panic("No number!")
		}
		size += num
		if i%2 == 0 {
			blocks = append(blocks, num)
		} else {
			freeSpace = append(freeSpace, num)
		}
	}
	disk := make([]string, size)
	offset := 0
	for i := 0; i < len(blocks); i++ {
		for b := range blocks[i] {
			disk[offset+b] = fmt.Sprint(i)
		}
		offset += blocks[i]
		if i >= len(freeSpace) {
			continue
		}
		for f := range freeSpace[i] {
			disk[offset+f] = "."
		}
		offset += freeSpace[i]
	}
	return disk, blocks
}

func part2(input string) int {
	disk, blocks := createDisk(input)
	for i := len(blocks) - 1; i > 0; i-- {
		fileId := fmt.Sprint(i)
		startIndex := slices.Index(disk, fileId)
		fileSize := 1
		for startIndex+fileSize < len(disk) {
			if disk[startIndex+fileSize] != fileId {
				break
			}
			fileSize++
		}
		head := 0
		for head < startIndex {
			if disk[head] != "." {
				head++
				continue
			}
			freeSpace := 1
			for head+freeSpace < len(disk) && disk[head+freeSpace] == "." {
				freeSpace++
			}
			if fileSize <= freeSpace {
				for j := range fileSize {
					disk[head+j], disk[startIndex+j] = disk[startIndex+j], disk[head+j]
				}
				break
			}
			head += freeSpace
		}
	}
	sum := 0
	for i, num := range disk {
		n, _ := strconv.Atoi(num)
		sum += i * n
	}
	return sum
}
