package matrix

import (
	"errors"
)

var ErrNotMatrix4 = errors.New("Given matrix is not a 4*4 matrix")

//Matrix4 defines 4*4 matrix
type Matrix4 [4][4]float64

//make array to Matrix4
func MakeMatrix4(value [16]float64) Matrix4 {
	var result Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = value[i*4+j]
		}
	}
	return result
}

func ZeroMatrix4() Matrix4 {
	return MakeMatrix4([16]float64{})
}
func UnitMatrix4() Matrix4 {
	return MakeMatrix4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
}

func (mat Matrix4) Equal(mat2 Matrix4) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if mat[i][j] != mat2[i][j] {
				return false
			}
		}
	}

	return true
}

//Return Scaled Matrix
func (mat Matrix4) Scaled(C float64) (result Matrix4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = mat[i][j] * C
		}
	}
	return
}

//Add another 4*4 matrix
func (mat Matrix4) Add(mat2 Matrix4) (result Matrix4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = mat[i][j] + mat2[i][j]
		}
	}

	return
}

//Subtract another 4*4 matrix
func (mat Matrix4) Sub(mat2 Matrix4) (result Matrix4) {
	result = mat.Add(mat2.Scaled(-1))
	return
}

//Multiply another 4*4 matrix
func (mat Matrix4) Mul(mat2 Matrix4) (result Matrix4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s := 0.0
			for k := 0; k < 4; k++ {
				s += mat[i][k] * mat2[k][j]
			}
			result[i][j] = s
		}
	}

	return result
}

//Transpose of the matrix
func (mat Matrix4) Trans() Matrix4 {
	var result Matrix4

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = mat[j][i]
		}
	}

	return result
}

func (mat Matrix4) minor(i, j int) Matrix3 {
	cnt := 0
	arr := [9]float64{}
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if r != i && c != j {
				arr[cnt] = mat[r][c]
				cnt++
			}
		}
	}

	return MakeMatrix3(arr)
}

//Determinant of the matrix
func (mat Matrix4) Det() float64 {
	var CA = mat.minor(0, 0)
	var CB = mat.minor(0, 1)
	var CC = mat.minor(0, 2)
	var CD = mat.minor(0, 3)

	A, B, C, D := mat[0][0], mat[0][1], mat[0][2], mat[0][3]
	return A*CA.Det() - B*CB.Det() + C*CC.Det() - D*CD.Det()
}

//Inverse matrix of the matrix
func (mat Matrix4) Inv() Matrix4 {
	var result Matrix4

	if mat.Det() == 0 {
		panic(ErrNoInverseMatrix)
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = mat.minor(i, j).Det()
			if (i+j)%2 == 1 {
				result[i][j] *= -1
			}
		}
	}

	result = result.Trans()

	return result.Scaled(1.0 / mat.Det())
}
