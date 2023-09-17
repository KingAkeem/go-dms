package dms

import (
	"fmt"
)

// LatLonError is used for errors with lat/lon values
type LatLonError struct {
	err string
}

func (e *LatLonError) Error() string {
	return e.err
}

// Position is a single coordinate within a DMS location
type Position struct {
	degrees   int
	minutes   int
	seconds   float64
	direction string
}

// DMS coordinate
type DMS struct {
	Latitude  Position
	Longitude Position
}

func (p Position) String() string {
	return fmt.Sprintf(`%d°%d'%v" %s`, p.degrees, p.minutes, p.seconds, p.direction)
}

func newPosition(decimalDegrees float64, direction string) Position {
	degrees := uint8(decimalDegrees)
	minutes := uint8((decimalDegrees - float64(degrees)) * 60)
	seconds := (decimalDegrees - float64(degrees) - float64(minutes)/60) * 3600

	return Position{
		degrees:   int(degrees),
		minutes:   int(minutes),
		seconds:   seconds,
		direction: direction,
	}
}

// New converts Decimal Degreees to Degree, Minute, Seconds coordinates
func New(lat, lon float64) (*DMS, error) {
	if lat < 0 || lon < 0 {
		return nil, &LatLonError{"latitude or longitude must be positive"}
	}
	if lat > 90 || lon > 180 {
		return nil, &LatLonError{"latitude must be less than 90 and longitude must be less than 180"}
	}

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

	dmsLat := newPosition(lat, latDirection)
	dmsLon := newPosition(lon, lonDirection)

	return &DMS{Latitude: dmsLat, Longitude: dmsLon}, nil
}
