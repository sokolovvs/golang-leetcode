package _67

import (
	"math"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	return _sumProcessed(_preprocessArgs(a, b))
}

func _preprocessArgs(a, b string) (string, string) {
	var tmpStrArr []string
	lenA := len(a)
	lenB := len(b)
	diff := int(math.Abs(float64(lenA) - float64(lenB)))

	if lenA != lenB {
		for i := 0; i < diff; i++ {
			tmpStrArr = append(tmpStrArr, "0")
		}

		if lenA > lenB {
			b = strings.Join(tmpStrArr, "") + b
		} else {
			a = strings.Join(tmpStrArr, "") + a
		}
	}

	return a, b
}

func _sumProcessed(a, b string) string {
	var buffer = 0

	l := len(a)
	resultArr := make([]string, l+1)

	for i := l - 1; i > -1; i-- {
		ai, _ := strconv.Atoi(string(a[i]))
		bi, _ := strconv.Atoi(string(b[i]))
		sumResult := ai + bi + buffer
		sumResultStr := strconv.Itoa(sumResult)

		if sumResult > 1 {
			buffer = 1

			if sumResult == 3 {
				resultArr[i+1] = "1"
			} else {
				resultArr[i+1] = "0"
			}
		} else {
			buffer = 0
			resultArr[i+1] = sumResultStr
		}
	}

	if buffer == 1 {
		resultArr[0] = "1"
	}

	return strings.Join(resultArr, "")
}
