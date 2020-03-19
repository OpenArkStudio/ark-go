package base

import "math"

func GetNearest2N() {

}

func IsZeroFloat32(value float32) bool {
	return math.Abs(float64(value)) < math.SmallestNonzeroFloat32
}

func IsZeroFloat64(value float64) bool {
	return math.Abs(value) < math.SmallestNonzeroFloat64
}

func IsFloat32Equal(lhs, rhs float32) bool {
	return math.Abs(float64(lhs-rhs)) < math.SmallestNonzeroFloat32
}

func IsFloat64Equal(lhs, rhs float64) bool {
	return math.Abs(lhs-rhs) < math.SmallestNonzeroFloat64
}
