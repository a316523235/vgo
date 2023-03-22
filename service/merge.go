package service

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/go-vgo/robotgo"
)

// GotoMergerPage open wait merge page
func GotoMergerPage() {
	// record by 125%
	posList := [][]int{{3822, 58, 1}, {3566, 190, 1}, {2072, 95, 1}, {2135, 394, 2}, {3613, 121, 2}, {3459, 213, 2}, {2291, 345, 2}, {2583, 298, 2}}
	for i, pos := range posList {
		if !Switch.IsTaskOpen() {
			break
		}
		robotgo.MoveClick(common.GetRightXy(pos[0], pos[1]))
		x, y := robotgo.GetMousePos()
		fmt.Println(i, "mleft pos:", x, y)
		robotgo.Sleep(pos[2])
	}
}

// GotoMergerLastSubmitToRelease merge last submit to release branch
func GotoMergerLastSubmitToRelease() {
	//[[3398 175] [2635 259] [2635 259] [3496 308] [3349 405] [3349 405] [3385 454] [2457 432] [2404 584] [2404 584] [2404 584]]
	// record by 125%
	posList := [][4]int{
		{3822, 58, 1},
		{3566, 190, 1},
		{2072, 95, 1},
		{2135, 394, 2},
		{3398, 175, 2},
		{2635, 259, 2},	// click two position
		{2700, 259, 2}, // click two position
		{3496, 308, 3},
		//{3349, 405, 2},
		{3349, 405, 3, 101}, 	//after click here must input 'release'
		{3385, 444, 2},
		{2457, 434, 2},		//click two position
		{2457, 444, 2},		//click two position
		// not to select user
		//{2457, 454, 2, 102},//click two position, after click find Assignee and select user
	}
	for i, pos := range posList {
		if !Switch.IsTaskOpen() {
			break
		}
		robotgo.MoveClick(common.GetRightXy(pos[0], pos[1]))
		x, y := robotgo.GetMousePos()
		fmt.Println(i, "mleft pos:", x, y)
		robotgo.Sleep(pos[2])

		if pos[3] == 101 {
			//robotgo.KeyPress()
			robotgo.TypeStr("release")
			robotgo.Sleep(3)
		}

		if pos[3] == 102 {
			// find assignee.png x,y
			x, y, err := FindBitMapXy("merge_assignee.png")
			fmt.Println(x, y, err)
			if err != nil {
				//{2651, 737, 2},	if not find, use default position
				x, y = 2651, 730
				//return
			}
			robotgo.Sleep(2)

			// move to input
			x, y = x + 150, y + 20
			fmt.Println(x, y)
			robotgo.MoveClick(common.GetRightXy(x, y))
			robotgo.Sleep(2)

			// move to search and input
			x, y = x, y - 275
			fmt.Println(x, y)
			robotgo.MoveClick(common.GetRightXy(x, y))
			robotgo.Sleep(2)

			// check user
			userName := "liq"
			projectMap := map[string]string{
				"merge_go-mye.png":       "lins",
				"merge_adx.png":          "lins",
				"merge_go-advertise.png": "lins",
			}
			for project, tempName  := range projectMap {
				_, _, err = FindBitMapXy(project)
				if err == nil {
					userName = tempName
					break
				}
			}

			robotgo.TypeStr(userName)
			robotgo.Sleep(2)

			// select user
			x, y = x, y + 50
			fmt.Println(x, y)
			robotgo.MoveClick(common.GetRightXy(x, y))
			robotgo.Sleep(2)
		}
	}
}
