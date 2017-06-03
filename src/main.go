package main

import "./myUtil"

func main() {
	x := make([]float64, 100)
	y := make([]float64, 100)
	for i := 0; i < len(x); i++ {
		x[i] = float64(i)
		y[i] = 2 * x[i]
	}

	myUtil.MyPlot(x, y)

}
