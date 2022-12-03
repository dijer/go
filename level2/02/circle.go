// Package circle is calculate diameter and circum ference
package circle

import (
	"math"
)

// Length is a float length fro calcDiameter and calcCircumference
type Length = float64

// Calculate Diameter
// @Param area Length
// @Return Length
func calcDiameter(area Length) Length {
	return 2 * math.Sqrt(area/math.Pi)
}

// Calculate circum ference
// @Param diameter Length
// @Return Length
func calcCircumference(diameter Length) Length {
	return math.Pi * diameter
}
