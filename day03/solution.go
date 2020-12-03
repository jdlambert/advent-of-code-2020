package main

import (
	"bufio"
	"fmt"
	"os"
)

func treeCount(trees []string, dRow int, dCol int) (count int) {

	for row, col := 0, 0; row < len(trees); row, col = row+dRow, (col+dCol)%len(trees[0]) {
		if trees[row][col] == '#' {
			count++
		}
	}

	return
}

func partOne(trees []string) int {
	return treeCount(trees, 1, 3)
}

func partTwo(trees []string) int64 {
	product := int64(treeCount(trees, 2, 1))
	for i := 1; i < 9; i += 2 {
		product *= int64(treeCount(trees, 1, i))
	}
	return product
}

func main() {
	reader, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(reader)
	trees := make([]string, 0)
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}

	fmt.Println(partOne(trees))
	fmt.Println(partTwo(trees))
}
