package coordinate

import "github.com/KingAkeem/go-dms/latlon"

type Coordinate interface {
	GetDecimalDegrees() latlon.LatLon
	String() string
}
