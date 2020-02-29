package array

func Sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func SumAll(nums ...[]int) (totalArr []int) {
	for _, v := range nums {
		total := Sum(v)
		totalArr = append(totalArr, total)
	}

	return totalArr
}
