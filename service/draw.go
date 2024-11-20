package service

import "strings"

type dayInfo struct {
	Day string
	Numbers []numberInfo
}

type numberInfo struct {
	Name string
	Num int
}


func createDrawConfig(string)  {
	awardStr := `序号	中奖奖品（展示内容）	奖品份数	12月15日	12月16日	12月17日	12月18日	12月19日	12月20日	12月21日
1	京东卡 500元	1				1			
2	京东卡 100元	2	1				1		
3	京东卡 20元	30	8	3	3	6	5	4	1
4	京东卡 10元	50	15	5	5	8	8	8	1
5	周十五蜂蜜露	10	2	1	1	2	2	2	
6	安满智孕宝孕妈奶粉	5	1	1		1	1	1	
7	宜婴红花山茶拉拉裤试用装6片  XL	5	1	1		1	1	1	
8	启初婴儿多维舒缓护臀膏 60g	10	2	1	1	2	2	1	1
9	孕味食足酸酸片	10	2	1	1	2	2	1	1
10	嫩芙温和柔肤身体精华油110ml	10	2	1	1	2	2	1	1
11	丝塔芙大白罐 250g	10	2	1	1	2	2	1	1
12	Hydromol孩舒抚湿疹沐浴露50ml	10	2	1	1	2	2	1	1
13	金秀儿艾灸贴+足浴包	10	2	1	1	2	2	1	1
14	博益贝贝舒缓凝露正装 28g	10	2	1	1	2	2	1	1
15	三诺爱看动态血糖仪	5	1	1		1	1	1	
16	乐扣乐扣人鱼线保温杯	5	1	1		1	1	1	
17	肌肤蕾海盐水喷雾	5	1		1	1	1	1	
18	Noromega虾青素胶囊	5	1		1	1	1	1	
19	Yep保温袋	3				1	1	1	
20	Yep子母包	1	1						
21	Yep大杯子	1				1
`

lines := strings.Split(awardStr, "\n")
lineArr := [][]string{}
	for _, line := range lines  {
		lineArr = append(lineArr, strings.Split(line, "	"))
	}

	res := []dayInfo{}
	rowCnt, colCnt := len(lineArr), len(lineArr[0])	//行数，列数
	// 案列循环
	for j := 3; j < colCnt; j++ {
		temp := dayInfo{
			Day:lineArr[0][j],
		}
	}
}
