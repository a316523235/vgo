package example

import (
	"fmt"
	"github.com/a316523235/wingo/common"
	"strings"
	"testing"
)

func Test100To125(t *testing.T)  {
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
	newPosList := [][]int{}

	strList := []string{}
	for _, arr := range posList {
		newArr := []int{arr[0] * 125 / 100, arr[1] * 125 / 100, arr[2]}
		newPosList = append(newPosList, newArr)

		strList = append(strList, "{" + common.IntJoin(newArr) + "}")
	}
	res := "{" + strings.Join(strList, ",") + "}"
	fmt.Println(res)
}