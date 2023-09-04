func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}

	l1, l2 := len(nums1), len(nums2)
	deltaForPartitionB := (l1 + l2 + 1) / 2
	left, right := 0, l1
	for left <= right {
		partitionA := (left + right) / 2
		partitionB := deltaForPartitionB - partitionA
		var maxLeftA, minRightA, maxLeftB, minRightB int
		setIf(&maxLeftA, partitionA == 0, math.MinInt, partitionA-1, nums1)
		setIf(&minRightA, partitionA == l1, math.MaxInt, partitionA, nums1)
		setIf(&maxLeftB, partitionB == 0, math.MinInt, partitionB-1, nums2)
		setIf(&minRightB, partitionB == l2, math.MaxInt, partitionB, nums2)
		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			if (l1+l2)%2 == 0 {
				return (math.Min(float64(minRightA), float64(minRightB)) +
					math.Max(float64(maxLeftA), float64(maxLeftB))) / 2.0
			} else {
				return math.Max(float64(maxLeftB), float64(maxLeftA))
			}
		} else if maxLeftA > minRightB {
			right = partitionA - 1
		} else {
			left = partitionA + 1
		}
	}

	return 0.0
}

func setIf(target *int, flag bool, val, elsIdx int, nums []int) {
	if flag {
		*target = val
	} else {
		*target = nums[elsIdx]
	}
}