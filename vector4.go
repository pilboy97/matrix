package matrix

import (
	"math"
)

type Vector4 [4]float64

func MakeVector4(x, y, z, w float64) Vector4 {
	return [4]float64{x, y, z, w}
}

func (v Vector4) Equal(v2 Vector4) bool {
	for i := 0; i < 4; i++ {
		if !Equal(v[i], v2[i]) {
			return false
		}
	}
	return true
}
func (v Vector4) Scaled(C float64) Vector4 {
	return MakeVector4(v[0]*C, v[1]*C, v[2]*C, v[3]*C)
}
func (v Vector4) Add(v2 Vector4) Vector4 {
	return MakeVector4(v[0]+v2[0], v[1]+v2[1], v[2]+v2[2], v[3]+v2[3])
}
func (v Vector4) Sub(v2 Vector4) Vector4 {
	return v.Add(v2.Scaled(-1))
}
func (v Vector4) Dot(v2 Vector4) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2] + v[3]*v2[3]
}
func (v Vector4) Len() float64 {
	return math.Sqrt(v.Dot(v))
}
func (v Vector4) Angle(v2 Vector4) float64 {
	if v.Len()*v2.Len() == 0 {
		panic(ErrZeroLength)
	}

	return math.Acos(v.Dot(v2) / (v.Len() * v2.Len()))
}

func (mat Matrix4) Transform(v Vector4) Vector4 {
	return MakeVector4(
		mat[0][0]*v[0]+mat[0][1]*v[1]+mat[0][2]*v[2]+mat[0][3]*v[3],
		mat[1][0]*v[0]+mat[1][1]*v[1]+mat[1][2]*v[2]+mat[1][3]*v[3],
		mat[2][0]*v[0]+mat[2][1]*v[1]+mat[2][2]*v[2]+mat[2][3]*v[3],
		mat[3][0]*v[0]+mat[3][1]*v[1]+mat[3][2]*v[2]+mat[3][3]*v[3])
}
