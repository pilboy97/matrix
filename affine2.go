package matrix

import "math"

func MakeADot2(x, y float64) Vector3 {
	return MakeVector3(x, y, 1)
}
func MakeAVector2(x, y float64) Vector3 {
	return MakeVector3(x, y, 0)
}

//Identity transform for 2d affine transform
func Identity2() Matrix3 {
	return UnitMatrix3()
}

//move transform for 2d affine transform
func Moved2(x, y float64) Matrix3 {
	return MakeMatrix3([9]float64{
		1, 0, x,
		0, 1, y,
		0, 0, 1,
	})
}

//scale transform for 2d affine transform
func Scaled2(x, y float64) Matrix3 {
	return MakeMatrix3([9]float64{
		x, 0, 0,
		0, y, 0,
		0, 0, 1,
	})
}

//rotate transform for 2d affine transform
func Rotated2(rad float64) Matrix3 {
	c, s := math.Cos(rad), math.Sin(rad)
	return MakeMatrix3([9]float64{
		c, -s, 0,
		s, c, 0,
		0, 0, 1,
	})
}

func (v Vector3) H() Vector3 {
	if v[2] == 0 {
		return v
	}
	return MakeVector3(v[0]/v[2], v[1]/v[2], 1)
}

func (mat Matrix3) Chained(mat2 Matrix3) Matrix3 {
	return mat2.Mul(mat)
}

//affine transform for the given vector
func (mat Matrix3) Proj(v Vector3) Vector3 {
	return mat.Transform(v)
}
