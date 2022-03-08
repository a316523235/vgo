package service

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/go-vgo/robotgo"
)

// GotoMergerPage open wait merge page
func GotoMergerPage() {
	posList := [][]int{{3822, 58, 1}, {3566, 190, 1}, {2072, 95, 1}, {2135, 394, 2}, {3613, 121, 2}, {3459, 213, 2}, {2291, 345, 2}, {2583, 298, 2}}
	for i, pos := range posList {
		if !Switch.IsTaskOpen() {
			break
		}
		robotgo.MoveClick(common.ToRightScreen(pos[0], pos[1]))
		x, y := robotgo.GetMousePos()
		fmt.Println(i, "mleft pos:", x, y)
		robotgo.Sleep(pos[2])
	}
}

// GotoMergerLastSubmitToRelease merge last submit to release branch
func GotoMergerLastSubmitToRelease() {
	//[[3398 175] [2635 259] [2635 259] [3496 308] [3349 405] [3349 405] [3385 454] [2457 432] [2404 584] [2404 584] [2404 584]]
	posList := [][]int{
		{3822, 58, 1},
		{3566, 190, 1},
		{2072, 95, 1},
		{2135, 394, 2},
		{3398, 175, 2},
		{2635, 259, 2},	// click two position
		{2700, 259, 2}, // click two position
		{3496, 308, 2},
		//{3349, 405, 2},
		{3349, 405, 2}, 	//after click here must input 'release'
		{3385, 454, 2},
		{2457, 432, 2},
	}
	for i, pos := range posList {
		if !Switch.IsTaskOpen() {
			break
		}
		robotgo.MoveClick(common.ToRightScreen(pos[0], pos[1]))
		x, y := robotgo.GetMousePos()
		fmt.Println(i, "mleft pos:", x, y)
		robotgo.Sleep(pos[2])

		if pos[0] == 3349 && pos[1] == 405 {
			//robotgo.KeyPress()
			robotgo.TypeStr("release")
			robotgo.Sleep(2)
		}
	}
}
