package latlon

import "fmt"

// LatLon represent latitude/longitude of geographic coordinates as decimal frations of a degree.
type LatLon struct {
	Latitude  float64
	Longitude float64
}

func (l LatLon) String() string {
	return fmt.Sprintf("%f°, %f°", l.Latitude, l.Longitude)
}

func (l LatLon) GetDecimalDegrees() LatLon {
	return l
}
