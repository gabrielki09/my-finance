package utils

func Ternary[T any](condition bool, ifVal, elseVal T) T {
	if condition {
		return ifVal
	}

	return elseVal
}
