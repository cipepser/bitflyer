package myUtil

import (
	"image/color"
	"math"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
)

// Candle represents a candle given data which give the start,
// end, low and high value.
type Candle struct {
	// X is used by taking the label of x to right space.
	X float64

	// start and end are the first and last value of the period respectively.
	// low and high are the lowest and highest value of the period respectively.
	start, end, low, high float64

	// Color is the fill color of the candle, which represents positive line and negative line.
	Color color.Color
}

// NewCandle construct an object for type Candle,
// with the given x and data.
func NewCandle(x float64, data []float64) (*Candle, error) {
	c := new(Candle)

	c.X = x

	c.start = data[0]
	c.end = data[len(data)-1]

	if c.start > c.end {
		c.Color = color.Black
	} else {
		c.Color = color.White
	}

	c.low = data[0]
	c.high = data[0]

	for _, d := range data {
		c.low = math.Min(c.low, d)
		c.high = math.Max(c.high, d)
	}

	return c, nil
}

// Candles is a slice of Candle.
type Candles struct {
	candles []Candle
}

// NewCandles construct an object for type Candles,
// with the data.
func NewCandles(data [][]float64) ([]Candle, error) {
	cs := make([]Candle, len(data))
	for i, d := range data {
		c, err := NewCandle(float64(i), d)
		if err != nil {
			return nil, err
		}
		cs[i] = *c
	}

	return cs, nil
}

// getMax returns the largest value of candles.
func getMax(cs []Candle) float64 {
	max := cs[0].high

	for _, c := range cs {
		if max < c.high {
			max = c.high
		}
	}

	return max
}

// getMax returns the smallest value of candles.
func getMin(cs []Candle) float64 {
	min := cs[0].low

	for _, c := range cs {
		if min > c.low {
			min = c.low
		}
	}

	return min
}

// CandleChart implements the Plotter interface, drawing
// a candle chart of candles.
type CandleChart struct {
	Candles

	// GlyphStyle is the style of the outside point glyphs.
	GlyphStyle draw.GlyphStyle

	// CandleStyle is the line style for the candle.
	CandleStyle draw.LineStyle

	// WhiskerStyle is the line style used to draw the whiskers.
	WhiskerStyle draw.LineStyle

	// Min and Max are the canvas size for Y-axis.
	Min, Max float64
}

// CandleChart creates as new candle chart plotter for
// the given data.
func NewCandleChart(data [][]float64) (*CandleChart, error) {
	cc := new(CandleChart)

	var err error
	cc.candles, err = NewCandles(data)
	if err != nil {
		return nil, err
	}

	cc.Min = getMin(cc.candles) * 0.9
	cc.Max = getMax(cc.candles) * 1.1

	cc.GlyphStyle = plotter.DefaultGlyphStyle
	cc.CandleStyle = NegativeLineStyle
	cc.WhiskerStyle = draw.LineStyle{
		Width: vg.Points(1),
	}

	return cc, nil
}

// Plot implements the Plot method of the plot.Plotter interface.
func (cc *CandleChart) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)

	var w vg.Length
	if len(cc.candles) < 2 {
		return
	} else {
		w = trX(cc.candles[1].X) - trX(cc.candles[0].X)
	}

	for _, candle := range cc.candles {
		x := trX(candle.X)

		l := trY(candle.low)
		h := trY(candle.high)
		var q1, q3 vg.Length

		if candle.start < candle.end {
			q1 = trY(candle.start)
			q3 = trY(candle.end)
		} else {
			q1 = trY(candle.end)
			q3 = trY(candle.start)
		}

		pts := []vg.Point{
			{x - w/2, q1},
			{x - w/2, q3},
			{x + w/2, q3},
			{x + w/2, q1},
			{x - w/2 - cc.CandleStyle.Width/2, q1},
		}

		poly := c.ClipPolygonY(pts)
		c.FillPolygon(candle.Color, poly)

		box := c.ClipLinesY(pts)
		c.StrokeLines(cc.CandleStyle, box...)

		whisks := c.ClipLinesY([]vg.Point{{x, q3}, {x, h}},
			[]vg.Point{{x, h}, {x, h}},
			[]vg.Point{{x, q1}, {x, l}},
			[]vg.Point{{x, l}, {x, l}})
		c.StrokeLines(cc.WhiskerStyle, whisks...)
	}

}

// DataRange implements the DataRange method
// of the plot.DataRanger interface.
func (cc *CandleChart) DataRange() (xmin, xmax, ymin, ymax float64) {
	return 0, float64(len(cc.candles)) * 1.3, cc.Min, cc.Max
}
