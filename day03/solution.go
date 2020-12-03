package main

import (
	"bufio"
	"fmt"
	"os"
)

func treeCount(trees [][]bool, dRow int, dCol int) (count int) {

	for row, col := 0, 0; row < len(trees); row, col = row+dRow, (col+dCol)%len(trees[0]) {
		if trees[row][col] {
			count++
		}
	}

	return
}

func partOne(trees [][]bool) int {
	return treeCount(trees, 1, 3)
}

func partTwo(trees [][]bool) int64 {
	product := int64(treeCount(trees, 2, 1))
	for i := 1; i < 9; i += 2 {
		product *= int64(treeCount(trees, 1, i))
	}
	return product
}

func main() {
	reader, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(reader)
	trees := make([][]bool, 0)
	for scanner.Scan() {
		row := make([]bool, 0)
		for _, item := range scanner.Text() {
			row = append(row, item == '#')
		}
		trees = append(trees, row)
	}

	fmt.Println(partOne(trees))
	fmt.Println(partTwo(trees))
}
