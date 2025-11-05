package main

import (
	"fmt"

	"github.com/hoffhannisyan/leetcode/tasks"
)

func main() {
	fmt.Println("LeetCode Solutions in Go")
	fmt.Println("========================")

	// --- LeetCode #1: Two Sum ---
	// nums := []int{2, 7, 11, 15}
	// target := 9
	// result := tasks.TwoSum(nums, target)
	// fmt.Println("Two Sum Output:", result)
	// fmt.Println()

	// --- LeetCode #2: Add Two Numbers ---
	// l1 := &tasks.ListNode{Val: 2, Next: &tasks.ListNode{Val: 4, Next: &tasks.ListNode{Val: 3}}}
	// l2 := &tasks.ListNode{Val: 5, Next: &tasks.ListNode{Val: 6, Next: &tasks.ListNode{Val: 4}}}
	// resultList := tasks.AddTwoNumbers(l1, l2)
	// fmt.Print("Add Two Numbers Output: [")
	// for node := resultList; node != nil; node = node.Next {
	// 	fmt.Print(node.Val)
	// 	if node.Next != nil {
	// 		fmt.Print(" ")
	// 	}
	// }
	// fmt.Print("]\n\n")

	// --- LeetCode #3: Longest Substring Without Repeating Characters ---
	// s := "pwwkew"
	// length := tasks.LengthOfLongestSubstring(s)
	// fmt.Println("Longest Substring Length:", length)
	// fmt.Println()

	// --- LeetCode #4: Median of Two Sorted Arrays ---
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	median := tasks.FindMedianSortedArrays(nums1, nums2)
	fmt.Println("Median of Two Sorted Arrays:", median)
	fmt.Println()

}
