package matrix

import (
	"math"
)

type Vector3 [3]float64

func MakeVector3(x, y, z float64) Vector3 {
	return [3]float64{x, y, z}
}

func (v Vector3) Equal(v2 Vector3) bool {
	for i := 0; i < 3; i++ {
		if !Equal(v[i], v2[i]) {
			return false
		}
	}
	return true
}
func (v Vector3) Scaled(C float64) Vector3 {
	return MakeVector3(v[0]*C, v[1]*C, v[2]*C)
}
func (v Vector3) Add(v2 Vector3) Vector3 {
	return MakeVector3(v[0]+v2[0], v[1]+v2[1], v[2]+v2[2])
}
func (v Vector3) Sub(v2 Vector3) Vector3 {
	return v.Add(v2.Scaled(-1))
}
func (v Vector3) Dot(v2 Vector3) float64 {
	return v[0]*v2[0] + v[1]*v2[1] + v[2]*v2[2]
}
func (v Vector3) Len() float64 {
	return math.Sqrt(v.Dot(v))
}
func (v Vector3) Angle(v2 Vector3) float64 {
	if v.Len()*v2.Len() == 0 {
		panic(ErrZeroLength)
	}

	return math.Acos(v.Dot(v2) / (v.Len() * v2.Len()))
}
func (v Vector3) Cross(v2 Vector3) Vector3 {
	return MakeVector3(
		v[1]*v2[2]-v[2]*v2[2],
		v[2]*v2[0]-v[0]*v2[2],
		v[0]*v2[1]-v[1]*v2[0])
}

func (mat Matrix3) Transform(v Vector3) Vector3 {
	return MakeVector3(
		mat[0][0]*v[0]+mat[0][1]*v[1]+mat[0][2]*v[2],
		mat[1][0]*v[0]+mat[1][1]*v[1]+mat[1][2]*v[2],
		mat[2][0]*v[0]+mat[2][1]*v[1]+mat[2][2]*v[2])
}
