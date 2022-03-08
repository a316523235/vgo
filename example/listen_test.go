package example

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/a316523235/wingo/service"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
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