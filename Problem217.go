package main

func containsDuplicate(nums []int) bool {
	hm := map[int]struct{}{}

	for i := 0; i < len(nums); i++ {
		if _, ok := hm[nums[i]]; ok {
			return true
		} else {
			hm[nums[i]] = struct{}{}
		}
	}
	return false
}
