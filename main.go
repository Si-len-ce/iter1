package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const ITER int = 50
const RADIUS float64 = 75
const PI4 float64 = 4 / math.Pi

func add(thing []pixel.Vec, element pixel.Vec) []pixel.Vec {
	result := make([]pixel.Vec, len(thing)+1)
	result[0] = element
	for i := 1; i < len(thing)+1; i++ {
		result[i] = thing[i-1]
	}
	return result
}

func nthFib(n int) float64 {
	pn := 1.
	nn := 1.
	for i := 1; i < n; i++ {
		nn = pn + nn
		pn = nn
	}
	return nn
}

func Run() {
	t := 0.
	cfg := pixelgl.WindowConfig{
		Title:  "fourier",
		Bounds: pixel.R(-600, -200, 600, 200),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	running := true
	i := 1
	center := pixel.V(-350, 0)
	pos := center
	imd := imdraw.New(win)
	radius := RADIUS * PI4
	prevpos := pos
	dn := 0.
	points := make([]pixel.Vec, 1)
	fn := 0.
	signal := [24]float64{100, 100, 100, 100, -100, -100, -100, -100, 100, 100, 100, 100, -100, -100, -100, -100, 100, 100, 100, 100, -100, -100, -100, -100}
	fourierY = 
	for running {
		win.Clear(colornames.Black)
		imd.Clear()
		imd.Reset()
		imd.Push(center)
		imd.Circle(RADIUS*PI4, 1)
		for ; i <= ITER; i++ {
			dn = nthFib(i)
			fn = dn
			radius = PI4 * (RADIUS / float64(dn))
			imd.Color = colornames.Gray

			prevpos = pos
			pos = prevpos.Add(pixel.V(radius*math.Cos(fn*t), radius*math.Sin(fn*t)))
			points[0] = pos

			imd.Push(prevpos, pos)
			imd.Line(1.)
			imd.Push(prevpos)
			imd.Circle(radius, 1.)
		}
		points = add(points, pos)
		if len(points) > 700 {
			points = points[:700]
		}
		for i = range points[:len(points)-1] {
			imd.Push(points[i], points[i+1])
			imd.Line(1.)
			imd.Push(pixel.V(float64(i+1), points[i+1].Y), pixel.V(float64(i), points[i].Y))
			imd.Line(1.)
		}
		imd.Push(pos, pixel.V(0, pos.Y))
		imd.Line(1.)
		imd.Draw(win)
		pos = center
		prevpos = pos
		i = 1
		t += 0.01
		win.Update()
		if win.Closed() {
			running = false
		}
	}
}

func main() {
	pixelgl.Run(Run)
}
