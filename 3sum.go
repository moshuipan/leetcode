package leetcode

func threeSum(nums []int) [][]int {
	m := make(map[int]int) //value: index,已经处理过的一级解
	type sencond struct {
		m map[int]int
	}
	m2 := make(map[int]sencond) //已经获取过的二级解
	var tmp [][]int
	//
	for i := 0; i < len(nums); i++ {
		if _, ok := m[-nums[i]]; ok { //该解已经处理过了
			continue
		}
		m[-nums[i]] = nums[i]
		other := make(map[int]int, len(nums)-i)
		for j := i + 1; j < len(nums); j++ {
			sum := -nums[i] - nums[j]
			if v, ok := other[sum]; ok {
				//去掉重复解
				v2, ok := m2[v]
				v21, ok1 := v2.m[nums[i]]
				v22, ok2 := v2.m[nums[j]]
				if !ok || ((v21 != nums[j] && ok1) && (v22 != nums[i] && ok2)) || (!ok1 && !ok2) {
					tmp = append(tmp, []int{nums[i], nums[j], v})
					if _, ok := m2[v]; !ok {
						m2[v] = sencond{m: map[int]int{}}
					}
					if m2[nums[i]].m == nil {
						m2[nums[i]] = sencond{m: map[int]int{}}
					}
					if m2[nums[j]].m == nil {
						m2[nums[j]] = sencond{m: map[int]int{}}
					}
					m2[v].m[nums[i]] = nums[j]
					m2[v].m[nums[j]] = nums[i]

					m2[nums[i]].m[v] = nums[j]
					m2[nums[i]].m[nums[j]] = v

					m2[nums[j]].m[v] = nums[i]
					m2[nums[j]].m[nums[i]] = v
				}
			}
			other[nums[j]] = nums[j]
		}
	}
	return tmp
}
