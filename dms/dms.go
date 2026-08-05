package dms

import (
	"fmt"
	"math"
)

// DecimalDegrees represent latitude/longitude of geographic coordinates as decimal fractions of a degree.
type DecimalDegrees struct {
	Latitude  float64
	Longitude float64
}

// DMSAngle represents a single angle for degrees, minutes, seconds measurements
type DMSAngle struct {
	degrees   int
	minutes   int
	seconds   float64
	direction string
}

func (d DMSAngle) String() string {
	return fmt.Sprintf("%d°%d'%.6f\" %s", d.degrees, d.minutes, d.seconds, d.direction)
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
	totalSeconds := math.Abs(decimalDegreeAngle) * 3600
	degrees := int(totalSeconds / 3600)
	remainingSeconds := totalSeconds - float64(degrees)*3600
	minutes := int(remainingSeconds / 60)
	seconds := remainingSeconds - float64(minutes)*60

	if seconds >= 59.9999999999 {
		seconds = 0
		minutes++
	}

	if minutes >= 60 {
		minutes = 0
		degrees++
	}

	return DMSAngle{
		degrees:   degrees,
		minutes:   minutes,
		seconds:   seconds,
		direction: direction,
	}
}

// New generates a DMS position from a set of decimal degree coordinates (latitude/longitude)
func NewDMS(latlon DecimalDegrees) (*DMS, error) {
	lat, lon := latlon.Latitude, latlon.Longitude

	var latDirection, lonDirection string
	if lat >= 0 {
		latDirection = "N"
	} else {
		latDirection = "S"
	}

	if lon >= 0 {
		lonDirection = "E"
	} else {
		lonDirection = "W"
	}

	if lat < -90 || lat > 90 {
		return nil, fmt.Errorf("latitude must be in range of -90 and 90, found %f", lat)
	}

	if lon < -180 || lon > 180 {
		return nil, fmt.Errorf("longitude must be in range of -180 and 180, found %f", lon)
	}

	latitude := newDMSAngle(lat, latDirection)
	longitude := newDMSAngle(lon, lonDirection)

	return &DMS{Latitude: latitude, Longitude: longitude}, nil
}
