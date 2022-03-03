package service

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/go-vgo/robotgo"
)

var strArr = []string{}

func ReadWord() {
	strArr = []string{}
	for i := 0; i < 4 * 5; i++ {
		robotgo.KeyTap("tab")
		robotgo.MilliSleep(100)
		robotgo.KeyTap("c","ctrl")
		robotgo.MilliSleep(100)
		str, err := robotgo.ReadAll()
		if err != nil {
			fmt.Println(err.Error())
		}
c			break
		}
		strArr = append(strArr, str)
	}

	robotgo.AddEvents("enter", "alt")
	WriteCode()
}

func WriteCode()  {
	//robotgo.MilliSleep(100)
	//robotgo.AddEvent("enter")
	//robotgo.MilliSleep(100)

	for i := 0; i < len(strArr) && i + 3 < len(strArr); i += 4 {
		s1, s2, s3, s4 := strArr[i], strArr[i+1], strArr[i+2], strArr[i+3]
		robotgo.KeyPress("enter")
		s := fmt.Sprintf("%s %s `json:\"%s%s\" yy:\"%s\"`", common.ToField(s1), s2, s1, common.AllowEmpty(s3), s4)
		robotgo.TypeStr(s)
	}
	//robotgo.TypeStr("here")
	//fmt.Println(strArr)
}


