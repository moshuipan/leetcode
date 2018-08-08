package leetcode

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。
// 请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。
// 你可以假设 nums1 和 nums2 均不为空。
// 示例 1:
// nums1 = [1, 3]
// nums2 = [2]
// 中位数是 2.0
// 示例 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// 中位数是 (2 + 3)/2 = 2.5
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	x, y := nums1, nums2
	xl, yl := len(x), len(y)
	var midIndexStart, midIndexEnd int
	all := xl + yl
	if all%2 == 0 {
		midIndexStart, midIndexEnd = all/2, (all/2)+1
	} else {
		midIndexStart, midIndexEnd = (all+1)/2, (all+1)/2
	}
	midIndexStart, midIndexEnd = midIndexStart-1, midIndexEnd-1
	var sum, xi, yi int
	for i := 0; i <= midIndexEnd; i++ {
		z := 0
		if xi < xl && yi < yl && x[xi] < y[yi] {
			z = x[xi]
			xi++
		} else if xi < xl && yi < yl && x[xi] >= y[yi] {
			z = y[yi]
			yi++
		} else if xi >= xl {
			z = y[yi]
			yi++
		} else if yi >= yl {
			z = x[xi]
			xi++
		}
		if i == midIndexStart {
			sum += z
			if midIndexEnd == midIndexStart {
				break
			}
		} else if i == midIndexEnd {
			sum += z
			break
		}
	}
	if midIndexEnd == midIndexStart {
		return float64(sum)
	} else {
		return float64(sum) / 2.0
	}
}
