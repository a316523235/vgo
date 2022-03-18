package example

import (
	"errors"
	"fmt"
	"github.com/a316523235/wingo/common"
	"github.com/a316523235/wingo/service"
	"github.com/go-vgo/robotgo"
	"strconv"
	"testing"
	"time"
)

type Trip struct {
	Icon []int
	FilePath string
}

const SavePath = "d:/hzj/bitmap/"

func (t *Trip) getScreen() {
	robotgo.AddEvent("mleft")
	x, y := robotgo.GetMousePos()
	t.Icon[0], t.Icon[1] = x, y
	robotgo.AddEvent("mleft")
	x, y = robotgo.GetMousePos()
	t.Icon[2], t.Icon[3] = x - t.Icon[0], y - t.Icon[1]
}

func (t *Trip) getShot() {
	bitMap := robotgo.CaptureScreen(t.Icon...)
	fileName := time.Now().Second()
	filePath := SavePath + strconv.Itoa(fileName) + ".tif"
	robotgo.SaveBitmap(bitMap, filePath)
	fmt.Println("has save bitmap")
}

func (t *Trip) FindPos() (err error, x, y int){
	bitMap := robotgo.OpenBitmap(t.FilePath)
	x, y = robotgo.FindBitmap(bitMap)
	if x * y < 0 {
		fmt.Println("find bitmap fail")
		return errors.New("find bitmap fail"), x, y
	}
	return nil, x, y
}

func TestGetScreen(t *testing.T)  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("please click first x y")
	robotgo.AddEvent("alt")
	x, y := robotgo.GetMousePos()
	fmt.Println("please click second x y")
	robotgo.Sleep(1)
	robotgo.AddEvent("alt")
	x2, y2 := robotgo.GetMousePos()
	w, h := x2 - x, y2 - y
	icon := []int{x, y, w, h}
	fmt.Println("has get icon:", x, y, w, h)

	robotgo.Sleep(1)

	bitMap := robotgo.CaptureScreen(icon...)
	filePath := SavePath + "test1.tif"
	res := robotgo.SaveBitmap(bitMap, filePath)
	fmt.Println("has save bitmap", res)

	newBitMap := robotgo.OpenBitmap(filePath)
	x, y = robotgo.FindBitmap(newBitMap)
	if x * y < 0 {
		fmt.Println("find bitmap fail x y：", x, y)
	}
	robotgo.MoveClick(x, y, "mleft", true)
}

func TestSaveIcon(t *testing.T)  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("please click first x y")
	robotgo.AddEvent("alt")
	x, y := robotgo.GetMousePos()
	x, y = common.GetRealPx(x, y)
	robotgo.Sleep(1)
	fmt.Println("please click second x y")
	robotgo.AddEvent("alt")
	x2, y2 := robotgo.GetMousePos()
	x2, y2 = common.GetRealPx(x2, y2)
	w, h := x2 - x, y2 - y
	icon := []int{x, y, w, h}
	fmt.Println("has get icon", x, y, w, h)

	robotgo.Sleep(1)

	bitMap := robotgo.CaptureScreen(icon...)
	filePath := SavePath + "test2.tif"
	res := robotgo.SaveBitmap(bitMap, filePath)
	fmt.Println("has save bitmap", res)
}

func TestOpenIcon(t *testing.T) {
	filePath := SavePath + "test2.tif"
	newBitMap := robotgo.OpenBitmap(filePath)
	x, y := robotgo.FindBitmap(newBitMap)
	fmt.Println("find bitmap fail x y：", x, y)
	if x  < 0 || y < 0 {
		fmt.Println("error x y")
		return
	}
	//465 501 67 74
	robotgo.Sleep(1)
	robotgo.MoveClick(x, y, "mleft")
}

func TestOpenIconByAltA(t *testing.T) {
	filePath := SavePath + "test3.png"
	newBitMap := robotgo.OpenBitmap(filePath)
	x, y := robotgo.FindBitmap(newBitMap)
	fmt.Println("find bitmap fail x y：", x, y)
	if x  < 0 || y < 0 {
		fmt.Println("error x y")
		return
	}
	//465 501 67 74
	robotgo.Sleep(1)
	robotgo.MoveClick(x, y, "mleft")
}

func TestSendMsg(t *testing.T)  {
	robotgo.Sleep(1)
	filePath := SavePath + "t1.jpg"
	newBitMap := robotgo.OpenBitmap(filePath)
	x, y := robotgo.FindBitmap(newBitMap)
	fmt.Println("find bitmap fail x y：", x, y)
	if x  < 0 || y < 0 {
		fmt.Println("error x y")
		return
	}
	robotgo.Sleep(1)
	robotgo.MoveClick(x, y, "mleft")

	// 590 360
	robotgo.MoveClick(x + 590, y + 360, "mleft")

	robotgo.TypeStr("你好")

	filePath = SavePath + "t2.png"
	newBitMap = robotgo.OpenBitmap(filePath)
	x, y = robotgo.FindBitmap(newBitMap)
	fmt.Println("find bitmap fail x y：", x, y)
	if x  < 0 || y < 0 {
		fmt.Println("error x y")
		return
	}
	//robotgo.MoveClick(x, y, "mleft")
}

func TestCBitmap(t *testing.T)  {
	bmp, free := loadBitmaps("t1.png")
	defer free()

	for {
		clickBitmap(bmp["t1.png"], false)
		//clickBitmap(bmp["chest.png"], true)
		//clickBitmap(bmp["eat.png"], false)
	}
}

func loadBitmaps(files ...string) (bitmaps map[string]robotgo.Bitmap, free func()) {
	freeFuncs := make([]func(), 0)
	bitmaps = make(map[string]robotgo.Bitmap)
	for _, f := range files {
		bitmap, freeFunc := readBitmap(SavePath +  f)
		bitmaps[f] = bitmap
		freeFuncs = append(freeFuncs, freeFunc)
	}

	free = func() {
		for key := range freeFuncs {
			freeFuncs[key]()
		}
	}
	return bitmaps, free
}

func readBitmap(file string) (bitmap robotgo.Bitmap, free func()) {
	cBitmap := robotgo.OpenBitmap(file)
	bitmap = robotgo.ToBitmap(cBitmap)
	free = func() {
		robotgo.FreeBitmap(cBitmap)
	}
	return bitmap, free
}

func clickBitmap(bmp robotgo.Bitmap, doubleClick bool) bool {
	fx, fy := robotgo.FindBitmap(robotgo.ToCBitmap(bmp))
	if fx != -1 && fy != -1 {
		robotgo.MoveMouse(fx, fy)
		robotgo.MouseClick("left", doubleClick)
		return true
	}

	return false
}

func TestFindBitMapXy(t *testing.T)  {
	robotgo.Sleep(2)
	x, y, err := service.FindBitMapXy("merge_assignee.png")
	fmt.Println(x, y, err)
	if err != nil {
		return
	}
	robotgo.Sleep(2)

	x, y = x + 150, y + 20
	fmt.Println(x, y)
	robotgo.MoveClick(common.GetRightXy(x, y))
	robotgo.Sleep(2)

	x, y = x, y - 275
	fmt.Println(x, y)
	robotgo.MoveClick(common.GetRightXy(x, y))
	robotgo.Sleep(2)
	//return

	// check user
	userName := "liq"
	projectMap := map[string]string{
		"merge_go-mye.png":           "lins",
		"merge_adx.png":          "lins",
		"merge_go-advertise.png": "lins",
	}
	for project, tempName  := range projectMap {
		_, _, err = service.FindBitMapXy(project)
		if err == nil {
			userName = tempName
			break
		}
	}

	robotgo.TypeStr(userName)
	robotgo.Sleep(2)

	x, y = x, y + 50
	fmt.Println(x, y)
	robotgo.MoveClick(common.GetRightXy(x, y))
	robotgo.Sleep(2)



	//2467 411 <nil>
	// 2613 434
	// 2616 532
	// 2680 586


	//2657 460
}