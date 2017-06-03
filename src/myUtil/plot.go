package myUtil

import (
	"log"
	"os/exec"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func MyPlot(x, y []float64) {
	if len(x) != len(y) {
		log.Fatal("length of x and y have to same.")
	}

	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = x[i]
		data[i].Y = y[i] * y[i]
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

	file := "line.png"
	if err = p.Save(5*vg.Inch, 3*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}
