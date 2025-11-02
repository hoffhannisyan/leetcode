package main

import (
	"fmt"

	"github.com/hoffhannisyan/leetcode/tasks"
)

func main() {
	fmt.Println("Running LeetCode tasks...")

	runTwoSum()
	runAddTwoNumbers()

	fmt.Println("All tasks executed successfully.")
}

// --- Task 1: Two Sum ---
func runTwoSum() {
	fmt.Println("=== Two Sum ===")

	nums := []int{2, 7, 11, 15}
	target := 9
	result := tasks.TwoSum(nums, target)

	fmt.Println("Input:", nums)
	fmt.Println("Target:", target)
	fmt.Println("Output:", result)
	fmt.Println()
}

// --- Task 2: Add Two Numbers ---
func runAddTwoNumbers() {
	fmt.Println("=== Add Two Numbers ===")

	l1 := &tasks.ListNode{
		Val: 2,
		Next: &tasks.ListNode{
			Val: 4,
			Next: &tasks.ListNode{
				Val: 3,
			},
		},
	}

	l2 := &tasks.ListNode{
		Val: 5,
		Next: &tasks.ListNode{
			Val: 6,
			Next: &tasks.ListNode{
				Val: 4,
			},
		},
	}

	result := tasks.AddTwoNumbers(l1, l2)

	fmt.Print("Input 1: [2 4 3]\nInput 2: [5 6 4]\nOutput:  [")
	for node := result; node != nil; node = node.Next {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" ")
		}
	}
	fmt.Print("]\n")
	fmt.Println()
}
