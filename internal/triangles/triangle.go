package triangles

import (
	"github.com/gotk3/gotk3/cairo"
)

type side int

const (
	sideP1P2 side = iota
	sideP2P3
	sideP3P1
)

type point struct {
	x, y float64
}
type triangle struct {
	p1, p2, p3 point
}

func createInitialTriangle(h float64, w float64) {
	t := triangle{
		p1: point{100, h - 100},
		p2: point{w - 100, h - 100},
		p3: point{w / 2, 100},
	}

	triangles = append(triangles, t)
}

func (t triangle) subDivide() (t1, t2, t3, t4 triangle) {
	// Get midpoints
	m1 := t.getMidPoint(sideP1P2)
	m2 := t.getMidPoint(sideP2P3)
	m3 := t.getMidPoint(sideP3P1)

	// Create 4 new triangles
	t1 = triangle{t.p1, m3, m1}
	t2 = triangle{m1, t.p2, m2}
	t3 = triangle{m2, t.p3, m3}
	t4 = triangle{m1, m2, m3}

	return
}

func (t triangle) getMidPoint(s side) point {
	var p1, p2 point

	// Get the correct points
	switch s {
	case sideP1P2:
		p1 = t.p1
		p2 = t.p2
	case sideP2P3:
		p1 = t.p2
		p2 = t.p3
	case sideP3P1:
		p1 = t.p3
		p2 = t.p1
	}

	// Get the x and y for the points, where x1<x2 and y1<y2
	x1, x2, y1, y2 := p1.x, p2.x, p1.y, p2.y
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	// Return midpoint
	return point{x1 + (x2-x1)/2, y1 + (y2-y1)/2}
}

func (t triangle) draw(ctx *cairo.Context) {
	ctx.SetSourceRGB(0, 0, 0)
	ctx.MoveTo(t.p1.x, t.p1.y)
	ctx.LineTo(t.p2.x, t.p2.y)
	ctx.LineTo(t.p3.x, t.p3.y)
	ctx.LineTo(t.p1.x, t.p1.y)
	ctx.Stroke()
}
