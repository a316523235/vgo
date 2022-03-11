package common

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"strconv"
	"strings"
)

// To100to125 return x * 100 / 125
func To100to125(x, y int) (x100, y100 int) {
	// 左屏缩放125%时， GetScaleSize:4800 1350， TestGetScaleSize:3840 1080
	// 右屏缩放100%时， GetScaleSize:3840 1350， TestGetScaleSize:3840 1080
	_, h := robotgo.GetScaleSize()
	_, sH := robotgo.GetScreenSize()
	//return x * 100 / 125, y * 100 / 125
	return x * sH / h, y * sH / h
}

// To100to125 return x * 100 / 100
func ToLeftScreen(x, y int) (x100, y100 int) {
	return x * 100 / 100, y * 100 / 100
}

// GetRealPx return x * 125 / 100
func GetRealPx(x, y int) (x100, y100 int)   {
	_, h := robotgo.GetScaleSize()
	_, sH := robotgo.GetScreenSize()

	fmt.Println(h, sH)

	return x * h / sH, y * h / sH
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
func ToField(wordStr string) string {
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