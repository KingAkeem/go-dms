package dms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDMS(t *testing.T) {
	dms, err := NewDMS(DecimalDegrees{Latitude: 23.33, Longitude: 42.55})
	assert.Nil(t, err)

	// test latitude
	assert.Equal(t, dms.Latitude.degrees, 23)
	assert.Equal(t, dms.Latitude.minutes, 19)
	assert.Equal(t, dms.Latitude.seconds, 47.99999999999392)
	assert.Equal(t, dms.Latitude.direction, "N")
	assert.Equal(t, dms.Latitude.String(), `23°19'47.99999999999392" N`)

	// test longitude
	assert.Equal(t, dms.Longitude.degrees, 42)
	assert.Equal(t, dms.Longitude.minutes, 32)
	assert.Equal(t, dms.Longitude.seconds, 59.9999999999898)
	assert.Equal(t, dms.Longitude.direction, "E")
	assert.Equal(t, dms.Longitude.String(), `42°32'59.9999999999898" E`)

	assert.Equal(t, dms.String(), `23°19'47.99999999999392" N 42°32'59.9999999999898" E`)

	dms, err = NewDMS(DecimalDegrees{Latitude: -66.434323, Longitude: -115.25})
	assert.Nil(t, err)

	// test latitude
	assert.Equal(t, dms.Latitude.degrees, 66)
	assert.Equal(t, dms.Latitude.minutes, 26)
	assert.Equal(t, dms.Latitude.seconds, 3.5628000000223814)
	assert.Equal(t, dms.Latitude.direction, "S")
	assert.Equal(t, dms.Latitude.String(), `66°26'3.5628000000223814" S`)

	// test longitude
	assert.Equal(t, dms.Longitude.degrees, 115)
	assert.Equal(t, dms.Longitude.minutes, 15)
	assert.Equal(t, dms.Longitude.seconds, 0.0)
	assert.Equal(t, dms.Longitude.direction, "W")
	assert.Equal(t, dms.Longitude.String(), `115°15'0" W`)

	assert.Equal(t, dms.String(), `66°26'3.5628000000223814" S 115°15'0" W`)
}
