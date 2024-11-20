package service

import (
	"testing"
	"time"
)

func TestNotify(t *testing.T)  {
	Notify("try test")
}

func TestStartNotify(t *testing.T) {
	StartNotify()

	for {
		time.Sleep(500 * time.Millisecond)
	}

	//location, _ := time.LoadLocation("Asia/Shanghai")
	//target := time.Date(2023, 11, 13,11,55,0, 0, location)
	//duration := target.Sub(time.Now())
	//fmt.Println(duration)

	//time.AfterFunc(duration, func() {
	//	Notify("now is 2023-11-13 11:45:00")
	//})
}

func TestStack(t *testing.T) {

	s := "[{([)}]"
	ret := isValid(s)
	if ret != true {
		t.Fatal("result error")
	}
}


func isValid(s string) bool {
	arr := []rune{}
	for _, v := range s {
		if v == '(' || v == '['|| v == '{' {
			arr = append(arr, v)
		} else {
			if len(arr) == 0 {
				return false
			} else {
				left := arr[len(arr)-1]
				if (left == '(' && v == ')') || (left == '[' && v == ']') || (left == '{' && v == '}') {
					arr = arr[:len(arr)-1]
				} else {
					return false
				}

			}
		}
	}
	return len(arr) == 0
}