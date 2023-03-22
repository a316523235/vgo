package example

import (
	"bufio"
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/a316523235/wingo/service"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
	"os"
	"strings"
	"testing"
	"time"
)

func TestListenRMouse(t *testing.T)  {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {

		fmt.Println(e.String())
		if e.Button == hook.MouseMap["left"] {
			fmt.Println("mouse left")
		}
	})

	s := hook.Start()
	<- hook.Process(s)
}

func TestWriteByWord(t *testing.T) {
	//time.Sleep(3 * time.Second)
	robotgo.AddEvents("4", "alt")
	service.ReadWord()
}

func TestCopy(t *testing.T)  {
	time.Sleep(3 * time.Second)
	robotgo.KeyTap("c","ctrl")
	robotgo.MilliSleep(100)
	str, err := robotgo.ReadAll()
	fmt.Println(str)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestWriteEnter(t *testing.T)  {
	str := robotgo.PasteStr("123")
	fmt.Println(str)
}

func TestToField(t *testing.T)  {
	str := common.ToField("user_id")
	fmt.Println(str)
}

func TestAltEnter(t *testing.T)  {
	robotgo.SetDelay(100, 100)
	robotgo.AddEvent("enter")
	robotgo.TypeStr("here")
}

func TestAlt4(t *testing.T)  {
	hook.Register(hook.KeyDown, []string{"4", "ctrl"}, func(e hook.Event) {
		fmt.Println("alt 4")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func TestSleepOne(t *testing.T)  {
	robotgo.EventHook(hook.KeyDown, []string{"h"}, func(e hook.Event) {
		fmt.Println("h")
		time.Sleep(3 * time.Second)
	})

	robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	//kk := []string{
	//	"num_enter",
	//	"capslock",
	//	"space",
	//	"print",
	//	"printscreen",
	//}

	robotgo.EventHook(hook.KeyDown, []string{"num_enter"}, func(e hook.Event) {
		fmt.Println("num_enter")
	})
	//robotgo.EventHook(hook.KeyDown, []string{"capslock"}, func(e hook.Event) {
	//	fmt.Println("capslock")
	//})
	//robotgo.EventHook(hook.KeyDown, []string{"space"}, func(e hook.Event) {
	//	fmt.Println("space")
	//})
	//robotgo.EventHook(hook.KeyDown, []string{"print"}, func(e hook.Event) {
	//	fmt.Println("print")
	//})
	//robotgo.EventHook(hook.KeyDown, []string{"printscreen"}, func(e hook.Event) {
	//	fmt.Println("printscreen")
	//})


	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func TestDoubleKey(t *testing.T)  {
	robotgo.EventHook(hook.KeyDown, []string{"1", "2"}, func(e hook.Event) {
		fmt.Println("double 1")
	})
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func TestDoubleKey2(t *testing.T)  {
	robotgo.AddEvents("q", "ctrl")
	fmt.Println("ok")
}

func TestDoubleKey3(t *testing.T)  {
	robotgo.EventHook(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		fmt.Println("1")
		robotgo.EventEnd()
	})

	robotgo.EventHook(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		fmt.Println("2")
		go robotgo.AddEvents("q", "ctrl")
	})

	robotgo.EventHook(hook.KeyDown, []string{"3"}, func(e hook.Event) {

	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func TestOther(t *testing.T)  {
	robotgo.EventHook(hook.KeyDown, []string{"alt", "alt"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		fmt.Println(e.String())
		fmt.Printf("xx3 {%d, %d},\n\r", x, y)
	})


	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func TestTryWrite(t *testing.T)  {
	robotgo.Sleep(3)
	fmt.Println("begin")

	find, newString := "vgo", "ngo"
	findUpper, newStringUpper := "Vgo", "Ngo"

	i := 0
	x, y := robotgo.GetMousePos()
	for {
		robotgo.Sleep(1)
		robotgo.KeyTap("home")
		robotgo.KeyTap("end", "shift")
		robotgo.KeyTap("c", "ctrl")
		str, _ := robotgo.ReadAll()
		fmt.Println("读取:" + str)
		newStr := strings.ReplaceAll(str, find, newString)
		if findUpper != "" {
			fmt.Println("替换:" + newStr)
			newStr = strings.ReplaceAll(newStr, findUpper, newStringUpper)
		}
		robotgo.TypeStr(newStr)
		fmt.Println("输出完毕")
		robotgo.Sleep(3)
		println(x, y)
		robotgo.Sleep(3)
		robotgo.KeyTap("down")
		robotgo.Click()
		robotgo.Sleep(3)
		newX, newY := robotgo.GetMousePos()
		println(newX, newY)
		i++
		if newY == y {
			fmt.Println("yi yang")
			break
		}
	}
}

func TestInput(t *testing.T)  {
	reader := bufio.NewReader(os.Stdin)
	str, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("readString err, msg" + err.Error())
	}
	fmt.Println(str)
}