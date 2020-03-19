package base

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"math"
	"strings"
)

type AFVector3D struct {
	X float64
	Y float64
	Z float64
}

func NewAFVector3D(x, y, z float64) *AFVector3D {
	return &AFVector3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func NewAFVector3DFromAFVector3D(v *AFVector3D) *AFVector3D {
	return &AFVector3D{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}

func NewAFVector3DFromString(str string) (*AFVector3D, error) {
	strArr := strings.Split(str, ",")
	if len(strArr) != 3 {
		return nil, errors.New("failed to new AFVector3D from string : " + str)
	}

	var float64Arr [3]float64
	for k, v := range strArr {
		data, err := cast.ToFloat64E(v)
		if err != nil {
			return nil, errors.New("failed to new AFVector3D from string : " + str)
		}
		float64Arr[k] = data
	}

	return &AFVector3D{
		X: float64Arr[0],
		Y: float64Arr[1],
		Z: float64Arr[2],
	}, nil
}

func (v *AFVector3D) ToString() string {
	return fmt.Sprintf("%.2f,%.2f,%.2f", v.X, v.Y, v.Z)
}

func (v *AFVector3D) IsZero() bool {
	return IsZeroFloat64(v.X) && IsZeroFloat64(v.Y) && IsZeroFloat64(v.Z)
}

func (v *AFVector3D) EqualTo(v1 *AFVector3D) bool {
	return false
}

func (v *AFVector3D) NotEqualTo(v1 *AFVector3D) bool {
	return !v.EqualTo(v1)
}

func (v *AFVector3D) Distance(v1 *AFVector3D) float64 {
	dx := v.X - v1.X
	dy := v.Y - v1.Y
	dz := v.Z - v1.Z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
