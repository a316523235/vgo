package common

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"strconv"
	"strings"
)

//{0, 0}, {1535, 0}, {1920, 0}, {3839, 0},
//{0, 863}, {1535, 863}, {1920, 1079}, {3839, 1079},

// record pos, left and right all use 1.25 rate
// so if use in left x,y = x * LeftRate / RecordRate, y * LeftRate / RecordRate,
// so if use in right x,y = x * RightRate / RecordRate, y * RightRate / RecordRate,
const RecordRate  = float64(1.25)

const LeftRate  =  float64(1.25)

const RightRate  = float64(1.0)

func GetCurrentScreenRate() float64 {
	w, h := robotgo.GetScaleSize()
	sW, sH := robotgo.GetScreenSize()
	fmt.Println("GetScaleSize:", w, h, "GetScreenSize:",  sW, sH)
	rate := float64(h) / float64(sH)
	return rate
}

func GetAutoXy(x, y int) (cX, cY int) {
	rate := GetCurrentScreenRate()
	fX, fY := float64(x) * rate / RecordRate , float64(y) * rate / RecordRate
	cX, cY = int(fX), int(fY)
	return cX, cY
}

func GetLeftXy(x, y int) (cX, cY int)  {
	fX, fY := float64(x) * LeftRate / RecordRate , float64(y) * LeftRate / RecordRate
	cX, cY = int(fX), int(fY)
	return cX, cY
}

func GetRightXy(x, y int) (cX, cY int)  {
	fX, fY := float64(x) * RightRate / RecordRate , float64(y) * RightRate / RecordRate
	cX, cY = int(fX), int(fY)
	return cX, cY
}

func ToAuto100(x, y int) (x100, y100 int) {
	w, h := robotgo.GetScaleSize()
	sW, sH := robotgo.GetScreenSize()
	fmt.Println("GetScaleSize:", w, "*", h, "  GetScreenSize:", sW , "*", sH)

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