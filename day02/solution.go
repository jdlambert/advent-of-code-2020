package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passwordInfo struct {
	first    int
	second   int
	target   string
	password string
}

func partTwo(passwords []passwordInfo) (valid int) {
	for _, password := range passwords {

		if (string(password.password[password.first]) == password.target) !=
			(string(password.password[password.second]) == password.target) {
			valid++
		}
	}
	return
}

func partOne(passwords []passwordInfo) (valid int) {
	for _, password := range passwords {
		count := strings.Count(password.password, password.target)

		if password.first <= count && count <= password.second {
			valid++
		}
	}
	return
}

func main() {
	reader, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(reader)
	passwords := make([]passwordInfo, 0)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")
		policy, password := splitLine[0], splitLine[1]

		policySplit := strings.Split(policy, " ")
		target := policySplit[1]
		letterRange := strings.Split(policySplit[0], "-")
		first, _ := strconv.Atoi(letterRange[0])
		second, _ := strconv.Atoi(letterRange[1])

		if password != "" {
			passwords = append(passwords, passwordInfo{first: first, second: second, target: target, password: password})
		}
	}

	fmt.Println(partOne(passwords))
	fmt.Println(partTwo(passwords))
}
