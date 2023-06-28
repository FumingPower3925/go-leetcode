func twoSum(nums []int, target int) []int {
	// For each number we check if there is any other follwing number that will make up to target
    for i, num := range nums {
		// We don't need to ckeck the previous numbers as they've been already checked
        for j := i+1; j < len(nums); j++ {
            if num + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}