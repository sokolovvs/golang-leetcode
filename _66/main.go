package _66

func plusOne(digits []int) []int {

	var buffer = 0
	var digitPlusOne int

	for i := len(digits) - 1; i >= 0; i-- {
		digitPlusOne = digits[i] + 1

		if digitPlusOne < 10 {
			digits[i] = digitPlusOne
			buffer = 0
			break
		} else {
			digits[i] = digitPlusOne - 10
			buffer = 1
			continue
		}
	}

	if buffer == 1 {
		arr := make([]int, len(digits)+1)

		arr[0] = 1

		for i := 1; i < len(digits); i++ {
			arr[i] = digits[i]
		}

		return arr
	}

	return digits
}
