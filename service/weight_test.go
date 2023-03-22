package service

import (
	"fmt"
	"math"
	"testing"
)

func TestWeightList(t *testing.T)  {
	res := babyWeightList(90)
	for i := 0; i < len(res); i++ {
		fmt.Println(fmt.Sprintf("第%d天 %f", i, res[i]))
	}
}

func TestWeightList2(t *testing.T)  {
	n := 1000.0
	m := 100.0
	p := math.Pow(1-1/m, n*(n-1)/2)
	fmt.Printf("碰撞概率为：%f\n", 1-p)
}
