package matrix

import (
	"math"
)

func MakeADot3(x, y, z float64) Vector4 {
	return MakeVector4(x, y, z, 1)
}
func MakeAVector3(x, y, z float64) Vector4 {
	return MakeVector4(x, y, z, 0)
}

//identity transform for 3d affine transform
func Identity3() Matrix4 {
	return UnitMatrix4()
}

//move transform for 3d affine transform
func Moved3(x, y, z float64) Matrix4 {
	return MakeMatrix4([16]float64{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	})
}

//scale transform for 3d affine transform
func Scaled3(x, y, z float64) Matrix4 {
	return MakeMatrix4([16]float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	})
}

//rotate transform around the x-axis
func RotatedX3(rad float64) Matrix4 {
	c, s := math.Cos(rad), math.Sin(rad)
	return MakeMatrix4([16]float64{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	})
}

//rotate transform around the y-axis
func RotatedY3(rad float64) Matrix4 {
	c, s := math.Cos(rad), math.Sin(rad)
	return MakeMatrix4([16]float64{
		c, 0, -s, 0,
		0, 1, 0, 0,
		s, 0, c, 0,
		0, 0, 0, 1,
	})
}

//rotate transform around the z-axis
func RotatedZ3(rad float64) Matrix4 {
	c, s := math.Cos(rad), math.Sin(rad)
	return MakeMatrix4([16]float64{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
}

//perspective transform
//	W : width
//	H : height
//	F : Far
//	N : near
func Perspective(W, H, F, N float64) Matrix4 {
	return MakeMatrix4([16]float64{
		2 * N / W, 0, 0, 0,
		0, 2 * N / H, 0, 0,
		0, 0, F / (F - N), -F * N / (F - N),
		0, 0, 1, 0,
	})
}

func (v Vector4) H() Vector4 {
	if v[3] == 0 {
		return v
	}
	return MakeVector4(v[0]/v[3], v[1]/v[3], v[2]/v[3], 1)
}
func (mat Matrix4) Proj(v Vector4) Vector4 {
	res := mat.Transform(v)
	return MakeVector4(res[0]/res[3], res[1]/res[3], res[2]/res[3], 1)
}

func (mat Matrix4) Chained(mat2 Matrix4) Matrix4 {
	return mat2.Mul(mat)
}
