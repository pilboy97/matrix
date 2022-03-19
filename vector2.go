package matrix

import (
	"errors"
	"math"
)

var ErrZeroLength = errors.New("vector has no length")

type Vector2 [2]float64

func MakeVector2(x, y float64) Vector2 {
	return [2]float64{x, y}
}

func (v Vector2) Equal(v2 Vector2) bool {
	for i := 0; i < 2; i++ {
		if !Equal(v[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v Vector2) Scaled(C float64) Vector2 {
	return MakeVector2(v[0]*C, v[1]*C)
}
func (v Vector2) Add(v2 Vector2) Vector2 {
	return MakeVector2(v[0]+v2[0], v[1]+v2[1])
}
func (v Vector2) Sub(v2 Vector2) Vector2 {
	return v.Add(v2.Scaled(-1))
}
func (v Vector2) Dot(v2 Vector2) float64 {
	return v[0]*v2[0] + v[1]*v2[1]
}
func (v Vector2) Len() float64 {
	return math.Sqrt(v.Dot(v))
}
func (v Vector2) Angle(v2 Vector2) float64 {
	if v.Len()*v2.Len() == 0 {
		panic(ErrZeroLength)
	}

	return math.Acos(v.Dot(v2) / (v.Len() * v2.Len()))
}

func (mat Matrix2) Transform(v Vector2) Vector2 {
	return MakeVector2(
		mat[0][0]*v[0]+mat[0][1]*v[1],
		mat[1][0]*v[0]+mat[1][1]*v[1])
}
