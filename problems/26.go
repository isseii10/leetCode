package problems

// in-place
// numsは広義単調増加 ここ重要、つまり重複しているときはその数字は連続して現れる
// 最初のk個がuniqueであれば良い。元々の位置関係もkeep
// kを返す

// 自分の回答
func removeDuplicates(nums []int) int {
	n := len(nums)
	duplicated := make([]int, n)
	duplicatedMap := make(map[int]int, n)
	for i, v := range nums {
		duplicatedMap[v]++
		if duplicatedMap[v] > 1 {
			duplicated[i] = 1
		}
	}
	for k, v := range duplicated {
		if v == 1 {
			found := false
			for i := k; i < n; i++ {
				if duplicated[i] == 0 {
					found = true
					nums[k], nums[i] = nums[i], nums[k]
					duplicated[k], duplicated[i] = duplicated[i], duplicated[k]
					break
				}
			}
			if !found {
				return k
			}

		}
	}
	// ここに到達するのはすでに全てuniqueな時
	return n
}

// 早い回答
// 尺取法
// lとrの間はかぶっていることが分かっている数字になるようにする
// lより前はuniqueなのが保たれるようにする
func removeDuplicates2(nums []int) int {
	n := len(nums)
	r := 0
	ans := n // 既に全てuniqueな時
	for l := 0; l < n-1; l++ {
		for r < n {
			if nums[l] != nums[r] {
				break
			}
			r++
		}
		if r == n {
			ans = l + 1
			break
		}
		nums[l+1] = nums[r]

		if r == l {
			r++
		}
	}
	return ans
}

// 0,0,1,1,1,2,2,3,3,4
// l
// r
//
// // 1loop終わり
// 0,0,1,1,1,2,2,3,3,4
// l
//   r
//
// // 2loop終わり(違うの出てきた)
// 0,0,1,1,1,2,2,3,3,4
// l
//     r
//
// // 3loop終わり
// 0,1,1,1,1,2,2,3,3,4
//   l
//       r
//
// // 4loop終わり
// 0,1,1,1,1,2,2,3,3,4
//   l
//         r
//
// // 5loop終わり(違うの出てきた)
// 0,1,1,1,1,2,2,3,3,4
//   l
//           r
//
//
// // 6loop終わり
// 0,1,2,1,1,2,2,3,3,4
//     l
//             r
//
// // 7loop終わり
// 0,1,2,1,1,2,2,3,3,4
//     l
//               r
//
// // 7loop終わり
// 0,1,2,1,1,2,2,3,3,4
//     l
//               r
//
