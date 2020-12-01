package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(reader)
	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	seen := make(map[int]bool)

	for _, num := range nums {
		complement := 2020 - num
		_, ok := seen[complement]
		if ok {
			fmt.Println(complement * num)
			break
		}
		seen[num] = true
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			total := nums[i] + nums[j] + nums[k]
			if total == 2020 {
				fmt.Println(nums[i] * nums[j] * nums[k])
				return
			} else if total < 2020 {
				j++
			} else {
				k--
			}
		}
	}
}
