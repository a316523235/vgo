package service

func babyWeight(day int) float64 {
	w := (2.6 + float64(day)/30.0 * 0.7) * 2
	return w
}

func babyWeightList(day int) []float64 {
	res := []float64{}
	for i := 0; i <= day; i++ {
		res = append(res, babyWeight(i))
	}
	return res
}