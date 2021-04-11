package _66

func plusOne(digits []int) []int {
	var buffer = false
	var digitPlusOne int

	for i := len(digits) - 1; i >= 0; i-- {
		digitPlusOne = digits[i] + 1

		if digitPlusOne < 10 {
			digits[i] = digitPlusOne
			buffer = false
			break
		} else {
			digits[i] = digitPlusOne - 10
			buffer = true
			continue
		}
	}

	if buffer == true {
		return append([]int{1}, digits...)
	}

	return digits
}
