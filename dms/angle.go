package dms

import (
	"fmt"
	"math"
)

// Angle represents a single angle for degrees, miutes, seconds measurements
type Angle struct {
	degrees   int
	minutes   int
	seconds   float64
	direction string
}

// String representation of an Angle
func (a Angle) String() string {
	return fmt.Sprintf(`%d°%d'%v" %s`, a.degrees, a.minutes, a.seconds, a.direction)
}

func newAngle(decimalDegreeAngle float64, direction string) Angle {
	decimalDegreeAngle = math.Abs(decimalDegreeAngle)
	degrees := uint8(decimalDegreeAngle)
	minutes := uint8((decimalDegreeAngle - float64(degrees)) * 60)
	seconds := (decimalDegreeAngle - float64(degrees) - float64(minutes)/60) * 3600

	return Angle{
		degrees:   int(degrees),
		minutes:   int(minutes),
		seconds:   seconds,
		direction: direction,
	}
}

func convertAngleToDecimalDegrees(a Angle) float64 {
	minuteDegrees := float64(a.minutes) / 60
	secondDegrees := float64(a.seconds) / 3600
	decimalDegrees := float64(a.degrees) + float64(minuteDegrees) + secondDegrees

	if a.direction == "S" || a.direction == "W" {
		return -decimalDegrees
	}

	return decimalDegrees
}
