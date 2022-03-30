package matrix

import "math"

func MakeDot2(x, y float64) Mat {
	return MakeMat(1, 3,
		x, y, 1,
	)
}
func MakeVec2(x, y float64) Mat {
	return MakeMat(1, 3,
		x, y, 0,
	)
}
func MakeDot3(x, y, z float64) Mat {
	return MakeMat(1, 4,
		x, y, z, 1,
	)
}
func MakeVec3(x, y, z float64) Mat {
	return MakeMat(1, 4,
		x, y, z, 0,
	)
}

func I2() Mat {
	return MakeMat(3, 3,
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)
}
func I3() Mat {
	return MakeMat(4, 4,
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}
func (m Mat) Scaled2(x, y float64) Mat {
	return m.Mul(MakeMat(3, 3,
		x, 0, 0,
		0, y, 0,
		0, 0, 1,
	))
}
func (m Mat) Moved2(x, y float64) Mat {
	return m.Mul(MakeMat(3, 3,
		1, 0, 0,
		0, 1, 0,
		x, y, 1,
	))
}
func (m Mat) Rotated2(rad float64) Mat {
	c, s := math.Cos(rad), math.Sin(rad)

	return m.Mul(MakeMat(3, 3,
		c, s, 0,
		-s, c, 0,
		0, 0, 1,
	))
}

func (m Mat) Scaled3(x, y, z float64) Mat {
	return m.Mul(MakeMat(4, 4,
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	))
}
func (m Mat) Moved3(x, y, z float64) Mat {
	return m.Mul(MakeMat(4, 4,
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		x, y, z, 1,
	))
}
func (m Mat) RotatedX3(rad float64) Mat {
	c, s := math.Cos(rad), math.Sin(rad)

	return m.Mul(MakeMat(4, 4,
		1, 0, 0, 0,
		0, c, s, 0,
		0, -s, c, 0,
		0, 0, 0, 1,
	))
}
func (m Mat) RotatedY3(rad float64) Mat {
	c, s := math.Cos(rad), math.Sin(rad)

	return m.Mul(MakeMat(4, 4,
		c, 0, -s, 0,
		0, 1, 0, 0,
		s, 0, c, 0,
		0, 0, 0, 1,
	))
}
func (m Mat) RotatedZ3(rad float64) Mat {
	c, s := math.Cos(rad), math.Sin(rad)

	return m.Mul(MakeMat(4, 4,
		c, s, 0, 0,
		-s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	))
}
