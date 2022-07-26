package service

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/go-vgo/robotgo"
)

type auth struct {
	Code string
	Name string
}

func AddDepartmentAuth() {
	// record by 125%
	//str, err := robotgo.ReadAll()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//authCode, authName := "yj/consumestat/index", "财务管理-消费统计(柚+)"
	authCode, authName := "yj/newly/index", "柚+新开客户统计"
	departmentList := []auth{
		{"admin", "不限"},
		{"east", "华东"},
		{"south", "华南"},
		{"north", "华北"},
		{"southwest_test", "西南测试"},
		{"test", "测试"},
	}

	posList := [][4]int{
		//{2691, 326, 3}, //add button
		{3517, 452, 2}, //add button
		{2615, 369, 1, 101},      //auth name
		{2616, 420, 1, 102},      //auth code
		{2576, 456, 1, 103},      //auth url
		{2589, 500, 1},      //auth type click
		{2582, 664, 1},      //auth type select
		{2583, 547, 1},      //auth level click
		{2589, 606, 1},      //auth level select
		{2572, 640, 1},      //is online
		{2648, 684, 1},      //is show
		{2648, 730, 1},      //is dir
		{3295, 790, 1},      //click save
	}

	for _, department := range departmentList {
		for i, pos := range posList {
			if !Switch.IsTaskOpen() {
				break
			}
			robotgo.MoveClick(common.GetRightXy(pos[0], pos[1]))
			x, y := robotgo.GetMousePos()
			fmt.Println(i, "mleft pos:", x, y)
			robotgo.Sleep(pos[2])

			switch pos[3] {
			case 101:
				str := authName + " - 渠道权限 - " + department.Name
				fmt.Println(str)
				robotgo.Click()
				robotgo.Click("left", true)
				robotgo.TypeStr(str)
				robotgo.Sleep(2)
			case 102:
				code := authCode + "/" + department.Code
				fmt.Println(code)
				robotgo.Click()
				robotgo.Click("left", true)
				robotgo.TypeStr(code)
				robotgo.Sleep(2)
			case 103:
				uri := authCode + "/" + department.Code
				fmt.Println(uri)
				robotgo.Click()
				robotgo.Click("left", true)
				robotgo.TypeStr(uri)
				robotgo.Sleep(2)
			}
		}
		robotgo.Sleep(4)
	}
}
