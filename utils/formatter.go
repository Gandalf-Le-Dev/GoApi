package utils

import "strconv"

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	PanicIfErr(err)
	return num
}
