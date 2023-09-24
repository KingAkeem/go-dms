package dms

import (
	"fmt"
	"math"
)

// DecimalDegrees represent latitude/longitude of geographic coordinates as decimal frations of a degree.
type DecimalDegrees struct {
	Latitude  float64
	Longitude float64
}

// DMSAngle represents a single angle for degrees, miutes, seconds measurements
type DMSAngle struct {
	degrees   int
	minutes   int
	seconds   float64
	direction string
}

func (d DMSAngle) String() string {
	return fmt.Sprintf(`%d°%d'%v" %s`, d.degrees, d.minutes, d.seconds, d.direction)
}

// DMS coordinate
type DMS struct {
	Latitude  DMSAngle
	Longitude DMSAngle
}

func (d DMS) String() string {
	return fmt.Sprintf(`%s %s`, d.Latitude, d.Longitude)
}

func newDMSAngle(decimalDegreeAngle float64, direction string) DMSAngle {
	decimalDegreeAngle = math.Abs(decimalDegreeAngle)
	degrees := uint8(decimalDegreeAngle)
	minutes := uint8((decimalDegreeAngle - float64(degrees)) * 60)
	seconds := (decimalDegreeAngle - float64(degrees) - float64(minutes)/60) * 3600

	return DMSAngle{
		degrees:   int(degrees),
		minutes:   int(minutes),
		seconds:   seconds,
		direction: direction,
	}
}

// New generates a DMS position from a set of decimal degree coordinates (latitude/longitude)
func NewDMS(latlon DecimalDegrees) (*DMS, error) {
	lat, lon := latlon.Latitude, latlon.Longitude

	var latDirection, lonDirection string
	if lat > 0 {
		latDirection = "N"
	} else {
		latDirection = "S"
	}

	if lon > 0 {
		lonDirection = "E"
	} else {
		lonDirection = "W"
	}

	if lat < -90 || lat > 90 {
		return nil, fmt.Errorf("latitude must be in range of -90 and 90, found %f", lat)
	}

	if lon < -180 || lon > 180 {
		return nil, fmt.Errorf("longitude must be in range of -180 and 180, foujnd %f", lon)
	}

	latitude := newDMSAngle(lat, latDirection)
	longitude := newDMSAngle(lon, lonDirection)

	return &DMS{Latitude: latitude, Longitude: longitude}, nil
}
