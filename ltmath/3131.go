package ltmath

import "slices"

func AddedInteger(nums1, nums2 []int) int {

	nums1Max := slices.Max(nums1)
	nums2Max := slices.Max(nums2)

	return nums2Max - nums1Max
}
