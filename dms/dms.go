package dms

import (
	"fmt"

	"github.com/KingAkeem/go-dms/coordinate"
	"github.com/KingAkeem/go-dms/latlon"
)

// DMS coordinate
type DMS struct {
	Latitude  Angle
	Longitude Angle
}

// GetDecimalDegrees returns the GPS coordinate for a DMS position
func (d DMS) GetDecimalDegrees() latlon.LatLon {
	return latlon.LatLon{
		Latitude:  convertAngleToDecimalDegrees(d.Latitude),
		Longitude: convertAngleToDecimalDegrees(d.Longitude),
	}
}

// String returns the representation of a DMS value
func (d DMS) String() string {
	return fmt.Sprintf(`%s %s`, d.Latitude, d.Longitude)
}

// New generates a DMS position from a coordinate using it's decimal degrees value
func New(coord coordinate.Coordinate) (*DMS, error) {
	latlon := coord.GetDecimalDegrees()
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
		return nil, fmt.Errorf("longitude must be in range of -180 and 180, found %f", lon)
	}

	latitude := newAngle(lat, latDirection)
	longitude := newAngle(lon, lonDirection)

	return &DMS{Latitude: latitude, Longitude: longitude}, nil
}
