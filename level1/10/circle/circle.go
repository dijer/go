package circle

import (
	"math"
)

func calcDiameter(area float64) float64 {
	return 2 * math.Sqrt(area/math.Pi)
}

func calcCircumference(diameter float64) float64 {
	return math.Pi * diameter
}
