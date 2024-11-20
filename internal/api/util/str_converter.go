package util

import (
	"fmt"
	"strconv"
)

func ConvertStringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Error converting %s to int from setting env: %s", s, err))
	}
	return num
}
