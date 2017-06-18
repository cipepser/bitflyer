package myUtil

import (
	"errors"
	"image/color"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
)

type candle struct {
	start, end, max, min float64
}

type fourStatPlot struct {
	// Location is the location of the box along its axis.
	Location float64

	// Color is the fill color of the candle.
	Color color.Color

	// Start and End are the first and last value of the period respectively.
	Start, End float64

	// Low and High are the lowest and highest value of the period respectively.
	Low, High float64
}

type CandleChart struct {
	fourStatPlot

	// Offset is added to the x location of each candle.
	// When the Offset is zero, the candles are drawn
	// centered at their x location.
	Offset vg.Length

	// Width is the width used to draw the candle.
	Width vg.Length

	// CapWidth is the width of the cap used to top
	// off a whisker.
	CapWidth vg.Length

	// GlyphStyle is the style of the outside point glyphs.
	GlyphStyle draw.GlyphStyle

	// CandleStyle is the line style for the candle.
	CandleStyle draw.LineStyle

	// MedianStyle is the line style for the median line.
	MedianStyle draw.LineStyle

	// WhiskerStyle is the line style used to draw the
	// whiskers.
	WhiskerStyle draw.LineStyle

	// Min and Max are the extreme values of the data.
	Min, Max float64
}

func NewCandleChart(w vg.Length, loc float64, c candle) (*CandleChart, error) {
	if w < 0 {
		return nil, errors.New("Negative candlechart width")
	}

	cc := new(CandleChart)
	var err error
	if cc.fourStatPlot, err = newfourStatPlot(loc, c); err != nil {
		return nil, err
	}

	cc.Width = w
	cc.Min = c.min * 0.8
	cc.Max = c.max * 1.2

	cc.GlyphStyle = plotter.DefaultGlyphStyle
	cc.CandleStyle = NegativeLineStyle
	cc.WhiskerStyle = draw.LineStyle{
		Width: vg.Points(1),
	}
	return cc, nil
}

func newfourStatPlot(loc float64, c candle) (fourStatPlot, error) {
	var cc fourStatPlot
	cc.Location = loc

	if c.start < c.end {
		cc.Start = c.start
		cc.End = c.end
	} else {
		cc.Start = c.end
		cc.End = c.start
	}

	cc.Color = color.Black

	cc.Low = c.min
	cc.High = c.max

	return cc, nil
}

func (cc *CandleChart) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)
	x := trX(cc.Location)
	if !c.ContainsX(x) {
		return
	}
	x += cc.Offset

	q1 := trY(cc.Start)
	q3 := trY(cc.End)
	Low := trY(cc.Low)
	High := trY(cc.High)

	pts := []vg.Point{
		{x - cc.Width/2, q1},
		{x - cc.Width/2, q3},
		{x + cc.Width/2, q3},
		{x + cc.Width/2, q1},
		{x - cc.Width/2 - cc.CandleStyle.Width/2, q1},
	}

	poly := c.ClipPolygonY(pts)
	c.FillPolygon(cc.Color, poly)

	box := c.ClipLinesY(pts)
	c.StrokeLines(cc.CandleStyle, box...)

	whisks := c.ClipLinesY([]vg.Point{{x, q3}, {x, High}},
		[]vg.Point{{x, High}, {x, High}},
		[]vg.Point{{x, q1}, {x, Low}},
		[]vg.Point{{x, Low}, {x, Low}})
	c.StrokeLines(cc.WhiskerStyle, whisks...)

}

// DataRange returns the minimum and maximum x
// and y values, implementing the plot.DataRanger
// interface.
func (cc *CandleChart) DataRange() (float64, float64, float64, float64) {
	return cc.Location, cc.Location, cc.Min, cc.Max
}
