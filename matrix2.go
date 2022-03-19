package matrix

import "errors"

var (
	ErrNotMatrix2      = errors.New("This matrix is not a 2*2 matrix")
	ErrNoInverseMatrix = errors.New("This matrix has no inverse matrix")
)

//Matrix2 defines 2*2 matrix
type Matrix2 [2][2]float64

//make array to matrix2
func MakeMatrix2(value [4]float64) Matrix2 {
	a, b, c, d := value[0], value[1], value[2], value[3]
	return Matrix2{{a, b}, {c, d}}
}

func ZeroMatrix2() Matrix2 {
	return MakeMatrix2([4]float64{0, 0, 0, 0})
}
func UnitMatrix2() Matrix2 {
	return MakeMatrix2([4]float64{1, 0, 1, 0})
}

func (mat Matrix2) Equal(mat2 Matrix2) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if !Equal(mat[i][j], mat2[i][j]) {
				return false
			}
		}
	}

	return true
}

//Return Scaled Matrix
func (mat Matrix2) Scaled(C float64) (result Matrix2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result[i][j] = mat[i][j] * C
		}
	}
	return
}

//Add another 2*2 matrix
func (mat Matrix2) Add(mat2 Matrix2) (result Matrix2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result[i][j] = mat[i][j] + mat2[i][j]
		}
	}

	return
}

//Subtract another 2*2 matrix
func (mat Matrix2) Sub(mat2 Matrix2) (result Matrix2) {
	result = mat.Add(mat2.Scaled(-1))
	return
}

//Multiply another 2*2 matrix
func (mat Matrix2) Mul(mat2 Matrix2) (result Matrix2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			s := 0.0
			for k := 0; k < 2; k++ {
				s += mat[i][k] * mat2[k][j]
			}
			result[i][j] = s
		}
	}

	return result
}

//Transpose of the matrix
func (mat Matrix2) Trans() Matrix2 {
	return Matrix2{{mat[0][0], mat[1][1]}, {mat[1][0], mat[0][1]}}
}

//Determinant of the matrix
func (mat Matrix2) Det() float64 {
	return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
}

//Inverse matrix of the matrix
func (mat Matrix2) Inv() Matrix2 {
	a, b, c, d := mat[0][0], mat[0][1], mat[1][0], mat[1][1]
	D := mat.Det()
	if D == 0 {
		panic(ErrNoInverseMatrix)
	}
	return Matrix2{{d, -b}, {-c, a}}.Scaled(1.0 / mat.Det())
}
