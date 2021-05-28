package util

import (
	"strconv"
)

func StrToUint32(str string) (uint32, error) {
	i64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	i32 := uint32(i64)
	return i32, nil
}

func Uint32ToStr(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}
