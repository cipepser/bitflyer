package myUtil

import (
	"image/color"
	"log"
	"os/exec"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
)

var (
	PositiveLineStyle = draw.LineStyle{
		Color:    color.Black,
		Width:    vg.Points(1),
		Dashes:   []vg.Length{},
		DashOffs: 0,
	}

	NegativeLineStyle = draw.LineStyle{
		Color:    color.Black,
		Width:    vg.Points(1),
		Dashes:   []vg.Length{},
		DashOffs: 0,
	}
)

// MySinglePlot is a wrapper of Line of package plotter with slice of float64 x.
func MySinglePlot(x []float64) {
	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = float64(i)
		data[i].Y = x[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	l, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}

	p.Add(l)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

// MyPlot is a wrapper of Line of package plotter with slice of float64 x and y.
func MyPlot(x, y []float64) {
	if len(x) != len(y) {
		log.Fatal("length of x and y have to same.")
	}

	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = x[i]
		data[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	l, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}

	p.Add(l)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

// MySingleScatter is a wrapper of Scatter of package plotter with slice of float64 x.
func MySingleScatter(x []float64) {
	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = float64(i)
		data[i].Y = x[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}

	s.Radius = vg.Length(1)

	p.Add(s)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}

}

// MyScatter is a wrapper of Scatter of package plotter with slice of float64 x and y.
func MyScatter(x, y []float64) {
	if len(x) != len(y) {
		log.Fatal("length of x and y have to same.")
	}

	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = x[i]
		data[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}

	s.Radius = vg.Length(2)

	p.Add(s)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

func MyPlotWithScatter(x, y []float64) {
	if len(x) != len(y) {
		log.Fatal("length of x and y have to same.")
	}

	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = x[i]
		data[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}

	s.Radius = vg.Length(2)
	p.Add(s)

	l, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}

	p.Add(l)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

// func MyCanclePlot(x [][4]float64) {
func MyCandleChart() {
	c := candle{
		start: 2,
		end:   3,
		min:   1,
		max:   4,
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// rand.Seed(int64(0))
	// n := 10
	// uniform := make(plotter.Values, n)
	// for i := 0; i < n; i++ {
	// 	uniform[i] = rand.Float64()
	// }

	w := vg.Points(20)
	// b, err := plotter.NewBoxPlot(w, 0, uniform)
	// if err != nil {
	// 	panic(err)
	// }
	// tmp := []float64{3, 4, 5, 6}

	cc, err := NewCandleChart(w, 0, c)
	if err != nil {
		panic(err)
	}
	// fmt.Println(cc)
	// b.Values = tmp
	// b.Median = 3

	// b.AdjLow = math.Inf(1)
	// fmt.Println(b.Location)

	// low := b.Quartile1 - 1.5*(b.Quartile3-b.Quartile1)
	// high := b.Quartile3 + 1.5*(b.Quartile3-b.Quartile1)
	// for i, v := range b.Values {
	// 	if v > high || v < low {
	// 		b.Outside = append(b.Outside, i)
	// 		continue
	// 	}
	// 	if v < b.AdjLow {
	// 		b.AdjLow = v
	// 	}
	// 	if v > b.AdjHigh {
	// 		b.AdjHigh = v
	// 	}
	// }

	// fmt.Println(b)

	p.Add(cc)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}

}
