package matrix_test

import (
	"math"
	"matrix"
	"testing"
)

func TestMoved3(t *testing.T) {
	v := matrix.MakeADot3(0, 0, 0)
	mat := matrix.Moved3(1, 3, 5)

	res := mat.Proj(v)

	if !res.Equal(matrix.MakeADot3(1, 3, 5)) {
		t.Fatal(mat, matrix.MakeADot3(1, 3, 5), res)
	}
}
func TestRotateX3(t *testing.T) {
	v := matrix.MakeADot3(0, 1, 0)
	mat := matrix.RotatedX3(math.Pi / 2)

	res := mat.Proj(v)

	if !res.Equal(matrix.MakeADot3(0, 0, 1)) {
		t.Fatal(mat, matrix.MakeADot3(0, 0, 1), res)
	}
}
func TestRotateY3(t *testing.T) {
	v := matrix.MakeADot3(1, 0, 0)
	mat := matrix.RotatedY3(math.Pi / 2)

	res := mat.Proj(v)

	if !res.Equal(matrix.MakeADot3(0, 0, 1)) {
		t.Fatal(mat, matrix.MakeADot3(0, 0, 1), res)
	}
}
func TestRotateZ3(t *testing.T) {
	v := matrix.MakeADot3(1, 0, 0)
	mat := matrix.RotatedZ3(math.Pi / 2)

	res := mat.Proj(v)

	if !res.Equal(matrix.MakeADot3(0, 1, 0)) {
		t.Fatal(mat, matrix.MakeADot3(0, 1, 0), res)
	}
}

func TestScaled3(t *testing.T) {
	v := matrix.MakeADot3(1, 1, 1)
	mat := matrix.Scaled3(3, 5, 7)

	res := mat.Proj(v)

	if !res.Equal(matrix.MakeADot3(3, 5, 7)) {
		t.Fatal(mat, matrix.MakeADot3(3, 5, 7), res)
	}
}
