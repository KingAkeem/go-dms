package dms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDMS(t *testing.T) {
	dms, err := NewDMS(DecimalDegrees{Latitude: 23.33, Longitude: 42.55})
	assert.Nil(t, err)

	// test latitude
	assert.Equal(t, 23, dms.Latitude.degrees)
	assert.Equal(t, 19, dms.Latitude.minutes)
	assert.InDelta(t, 48.0, dms.Latitude.seconds, 1e-9)
	assert.Equal(t, "N", dms.Latitude.direction)
	assert.Equal(t, `23°19'48.000000" N`, dms.Latitude.String())

	// test longitude
	assert.Equal(t, 42, dms.Longitude.degrees)
	assert.Equal(t, 33, dms.Longitude.minutes)
	assert.InDelta(t, 0.0, dms.Longitude.seconds, 1e-9)
	assert.Equal(t, "E", dms.Longitude.direction)
	assert.Equal(t, `42°33'0.000000" E`, dms.Longitude.String())

	assert.Equal(t, `23°19'48.000000" N 42°33'0.000000" E`, dms.String())

	dms, err = NewDMS(DecimalDegrees{Latitude: -66.434323, Longitude: -115.25})
	assert.Nil(t, err)

	// test latitude
	assert.Equal(t, dms.Latitude.degrees, 66)
	assert.Equal(t, dms.Latitude.minutes, 26)
	assert.Equal(t, dms.Latitude.seconds, 3.5628000000223814)
	assert.Equal(t, dms.Latitude.direction, "S")
	assert.Equal(t, dms.Latitude.String(), `66°26'3.5628000000223814" S`)

	// test longitude
	assert.Equal(t, 115, dms.Longitude.degrees)
	assert.Equal(t, 15, dms.Longitude.minutes)
	assert.InDelta(t, 0.0, dms.Longitude.seconds, 1e-9)
	assert.Equal(t, "W", dms.Longitude.direction)
	assert.Equal(t, `115°15'0.000000" W`, dms.Longitude.String())

	assert.Equal(t, `66°26'3.562800" S 115°15'0.000000" W`, dms.String())
}

func TestDMSBoundaryValues(t *testing.T) {
	dms, err := NewDMS(DecimalDegrees{Latitude: 0, Longitude: 0})
	assert.NoError(t, err)
	assert.Equal(t, `0°0'0.000000" N 0°0'0.000000" E`, dms.String())

	dms, err = NewDMS(DecimalDegrees{Latitude: 90, Longitude: 180})
	assert.NoError(t, err)
	assert.Equal(t, `90°0'0.000000" N 180°0'0.000000" E`, dms.String())

	dms, err = NewDMS(DecimalDegrees{Latitude: -90, Longitude: -180})
	assert.NoError(t, err)
	assert.Equal(t, `90°0'0.000000" S 180°0'0.000000" W`, dms.String())
}

func TestDMSInvalidRange(t *testing.T) {
	_, err := NewDMS(DecimalDegrees{Latitude: 91, Longitude: 0})
	assert.ErrorContains(t, err, "latitude must be in range")

	_, err = NewDMS(DecimalDegrees{Latitude: 0, Longitude: 181})
	assert.ErrorContains(t, err, "longitude must be in range")
}
