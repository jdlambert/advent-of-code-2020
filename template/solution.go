package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}