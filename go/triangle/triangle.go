package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

// Pick values for the following identifiers used by the test program.
const NaT = 0 // not a triangle
const Equ = 1 // equilateral
const Iso = 2 // isosceles
const Sca = 3 // scalene

// Organize your code for readability.

type Kind int

func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}

	sort.Sort(sort.Float64Slice(sides))

	switch {
	case invalidSide(sides):
		return NaT
	case isEqual(sides):
		return Equ
	case isIsoceles(sides):
		return Iso
	case !isInEqualTriangle(sides):
		return NaT
	default:
		return Sca
	}
}

func invalidSide(sides []float64) bool {
	for _, side := range sides {
		if math.IsNaN(side) || math.Inf(1) == side || math.Inf(-1) == side || side <= 0 {
			return true
		}
	}
	return false
}

func isIsoceles(sides []float64) bool {
	return (sides[2] == sides[1] || sides[1] == sides[0]) && isInEqualTriangle(sides)
}

func isInEqualTriangle(sides []float64) bool {
	return (sides[0] + sides[1]) >= sides[2]
}

func isEqual(sides []float64) bool {
	return sides[2] == sides[0]
}
