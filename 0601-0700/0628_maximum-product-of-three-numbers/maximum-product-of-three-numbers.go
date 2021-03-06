package main

func maximumProduct(nums []int) int {
	max0, max1, max2 := -1<<31, -1<<31, -1<<31
	min0, min1 := 1<<31-1, 1<<31-1
	for _, num := range nums {
		if num > max2 {
			max0, max1, max2 = max1, max2, num
		} else if num > max1 {
			max0, max1 = max1, num
		} else if num > max0 {
			max0 = num
		}

		if num < min0 {
			min0, min1 = num, min0
		} else if num < min1 {
			min1 = num
		}
	}

	if max2 < 0 {
		return max0 * max1 * max2
	} else {
		return Max(min0*min1, max0*max1) * max2
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
