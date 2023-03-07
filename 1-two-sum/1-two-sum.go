package leetcode

// func twoSum(nums []int, target int) []int {
// 	n := len(nums)
// 	for i := 0; i < n-1; i++ {
// 		for j := i + 1; j < n; j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	imap := make(map[int]int)
    // index = walked indeger, value = index of value
	for currI, currNum := range nums {
		if requiredI, ok := imap[target-currNum]; ok { // checking if we already have required value
			return []int{requiredI, currI} // then return pair of found index and cuurend
		}
		imap[currNum] = currI
	}
	return []int{}
}
