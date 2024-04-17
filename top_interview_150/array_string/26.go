package arraystring

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
// 二つのポインタを走査する
func removeDuplicates_2(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			// もしslow = fast + 1の場合でここに入った時は、nums[slow]は重複していない。
			// それ以外の場合はslow以降が重複しているので、incrementしてnums[fast]を入れてあげる
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
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
