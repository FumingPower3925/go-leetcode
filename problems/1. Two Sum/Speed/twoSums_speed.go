func twoSum(nums []int, target int) []int {
    // Store the nums value in the key and the position in the value
    hist := make(map[int]int)
        for i, num := range nums {
            // We search for the "other" number expected value
            if j, ok := hist[target-num]; ok {
                return []int{j, i}
            } else {
                // We add de new element to the histogram
                hist[num] = i
            }
        }
    return []int{}
}