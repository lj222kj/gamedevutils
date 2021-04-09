package gogamedevutils

import (
	"math"
)

type Vec2 struct {
	x, y float64
}

func (v *Vec2) Add(v2 Vec2) {
	v.y = v.y + v2.y
	v.x = v.x + v2.x
}

func (v *Vec2) Sub(v2 Vec2) {
	v.y = v.y - v2.y
	v.x = v.x - v2.x
}

func (v *Vec2) Mul(n float64) {
	v.y = v.y * n
	v.x = v.x * n
}

func (v *Vec2) Div(n float64) {
	v.y = v.y / n
	v.x = v.x / n
}

func (v *Vec2) Mag() float64 {
	return math.Sqrt(v.x*v.x + v.x*v.y)
}

func (v *Vec2) Normalize() {
	m := v.Mag()
	if m > 0 {
		v.Div(m)
	}
}
