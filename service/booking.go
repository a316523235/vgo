package service

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/go-vgo/robotgo"
)

func Booking()  {
	//[[1194 842] [451 129] [480 219] [749 458] [804 534] [786 471] [734 532] [616 609] [458 748] [746 688] [1088 473] [1088 473] [1088 473]]
	posList := [][]int{
		{1194,842,1},
		{480,130,1},// click to search
		{480,219,1},
		{749,458,1},
		{804,534,1},	// after right click
		{786,471,1},
		{734,532,1},
		{616,609,1},	// input name
		{458,748,1},	// input sn
		{746,688,1},	// scroll down
		{1088,473,1},
		{1088,473,1},
		{1088,473,1},
	}
	for i, pos := range posList {
		if !Switch.IsTaskOpen() {
			break
		}
		robotgo.MoveClick(common.GetLeftXy(pos[0], pos[1]))
		x, y := robotgo.GetMousePos()
		fmt.Println(i, "mleft pos:", x, y)
		robotgo.Sleep(pos[2])

		if pos[0] == 480 && pos[1] == 130 {
			//robotgo.KeyPress()
			robotgo.TypeStr("美柚餐")
			robotgo.Sleep(2)
		} else if pos[0] == 804 && pos[1] == 534 {
			//robotgo.KeyPress()
			robotgo.Click("right")
			robotgo.Sleep(2)
		} else if pos[0] == 616 && pos[1] == 609 {
			//robotgo.KeyPress()
			robotgo.TypeStr("黄智健")
			robotgo.Sleep(2)
		}
	}
}