package utils

import "strconv"

func ConvertUint(value string) (uint, error) {
	u64, err := strconv.ParseUint(value, 10, 32)

	if err != nil {
		return 1, err
	}

	return uint(u64), err
}