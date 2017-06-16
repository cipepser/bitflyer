package myUtil

import (
	"log"
	"os/exec"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
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
