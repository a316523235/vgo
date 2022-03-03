package common

import (
	"strconv"
	"strings"
)

// To100 return x * 100 / 125
func To100(x, y int) (x100, y100 int) {
	return x * 100 / 125, y * 100 / 125
}

// IntJoin [1,2,3] to 1,2,3
func IntJoin(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	strArr := []string{}
	for _, d := range arr {
		strArr = append(strArr, strconv.Itoa(d))
	}
	return strings.Join(strArr, ",")
}

// AllowEmpty "否" to ",omitempty"
func AllowEmpty(wordStr string) string {
	if wordStr == "是" {
		return ""
	}
	return ",omitempty"
}

// ToField "user_id" to "UserId"
func ToFiecld(wordStr string) string {
	arr :=strings.Split(wordStr, "_")
	newArr := []string{}
	for _, str := range arr {
		newArr = append(newArr, FirstUp(str))
	}
	return strings.Join(newArr, "")
}

func FirstUp(wordStr string) string {
	vv := []rune(wordStr)
	if len(vv) > 0 && vv[0] >= 97 && vv[0] <= 122 {
		vv[0] -= 32
	}
	return string(vv)
}