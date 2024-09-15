package main

import (
	"LeetCodeGo/problems"
	"log"
)

func main() {
	input := make([]int, 0)
	input = append(input, 10, 20, 30, 40)
	target := 30
	output := problems.TwoSum(input, target)
	log.Println(output)
}
