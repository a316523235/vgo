package service

import (
	"fmt"
	"github.com/a316523235/wingo/conf"
	"github.com/go-vgo/robotgo"
	"strings"
)

var currentLine = 0
var lineArr = []string{}

func ReadEn()  {
	if len(lineArr) == 0 {
		lineArr = strings.Split(conf.EN1, "\n")
		currentLine = 0
	}

	for ; currentLine < len(lineArr); currentLine++ {
		if !Switch.IsTaskOpen() {
			break
		}
		str := lineArr[currentLine]
		//robotgo.TypeStr(str + "\n")
		fmt.Println(str + "\n")
		robotgo.Sleep(15)
	}
}


