package gogamedevutils

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	pos := Vec2{
		x: 100,
		y: 100,
	}
	speed := Vec2{
		x: 2,
		y: 5,
	}

	expected := Vec2{
		x: 102,
		y: 105,
	}

	pos.Add(speed)
	if pos.x != expected.x || pos.y != expected.y {
		t.Errorf("Add failed, expected %v, got %v", expected, pos)
	}
}

func TestSub(t *testing.T) {
	pos := Vec2{
		x: 100,
		y: 100,
	}
	speed := Vec2{
		x: 2,
		y: 5,
	}

	expected := Vec2{
		x: 98,
		y: 95,
	}

	pos.Sub(speed)
	if pos.x != expected.x || pos.y != expected.y {
		t.Errorf("Add failed, expected %v, got %v", expected, pos)
	}
}

func TestMul(t *testing.T) {
	u := Vec2{
		x: 10,
		y: 10,
	}
	scale := 0.5

	expected := Vec2{
		x: 5,
		y: 5,
	}

	u.Mul(scale)

	if u.x != expected.x || u.y != expected.y {
		t.Errorf("Add failed, expected %v, got %v", expected, u)
	}
}

func TestDiv(t *testing.T) {
	u := Vec2{
		x: 10,
		y: 10,
	}
	div := 2.0

	expected := Vec2{
		x: 5,
		y: 5,
	}

	u.Div(div)

	if u.x != expected.x || u.y != expected.y {
		t.Errorf("Add failed, expected %v, got %v", expected, u)
	}
}

func TestMag(t *testing.T) {
	u := Vec2{
		x: 10,
		y: 10,
	}

	hyp := 14.14
	// beware of floating points
	m := math.Floor(u.Mag()*100) / 100
	if m != hyp {
		t.Errorf("Add failed, expected %v, got %v", hyp, m)
	}
}

func TestNorm(t *testing.T) {
	m := Vec2{
		x: 10,
		y: 10,
	}
	expect := Vec2{
		x: 0.70,
		y: 0.70,
	}
	m.Normalize()
	// beware of floating points
	res := Vec2{
		x: math.Floor(m.x*100) / 100,
		y: math.Floor(m.y*100) / 100,
	}
	if res.x != expect.x || res.y != expect.y {
		t.Errorf("Add failed, expected %v, got %v", expect, res)
	}
}
