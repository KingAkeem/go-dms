package dms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDMS(t *testing.T) {
	lat := 23.33
	lon := 42.55
	dms, err := New(lat, lon)
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

	// test full DMS
	assert.Equal(t, dms.String(), `23°19'47.99999999999392" N 42°32'59.9999999999898" E`)
}
