package service

/*
//#if defined(IS_MACOSX)
	#cgo darwin CFLAGS: -x objective-c -Wno-deprecated-declarations
	#cgo darwin LDFLAGS: -framework Cocoa -framework OpenGL -framework IOKit
	#cgo darwin LDFLAGS: -framework Carbon -framework CoreFoundation
//#elif defined(USE_X11)
	#cgo linux CFLAGS: -I/usr/src
	#cgo linux LDFLAGS: -L/usr/src -lX11 -lXtst -lm
	// #cgo linux LDFLAGS: -lX11-xcb -lxcb -lxcb-xkb -lxkbcommon -lxkbcommon-x11
//#endif
	// #cgo windows LDFLAGS: -lgdi32 -luser32 -lpng -lz
	#cgo windows LDFLAGS: -lgdi32 -luser32
// #include <AppKit/NSEvent.h>
#include "screen/goScreen.h"
#include "mouse/goMouse.h"
#include "key/goKey.h"
//#include "event/goEvent.h"
#include "window/goWindow.h"
*/

import "C"
import (
	"bufio"
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ReplaceCode()  {
	//win2 := robotgo.GetActive()
	reader := bufio.NewReader(os.Stdin)
	//wingoIds, err := robotgo.FindIds("wingo.exe")
	//if err != nil {
	//	fmt.Println("find wingo.exe pid error, msg: " + err.Error())
	//}
	//wingoPid := wingoIds[0]
	//err = robotgo.ActivePID(wingoPid)
	//if err != nil {
	//	fmt.Println("wingo.exe active error, msg: " + err.Error())
	//}
	fmt.Println("please input find string")
	bt, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("readString err, msg" + err.Error())
	}
	find := string(bt)
	fmt.Println("please input replace string")
	bt, _, err = reader.ReadLine()
	newString := string(bt)

	findUpper, newStringUpper := firstCharUpper(find), firstCharUpper(newString)
	

	robotgo.Sleep(1)
	robotgo.SetActive(CurrentWin)
	robotgo.MoveClick(CurrentPosX, currentPoxY)
	robotgo.Sleep(1)
	robotgo.TypeStr("begin")
	robotgo.KeyTap("enter")


	i := 0
	robotgo.KeyTap("insert")
	for {
		_, y := robotgo.GetMousePos()
		robotgo.Sleep(2)
		robotgo.KeyTap("home")
		robotgo.Sleep(2)
		robotgo.KeyTap("end", "shift")
		robotgo.Sleep(3)
		robotgo.KeyTap("c", "ctrl")
		robotgo.Sleep(2)
		str, _ := robotgo.ReadAll()
		robotgo.Sleep(2)
		fmt.Println("读取:" + str)
		newStr := strings.ReplaceAll(str, find, newString)
		robotgo.Sleep(2)
		if findUpper != "" {
			fmt.Println("替换:" + newStr)
			newStr = strings.ReplaceAll(str, findUpper, newStringUpper)
		}
		robotgo.TypeStr("new new " + newStr)
		robotgo.Sleep(2)
		//robotgo.KeyTap("down")
		robotgo.Sleep(2)
		_, newY := robotgo.GetMousePos()
		i++
		if i > 10 {
			break
		}
		if newY == y {
			fmt.Println("yi yang")
		}
	}

	robotgo.KeyTap("insert")

	fmt.Println("---over---")
	robotgo.KeyTap("esc")
}

func showWingo()  {
	//wingoIds, err := robotgo.FindIds("wingo.exe")
	//if err != nil {
	//	fmt.Println("find wingo.exe pid error, msg: " + err.Error())
	//}
	//wingoPid := wingoIds[0]
	//err = robotgo.ActivePID(wingoPid)

	err := robotgo.ActiveName("wingo")
	if err != nil {
		fmt.Println("find wingo.exe pid error, msg: " + err.Error())
	}
}

func firstCharUpper(str string) string {
	r, size := utf8.DecodeRuneInString(str)
	if unicode.IsLetter(r) {
		firstChar := string(unicode.ToUpper(r))
		result := firstChar + str[size:]
		return result
	}
	return ""
}
