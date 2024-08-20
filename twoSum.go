package LeetCodeGo

func twoSum(nums []int, target int) []int {
	nums = removeBiggerFromArray(nums, target)
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			continue
		}
		num1 := nums[i]
		//fmt.Printf("num1= %d\n",num1)
		for j := 1; j < len(nums); j++ {
			if nums[j] == -1 {
				continue
			}
			num2 := nums[j]
			//fmt.Printf("num2= %d\n",num2)
			sum := num1 + num2
			if sum == target {
				//fmt.Printf("sum: %d",sum)
				list := make([]int, 0)
				list = append(list, i)
				list = append(list, j)
				//fmt.Println(list)
				return list
			}
		}
	}
	return nil
}

func removeBiggerFromArray(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		if target < nums[i] {
			nums[i] = -1
		}
	}
	//fmt.Printf("% d\n",nums)
	return nums
}

// failed at nums= [2,5,5,11]
// create a unique function
