package example

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/a316523235/wingo/service"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"testing"
	"time"
)

func TestMouse1(t *testing.T) {
	robotgo.MoveMouse(100, 100)
}

func TestMouse2(t *testing.T) {
	// 慢慢移动
	robotgo.MoveMouseSmooth(100, 200)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}

func TestMouse3(t *testing.T) {
	// 单击、双击；合起来变为3击， 就会选中一行
	robotgo.MouseClick()	//单击
	robotgo.MouseClick("left", true)	//双击

}

func TestMouse5(t *testing.T) {
	robotgo.MoveClick(15, 30)
	//robotgo.MoveClick(15, 30, "left", true)
}

func TestMouse6(t *testing.T) {
	// 切换按下或抬起状态 (会一直保持着)
	//robotgo.MouseToggle("down")
	robotgo.MouseToggle("down", "left")
}

func TestMouse7(t *testing.T) {
	//从500,400鼠标按下并保存按下状态划到500,500
	robotgo.MoveMouse(500, 400)
	robotgo.MouseToggle("down")
	robotgo.DragMouse(500, 500)
	robotgo.MouseToggle("up")
}

func TestMouse8(t *testing.T) {
	// get mouse position
	robotgo.Sleep(3)
	x,y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
}

func TestRecordClickPosition(t *testing.T)  {
	service.RecordClickPosition()
}

func TestRecordClickPosition2(t *testing.T)  {
	service.RecordClickPositionV2()
}

func TestGotoMergerPage(t *testing.T) {
	service.GotoMergerPage()
}

func TestListenKeys(t *testing.T) {
	s := hook.Start()

	for i := 0; i < 10; {
		e := <-s
		if e.Kind == hook.KeyHold  {
			i++
			fmt.Println(e.Keycode)
		}
	}
	//
	//chKey1, chKey2, chKey3 := make(chan bool), make(chan bool), make(chan bool)
	//out := time.After(30 * time.Second)
	//
	//go func() {
	//	chKey3 <- robotgo.AddEvents("q", "ctrl", "shift")
	//}()
	//
	//go func() {
	//	chKey2 <- robotgo.AddEvents("w", "ctrl", "shift")
	//}()
	//
	//
	//select {
	//case <-chKey1:
	//	fmt.Println("num 1")
	//case <-chKey2:
	//	fmt.Println("num 2")
	//case <-chKey3:
	//	fmt.Println("num 3")
	//case <-out:
	//	fmt.Println("time out")
	//}
}

func TestGoHook(t *testing.T)  {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func TestGotoMergerLastSubmitToRelease(t *testing.T)  {
	service.GotoMergerLastSubmitToRelease()
}

func TestMouseDwon(t *testing.T)  {
	robotgo.ScrollMouse(500, "down")
}
func TestBooking(t *testing.T)  {
	time.Sleep(2 * time.Second)
	service.Booking()
}

func TestTemp1(t *testing.T)  {
	//{0, 0}, {1535, 0}, {1920, 0}, {3839, 0},
	//{0, 863}, {1535, 863}, {1920, 1079}, {3839, 1079},
	robotgo.Sleep(2)
	x, y := common.GetRightXy(3822,58)
	fmt.Println(x, y)
	robotgo.Move(x, y)
}

