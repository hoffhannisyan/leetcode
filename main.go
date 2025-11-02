package main

import (
	"fmt"

	"github.com/hoffhannisyan/leetcode/tasks"
)

func main() {
	fmt.Println("Running LeetCode tasks...")
	fmt.Println()

	runTwoSum()

	fmt.Println()
	fmt.Println("✅ All tasks executed successfully.")
}

func runTwoSum() {
	fmt.Println("=== Two Sum ===")

	nums := []int{2, 7, 11, 15}
	target := 9

	result := tasks.TwoSum(nums, target)

	fmt.Println("Input: ", nums)
	fmt.Println("Target:", target)
	fmt.Println("Output:", result)
}
