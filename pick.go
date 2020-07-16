package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/cstockton/go-conv"
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"strconv"
)

var container *fyne.Container
var iterations int
var height int
var width int
var iterationstring *widget.Entry
var lengthstring *widget.Entry
var saveimage bool
var dc *gg.Context

func main() {
	height = 900
	width = 1600
	dc = gg.NewContext(width, height)
	application := app.New()
	window := application.NewWindow("ToothPick Fractal")
	container = fyne.NewContainer()
	iterationstring = widget.NewEntry()
	slide := widget.NewSlider(-math.Pi, math.Pi)
	//slide.OnChanged = renderstart
	slide.Step = 0.001
	iterationstring.PlaceHolder = "Number Of Iterations (for best results put a power of 2)"
	renderbtton := widget.NewButton("Render", func() { renderstart(slide.Value); saveimage = true })
	lengthstring = widget.NewEntry()
	lengthstring.PlaceHolder = "Length of Each 'ToothPick'"
	window.SetContent(widget.NewVBox(container, renderbtton, iterationstring, lengthstring, slide))
	window.ShowAndRun()
}

func renderstart(initangle float64) {
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	container.Resize(fyne.NewSize(width, height))
	iterations, _ = strconv.Atoi(iterationstring.Text)
	length, _ := conv.Float64(lengthstring.Text)
	centerx := float64(width) / 2
	centery := float64(height) / 2
	x1, y1, x2, y2 := drawline(centerx, centery, 0, 255, 255, 255, 255, length)
	recurse(x1, y1, math.Pi/2, iterations, length)
	recurse(x2, y2, math.Pi/2, iterations, length)
}

func recurse(x, y, angle float64, iter int, lent float64) {
	linecolor := new(color.RGBA)
	linecolor.R, _ = conv.Uint8(map1(x, 0, float64(width), 100, 255))
	linecolor.G, _ = conv.Uint8(map1(y, 0, float64(height), 100, 255))
	linecolor.B, _ = conv.Uint8(map1(float64(iter), 0, float64(iterations), 100, 255))
	r := linecolor.R
	g := linecolor.G
	b := linecolor.B
	x1, y1, x2, y2 := drawline(x, y, angle, r, g, b, 170, lent)
	if lent > 3 {
		recurse(x1, y1, angle+math.Pi/2, iter-1, lent/math.Sqrt2)
		recurse(x2, y2, angle+math.Pi/2, iter-1, lent/math.Sqrt2)
	}
}

func drawline(x, y, angle float64, r uint8, g uint8, b uint8, a uint8, lent float64) (float64, float64, float64, float64) {
	x1 := x + lent*math.Cos(angle)
	y1 := y + lent*math.Sin(angle)
	x2 := x - lent*math.Cos(angle)
	y2 := y - lent*math.Sin(angle)
	dc.DrawLine(x1, y1, x2, y2)
	dc.SetRGBA255(int(r), int(g), int(b), 255)
	dc.SetLineWidth(1)
	dc.Stroke()
	frameno++
	if frameno > 22266 {
		dc.SavePNG(strconv.Itoa(frameno) + ".png")
	}

	return x1, y1, x2, y2
}

var frameno int

func map1(value float64, istart float64, istop float64, ostart float64, ostop float64) float64 {
	return ostart + (ostop-ostart)*((value-istart)/(istop-istart))
}
