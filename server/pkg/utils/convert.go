package utils

import (
	"strconv"
)

func ToInt(n string) int {
	res, err := strconv.ParseInt(n, 10, 32)
	if err != nil {
		return 0
	}
	return int(res)
}
