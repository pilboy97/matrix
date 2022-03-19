package matrix_test

import (
	"math"
	"matrix"
	"testing"
)

func TestMoved2(t *testing.T) {
	v := matrix.MakeADot2(0, 0)
	mat := matrix.Moved2(1, 3)

	res := mat.Proj(v)

	if !res.Equal(matrix.MakeADot2(1, 3)) {
		t.Fatal(mat, matrix.MakeADot2(1, 3), res)
	}
}
func TestRotate2(t *testing.T) {
	v := matrix.MakeADot2(1, 0)
	mat := matrix.Rotated2(math.Pi / 2)

	res := mat.Proj(v)

	if !res.Equal(matrix.MakeADot2(0, 1)) {
		t.Fatal(mat, matrix.MakeADot2(0, 1), res)
	}
}

func TestScaled2(t *testing.T) {
	v := matrix.MakeADot2(1, 1)
	mat := matrix.Scaled2(3, 5)

	res := mat.Proj(v)

	if !res.Equal(matrix.MakeADot2(3, 5)) {
		t.Fatal(mat, matrix.MakeADot2(3, 5), res)
	}
}
