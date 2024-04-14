package bitmanipulation

func singleNumber(nums []int) int {
	ret := 0
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = true
			ret += v
			continue
		}
		ret -= v
	}
	return ret
}

// xor
func singleNumber2(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}
