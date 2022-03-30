package matrix

import (
	"errors"
	"fmt"
	"strings"
)

var ErrWrongSize = errors.New("size of the matrix is wrong")
var ErrNotSquareMatrix = errors.New("the matrix is not a square matrix")
var ErrHasNoInvereMatix = errors.New("the matrix has no inverse matrix")

type Mat struct {
	r, c int
	arr  []float64
}

func MakeMat(r, c int, elem ...float64) Mat {
	if r*c != len(elem) {
		panic(ErrWrongSize)
	}

	return Mat{
		r:   r,
		c:   c,
		arr: elem,
	}
}
func JoinRow(mat ...Mat) Mat {
	var R, C int
	R = mat[0].r
	C = mat[0].c

	for i := 1; i < len(mat); i++ {
		if mat[i].c != C {
			panic(ErrWrongSize)
		}

		R += mat[i].r
	}

	result := Mat{
		r:   R,
		c:   C,
		arr: make([]float64, R*C),
	}

	var cnt = 0

	for _, m := range mat {
		for i := 0; i < m.r; i++ {
			for j := 0; j < m.c; j++ {
				result.arr[(i+cnt)*C+j] = m.arr[i*C+j]
			}
			cnt++
		}
	}

	return result
}
func JoinCol(mat ...Mat) Mat {
	var R, C int
	R = mat[0].r
	C = mat[0].c

	for i := 1; i < len(mat); i++ {
		if mat[i].r != R {
			panic(ErrWrongSize)
		}

		C += mat[i].c
	}

	result := Mat{
		r:   R,
		c:   C,
		arr: make([]float64, R*C),
	}

	var cnt = 0

	for _, m := range mat {
		for i := 0; i < m.r; i++ {
			for j := 0; j < m.c; j++ {
				result.arr[i*C+j+cnt] = m.arr[i*m.c+j]
			}
		}
		cnt += m.c
	}

	return result
}

func (m Mat) R() int {
	return m.r
}
func (m Mat) C() int {
	return m.c
}
func (m Mat) Elem(r, c int) float64 {
	return m.arr[m.c*r+c]
}
func (m Mat) Row(r int) Mat {
	result := Mat{
		r:   1,
		c:   m.c,
		arr: make([]float64, m.c),
	}
	for i := 0; i < m.c; i++ {
		result.arr[i] = m.arr[r*m.c+i]
	}

	return result
}
func (m Mat) Col(c int) Mat {
	result := Mat{
		r:   m.r,
		c:   1,
		arr: make([]float64, m.r),
	}
	for i := 0; i < m.r; i++ {
		result.arr[i] = m.arr[i*m.c+c]
	}

	return result
}

func (m Mat) String() string {
	result := make([]string, m.r+2)

	result[0] = "["
	for i := 0; i < m.r; i++ {
		var rows []string
		rows = append(rows, "[")

		for j := 0; j < m.c; j++ {
			rows = append(rows, fmt.Sprintf("%.5g", m.Elem(i, j)))
			if j+1 < m.c {
				rows = append(rows, ",")
			}
		}

		rows = append(rows, "]")
		if i+1 < m.r {
			rows = append(rows, ",")
		}

		result[i+1] = strings.Join(rows, "")
	}
	result[m.r+1] = "]"

	return strings.Join(result, "")
}

func (m Mat) Scaled(C float64) Mat {
	result := Mat{
		r:   m.r,
		c:   m.c,
		arr: make([]float64, len(m.arr)),
	}

	for i := 0; i < m.r; i++ {
		for j := 0; j < m.c; j++ {
			result.arr[i*m.c+j] = m.Elem(i, j) * C
		}
	}

	return result
}

func (m Mat) Equal(m2 Mat) bool {
	if m.r != m2.r || m.c != m2.c {
		panic(ErrWrongSize)
	}

	for i := 0; i < m.r; i++ {
		for j := 0; j < m.c; j++ {
			if !ASame(m.Elem(i, j), m2.Elem(i, j)) {
				return false
			}
		}
	}

	return true
}

func (m Mat) Add(m2 Mat) Mat {
	if m.r != m2.r || m.c != m2.c {
		panic(ErrWrongSize)
	}

	result := Mat{
		r:   m.r,
		c:   m.c,
		arr: make([]float64, len(m.arr)),
	}

	for i := 0; i < m.r; i++ {
		for j := 0; j < m.c; j++ {
			result.arr[i*m.c+j] = m.Elem(i, j) + m2.Elem(i, j)
		}
	}

	return result
}

func (m Mat) Sub(m2 Mat) Mat {
	return m.Add(m2.Scaled(-1))
}

func (m Mat) Mul(m2 Mat) Mat {
	if m.c != m2.r {
		panic(ErrWrongSize)
	}

	result := Mat{
		r:   m.r,
		c:   m2.c,
		arr: make([]float64, m.r*m2.c),
	}

	for i := 0; i < m.r; i++ {
		for j := 0; j < m2.c; j++ {
			sum := 0.0
			for k := 0; k < m.c; k++ {
				sum += m.Elem(i, k) * m2.Elem(k, j)
			}
			result.arr[m.c*i+j] = sum
		}
	}

	return result
}

func (m Mat) adj(r, c int) Mat {
	result := Mat{
		r:   m.r - 1,
		c:   m.c - 1,
		arr: make([]float64, (m.r-1)*(m.c-1)),
	}

	var R = 0
	for i := 0; i < m.r; i++ {
		var C = 0
		for j := 0; j < m.c; j++ {
			if !(i == r || j == c) {
				result.arr[R*(m.c-1)+C] = m.Elem(i, j)
				C++
			}
		}
		if i != r {
			R++
		}
	}

	return result
}
func (m Mat) Det() float64 {
	if m.r != m.c {
		panic(ErrNotSquareMatrix)
	}

	if m.r == 1 {
		return m.Elem(0, 0)
	}
	if m.r == 2 {
		return m.Elem(0, 0)*m.Elem(1, 1) - m.Elem(0, 1)*m.Elem(1, 0)
	}

	var sum = 0.0

	for i := 0; i < m.r; i++ {
		if i%2 == 0 {
			sum += m.Elem(i, 0) * m.adj(i, 0).Det()
		} else {
			sum -= m.Elem(i, 0) * m.adj(i, 0).Det()
		}
	}

	return sum
}
func (m Mat) Trans() Mat {
	result := Mat{
		r:   m.c,
		c:   m.r,
		arr: make([]float64, m.r*m.c),
	}

	for i := 0; i < m.c; i++ {
		for j := 0; j < m.r; j++ {
			result.arr[i*m.c+j] = m.Elem(j, i)
		}
	}

	return result
}
func (m Mat) Inv() Mat {
	if m.r != m.c {
		panic(ErrNotSquareMatrix)
	}

	if m.r == 1 {
		return MakeMat(1, 1, 1/m.Elem(0, 0))
	}
	if m.r == 2 {
		A, B, C, D := m.Elem(0, 0), m.Elem(0, 1), m.Elem(1, 0), m.Elem(1, 1)
		return MakeMat(2, 2, D, -B, -C, A).Scaled(1 / (A*D - B*C))
	}

	I := m.Det()
	if I == 0 {
		panic(ErrHasNoInvereMatix)
	}

	result := Mat{
		r:   m.r,
		c:   m.c,
		arr: make([]float64, m.r*m.c),
	}

	for i := 0; i < m.r; i++ {
		for j := 0; j < m.c; j++ {
			result.arr[i*m.c+j] = m.adj(i, j).Det()
			if (i+j)%2 == 1 {
				result.arr[i*m.r+j] *= -1
			}
		}
	}

	return result.Trans().Scaled(1 / I)
}
