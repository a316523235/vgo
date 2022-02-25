package example

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/go-vgo/robotgo"
	"os"
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
	robotgo.Sleep(3)
	x,y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
}

func TestRecordClickPosition(t *testing.T)  {
	chMLeft := make(chan bool)
	chEsc := make(chan bool)
	cnt := 0
	posList := [][]int{}

	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			mleft := robotgo.AddEvent("mleft")
			chMLeft <- mleft
		}
	}()

	for {
		select {
		case <-chMLeft:
			x, y := robotgo.GetMousePos()
			cnt++
			fmt.Println(cnt ,"mleft pos:", x, y)
			posList = append(posList, []int{x,y})
			if cnt > 10 {
				fmt.Println(posList)
				os.Exit(1)
			}
		case <-chEsc:
			fmt.Println("esc over")
			os.Exit(1)
		}
	}
}

func Esc() {
	chEsc := make(chan bool)
	out := time.After(30 * time.Second)

	go func() {
		chEsc <- robotgo.AddEvent("mleft")
	}()

	select {
	case <-chEsc:
		fmt.Println("esc over")
		os.Exit(1)
	case <-out:
		fmt.Println("timeout over")
		os.Exit(1)
	}
}

func TestGotoMergerPage(t *testing.T) {
	// open wait merge
	posList := [][]int{{3822,58},{3566,190},{2072,95},{2135,394},{3613,121},{3459,213},{2291,345},{2583,298}}
	for i, pos := range posList {
		robotgo.MoveClick(common.To100(pos[0], pos[1]))
		x, y := robotgo.GetMousePos()
		fmt.Println(i ,"mleft pos:", x, y)
		robotgo.Sleep(1)
	}
}



