package _136

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n] += 1
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	panic("Invalid input data")
}
