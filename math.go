package matrix

import "math"

const Epsilon = 0.00000001

func Equal(v, v2 float64) bool {
	return math.Abs(v-v2) < Epsilon
}
