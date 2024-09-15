package problems

import (
	"log"
	"sort"
)

func TwoSum(nums []int, target int) []int {
	return nil
}

// here's brute force approach O(n^2): 59 ms
func _(nums []int, target int) []int {
	for current := range nums {
		next := current + 1
		for next = range nums {
			if current == next {
				continue
			}
			if nums[current]+nums[next] == target {
				return []int{current, next}
			}
		}
	}
	return nil
}

// two pointer approach: O( n * log(n) ) : 3 ms,3.96mb
func _(nums []int, target int) []int {
	// copy array with indexes:
	// ex: nums= 20,10 :
	// indNums: [20,1],[10,2],[20,10]
	indNums := make([][2]int, len(nums))
	for i, num := range nums {
		indNums[i] = [2]int{num, i}
	}

	// sort the array
	sort.Slice(indNums, func(i, j int) bool {
		return indNums[i][0] < indNums[j][0]
	})
	//log.Println(indNums)
	// indNums : [10,2],[20,1]
	left, right := 0, len(nums)-1
	for left < right {
		//	log.Printf("l:%d,indNums[l]:%d,r:%d,indNums[r]:%d\n", left, indNums[left], right, indNums[right])
		sum := indNums[left][0] + indNums[right][0]
		if sum == target {
			return []int{indNums[left][1], indNums[right][1]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

// two pass  hashmap approach: 15ms
// A = nums[i], number2 must be target-A, look for complement in map
// this does two passes at the map
func _(nums []int, target int) []int {

	m := make(map[int]int)
	//init hashmap
	for i, num := range nums {
		m[num] = i
	}
	log.Println(m)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}
	}
	return nil

}

// one pass hashmap approach: 0ms,4.07mb
// this one also from @kartikdevsharmaa
// while checking a number, we also check if the complement is already there.
func _(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
