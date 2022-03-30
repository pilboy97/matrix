package matrix

import (
	"math"

	"golang.org/x/exp/constraints"
)

const Epsilon = 0.00000001

func InRange[T constraints.Ordered](st, ed, value T) bool {
	if st > ed {
		return InRange(ed, st, value)
	}

	return st <= value && value <= ed
}

func ASame(x, y float64) bool {
	return math.Abs(x-y) <= Epsilon
}
