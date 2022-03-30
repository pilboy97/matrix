package matrix_test

import (
	"math"
	"matrix"
	"testing"
)

func TestMove2(t *testing.T) {
	V := matrix.MakeDot2(1, 3)
	D := matrix.MakeVec2(1, 3)

	if !V.Moved2(-1, -3).Equal(matrix.MakeDot2(0, 0)) {
		t.Log(V.Moved2(-1, -3))
		t.Fail()
	}
	if !D.Moved2(-1, -3).Equal(D) {
		t.Log(D.Moved2(-1, -3))
		t.Fail()
	}

	V = matrix.MakeDot3(1, 3, 5)
	D = matrix.MakeVec3(1, 3, 5)

	if !V.Moved3(-1, -3, -5).Equal(matrix.MakeDot3(0, 0, 0)) {
		t.Log(V.Moved3(-1, -3, -5))
		t.Fail()
	}
	if !D.Moved3(-1, -3, -5).Equal(D) {
		t.Log(D.Moved3(-1, -3, -5))
		t.Fail()
	}
}

func TestScale(t *testing.T) {
	V := matrix.MakeDot2(1, 3)
	D := matrix.MakeVec2(1, 3)

	if !V.Scaled2(1, 3).Equal(matrix.MakeDot2(1, 9)) {
		t.Log(V.Scaled2(1, 3))
		t.Fail()
	}
	if !D.Scaled2(1, 3).Equal(matrix.MakeVec2(1, 9)) {
		t.Log(D.Scaled2(1, 3))
		t.Fail()
	}

	V = matrix.MakeDot3(1, 3, 5)
	D = matrix.MakeVec3(1, 3, 5)

	if !V.Scaled3(1, 3, 5).Equal(matrix.MakeDot3(1, 9, 25)) {
		t.Log(V.Scaled3(1, 3, 5))
		t.Fail()
	}
	if !D.Scaled3(1, 3, 5).Equal(matrix.MakeVec3(1, 9, 25)) {
		t.Log(D.Scaled3(1, 3, 5))
		t.Fail()
	}
}
func TestRotate(t *testing.T) {
	D := matrix.MakeDot2(1, 0)
	V := matrix.MakeVec2(1, 0)

	if !D.Rotated2(math.Pi / 2).Equal(matrix.MakeDot2(0, 1)) {
		t.Log(D.Rotated2(math.Pi / 2))
		t.Fail()
	}
	if !V.Rotated2(math.Pi / 2).Equal(matrix.MakeVec2(0, 1)) {
		t.Log(V.Rotated2(math.Pi / 2))
		t.Fail()
	}

	D = matrix.MakeDot3(1, 0, 0)
	V = matrix.MakeVec3(1, 0, 0)

	if !D.RotatedZ3(math.Pi / 2).Equal(matrix.MakeDot3(0, 1, 0)) {
		t.Log(D.RotatedZ3(math.Pi / 2))
		t.Fail()
	}
	if !V.RotatedZ3(math.Pi / 2).Equal(matrix.MakeVec3(0, 1, 0)) {
		t.Log(V.RotatedZ3(math.Pi / 2))
		t.Fail()
	}

	D = matrix.MakeDot3(0, 0, 1)
	V = matrix.MakeVec3(0, 0, 1)

	if !D.RotatedY3(math.Pi / 2).Equal(matrix.MakeDot3(1, 0, 0)) {
		t.Log(D.RotatedY3(math.Pi / 2))
		t.Fail()
	}
	if !V.RotatedY3(math.Pi / 2).Equal(matrix.MakeVec3(1, 0, 0)) {
		t.Log(D.RotatedY3(math.Pi / 2))
		t.Fail()
	}

	D = matrix.MakeDot3(0, 1, 0)
	V = matrix.MakeVec3(0, 1, 0)

	if !D.RotatedX3(math.Pi / 2).Equal(matrix.MakeDot3(0, 0, 1)) {
		t.Log(D.RotatedX3(math.Pi / 2))
		t.Fail()
	}
	if !V.RotatedX3(math.Pi / 2).Equal(matrix.MakeVec3(0, 0, 1)) {
		t.Log(D.RotatedX3(math.Pi / 2))
		t.Fail()
	}
}
