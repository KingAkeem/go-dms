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

	assert.Equal(t, dms.Latitude.degrees, 23)
	assert.Equal(t, dms.Latitude.minutes, 19)
	assert.Equal(t, dms.Latitude.seconds, 47.99999999999392)
	assert.Equal(t, dms.Latitude.direction, "N")

	assert.Equal(t, dms.Latitude.String(), `23°19'47.99999999999392" N`)
}
