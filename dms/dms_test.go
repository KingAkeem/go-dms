package dms

import (
	"math"
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
	assert.Equal(t, 66, dms.Latitude.degrees)
	assert.Equal(t, 26, dms.Latitude.minutes)
	assert.InDelta(t, 3.5628, dms.Latitude.seconds, 1e-9)
	assert.Equal(t, "S", dms.Latitude.direction)
	assert.Equal(t, `66°26'3.562800" S`, dms.Latitude.String())

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

// Run with: go test -fuzz=FuzzNewDMS -fuzztime=30s ./dms
func FuzzNewDMS(f *testing.F) {
	seeds := []DecimalDegrees{
		{Latitude: -90, Longitude: -180},
		{Latitude: math.Nextafter(-90, 0), Longitude: math.Nextafter(-180, 0)},
		{Latitude: math.Nextafter(-90, math.Inf(-1)), Longitude: math.Nextafter(-180, math.Inf(-1))},
		{Latitude: 0, Longitude: 0},
		{Latitude: math.Nextafter(0, math.Inf(-1)), Longitude: math.Nextafter(0, math.Inf(-1))},
		{Latitude: math.Nextafter(0, math.Inf(1)), Longitude: math.Nextafter(0, math.Inf(1))},
		{Latitude: 90, Longitude: 180},
		{Latitude: math.Nextafter(90, 0), Longitude: math.Nextafter(180, 0)},
		{Latitude: math.Nextafter(90, math.Inf(1)), Longitude: math.Nextafter(180, math.Inf(1))},
		{Latitude: -180, Longitude: -90},
		{Latitude: 180, Longitude: 90},
	}
	for _, seed := range seeds {
		f.Add(seed.Latitude, seed.Longitude)
	}

	assertNormalized := func(t *testing.T, name string, angle DMSAngle, maxDegrees int) {
		t.Helper()
		if angle.degrees < 0 || angle.degrees > maxDegrees {
			t.Fatalf("%s degrees are not normalized: %d", name, angle.degrees)
		}
		if angle.minutes < 0 || angle.minutes >= 60 {
			t.Fatalf("%s minutes are not normalized: %d", name, angle.minutes)
		}
		if math.IsNaN(angle.seconds) || math.IsInf(angle.seconds, 0) || angle.seconds < 0 || angle.seconds >= 60 {
			t.Fatalf("%s seconds are not normalized: %f", name, angle.seconds)
		}
		if angle.degrees == maxDegrees && (angle.minutes != 0 || angle.seconds != 0) {
			t.Fatalf("%s exceeds its boundary: %d degrees, %d minutes, %f seconds", name, angle.degrees, angle.minutes, angle.seconds)
		}
	}

	f.Fuzz(func(t *testing.T, latitude, longitude float64) {
		if math.IsNaN(latitude) || math.IsInf(latitude, 0) || math.IsNaN(longitude) || math.IsInf(longitude, 0) {
			t.Skip()
		}

		got, err := NewDMS(DecimalDegrees{Latitude: latitude, Longitude: longitude})
		if latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
			if err == nil {
				t.Fatalf("NewDMS(%f, %f) returned no error for out-of-range input", latitude, longitude)
			}
			return
		}

		if err != nil {
			t.Fatalf("NewDMS(%f, %f) returned an unexpected error: %v", latitude, longitude, err)
		}
		if got == nil {
			t.Fatalf("NewDMS(%f, %f) returned a nil result", latitude, longitude)
		}

		assertNormalized(t, "latitude", got.Latitude, 90)
		assertNormalized(t, "longitude", got.Longitude, 180)
	})
}
