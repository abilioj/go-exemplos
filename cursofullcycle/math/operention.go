package math

import "errors"

// Sum - soma dois inteito
func Sum(a int, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("numero não pode ser '0'")
	}
	res := a + b
	return res, nil
}

func SumAll(x ...int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}
