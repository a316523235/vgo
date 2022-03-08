package service

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/go-vgo/robotgo"
	"time"
)

var strArr = []string{}

func ReadWord() {
	time.Sleep(1 * time.Second)
	strArr = []string{}
	for i := 0; Switch.IsTaskOpen() && i < 4 * 20; i++ {
		robotgo.KeyTap("tab")
		robotgo.MilliSleep(100)
		robotgo.KeyTap("c","ctrl")
		robotgo.MilliSleep(100)
		str, err := robotgo.ReadAll()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// empty || current str == prev str
		if str == "" || (i > 0 && i % 4 == 3 && str == strArr[i - 1] && str == strArr[i-2] && str == strArr[i-3]) {
			break
		}
		strArr = append(strArr, str)
		robotgo.MilliSleep(500)
	}

	fmt.Println("read over, please press \"enter\"")
	robotgo.AddEvents("enter")
	WriteCode()
}

func WriteCode()  {
	//robotgo.MilliSleep(100)
	//robotgo.AddEvent("enter")
	//robotgo.MilliSleep(100)

	for i := 0; Switch.IsTaskOpen()  && i < len(strArr) && i + 3 < len(strArr); i += 4 {
		s1, s2, s3, s4 := strArr[i], strArr[i+1], strArr[i+2], strArr[i+3]
		robotgo.KeyPress("enter")
		s := fmt.Sprintf("%s %s `json:\"%s%s\" yy:\"%s\"`", common.ToField(s1), s2, s1, common.AllowEmpty(s3), s4)
		robotgo.TypeStr(s)
		robotgo.MilliSleep(500)
	}
	//robotgo.TypeStr("here")
	fmt.Println("---over---")
}


