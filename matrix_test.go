package matrix_test

import (
	"matrix"
	"testing"
)

func TestAdd(t *testing.T) {
	X := matrix.MakeMat(2, 3, 1, 3, 2, 4, 6, 2)
	Y := matrix.MakeMat(2, 3, 6, 3, 5, 2, 3, 4)
	Z := matrix.MakeMat(2, 3, 7, 6, 7, 6, 9, 6)

	if !Z.Equal(X.Add(Y)) {
		t.Log(X.Add(Y))
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	X := matrix.MakeMat(1, 2, 1, 3)
	Y := matrix.MakeMat(1, 2, 3, 2)
	Z := matrix.MakeMat(1, 2, -2, 1)

	if !Z.Equal(X.Sub(Y)) {
		t.Log(X.Sub(Y))
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	X := matrix.MakeMat(1, 2, 1, 3)
	Y := matrix.MakeMat(2, 2, 3, 2, 4, 5)
	Z := matrix.MakeMat(1, 2, 15, 17)

	if !Z.Equal(X.Mul(Y)) {
		t.Log(X.Mul(Y))
		t.Fail()
	}
}

func TestDet(t *testing.T) {
	X := matrix.MakeMat(3, 3, 1, 2, 4, -1, 3, 0, 4, 1, 0)
	Z := -52.0

	if !matrix.ASame(X.Det(), Z) {
		t.Log(matrix.ASame(X.Det(), Z))
		t.Fail()
	}
}

func TestInv(t *testing.T) {
	X := matrix.MakeMat(3, 3, 1, 2, 0, 3, 2, 1, 2, 1, 0)
	Z := matrix.MakeMat(3, 3, -1, 0, 2, 2, 0, -1, -1, 3, -4).Scaled(1.0 / 3)

	if !Z.Equal(X.Inv()) {
		t.Log(X.Inv())
		t.Fail()
	}

	X = matrix.MakeMat(2, 2, 1, 2, 3, 4)
	Z = matrix.MakeMat(2, 2, -4, 2, 3, -1).Scaled(1.0 / 2)

	if !Z.Equal(X.Inv()) {
		t.Log(X.Inv())
		t.Fail()
	}
}

func TestMatrix(t *testing.T) {
	X := matrix.MakeMat(2, 3, 1, 2, 3, 4, 5, 6)

	if !X.Row(1).Equal(matrix.MakeMat(1, 3, 4, 5, 6)) {
		t.Log(X.Row(1))
		t.Fail()
	}

	if !X.Col(1).Equal(matrix.MakeMat(2, 1, 2, 5)) {
		t.Log(X.Col(1))
		t.Fail()
	}
}
func TestJoin(t *testing.T) {
	X := matrix.MakeMat(1, 3, 1, 2, 3)
	Y := matrix.MakeMat(1, 3, 4, 5, 6)
	Z := matrix.MakeMat(2, 3, 1, 2, 3, 4, 5, 6)

	if !matrix.JoinRow(X, Y).Equal(Z) {
		t.Log(matrix.JoinRow(X, Y))
		t.Fail()
	}

	X = matrix.MakeMat(3, 1, 1, 2, 3)
	Y = matrix.MakeMat(3, 1, 4, 5, 6)
	Z = matrix.MakeMat(3, 2, 1, 4, 2, 5, 3, 6)

	if !matrix.JoinCol(X, Y).Equal(Z) {
		t.Log(matrix.JoinCol(X, Y))
		t.Fail()
	}
}
