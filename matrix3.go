package matrix

import (
	"errors"
)

var ErrNotMatrix3 = errors.New("Given matrix is not a 3*3 matrix")

//Matrix3 defines 3*3 matrix
type Matrix3 [3][3]float64

//make array to Matrix3
func MakeMatrix3(value [9]float64) Matrix3 {
	var result Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = value[i*3+j]
		}
	}
	return result
}

func ZeroMatrix3() Matrix3 {
	return MakeMatrix3([9]float64{})
}
func UnitMatrix3() Matrix3 {
	return MakeMatrix3([9]float64{1, 0, 0, 0, 1, 0, 0, 0, 1})
}

func (mat Matrix3) Equal(mat2 Matrix3) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !Equal(mat[i][j], mat2[i][j]) {
				return false
			}
		}
	}

	return true
}

//Return Scaled Matrix
func (mat Matrix3) Scaled(C float64) (result Matrix3) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = mat[i][j] * C
		}
	}
	return
}

//Add another 3*3 matrix
func (mat Matrix3) Add(mat2 Matrix3) (result Matrix3) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = mat[i][j] + mat2[i][j]
		}
	}

	return
}

//Subtract another 3*3 matrix
func (mat Matrix3) Sub(mat2 Matrix3) (result Matrix3) {
	result = mat.Add(mat2.Scaled(-1))
	return
}

//Multiply another 3*3 matrix
func (mat Matrix3) Mul(mat2 Matrix3) (result Matrix3) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s := 0.0
			for k := 0; k < 3; k++ {
				s += mat[i][k] * mat2[k][j]
			}
			result[i][j] = s
		}
	}

	return result
}

//Transpose of the matrix
func (mat Matrix3) Trans() Matrix3 {
	var result Matrix3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = mat[j][i]
		}
	}

	return result
}

func (mat Matrix3) minor(i, j int) Matrix2 {
	cnt := 0
	arr := [4]float64{}
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if r != i && c != j {
				arr[cnt] = mat[r][c]
				cnt++
			}
		}
	}

	return MakeMatrix2(arr)
}

//Determinant of the matrix
func (mat Matrix3) Det() float64 {
	var CA = mat.minor(0, 0)
	var CB = mat.minor(0, 1)
	var CC = mat.minor(0, 2)

	A, B, C := mat[0][0], mat[0][1], mat[0][2]
	return A*CA.Det() - B*CB.Det() + C*CC.Det()
}

//Inverse matrix of the matrix
func (mat Matrix3) Inv() Matrix3 {
	var result Matrix3

	if mat.Det() == 0 {
		panic(ErrNoInverseMatrix)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = mat.minor(i, j).Det()
			if (i+j)%2 == 1 {
				result[i][j] *= -1
			}
		}
	}

	result = result.Trans()

	return result.Scaled(1.0 / mat.Det())
}
