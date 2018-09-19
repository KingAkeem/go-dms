package dms 

type LatLonError struct {
    err string
}

func (e *LatLonError) Error() string {
    return e.err 
}

type DMS struct {
    degrees uint8
    minutes uint8
    seconds float64 
}

func NewDMS(lat, lon float64) (*DMS, *DMS, error) {
    if lat < 0 || lon < 0 {
        return nil, nil, &LatLonError{"Latitude or longitude must be positive."} 
    }
    if lat > 90 || lon > 180 {
        return nil, nil, &LatLonError{"Latitude must be less than 90 and longitude must be less than 180."}
    }

    latitude := uint8(lat)
    latitudeMinutes := uint8((lat - float64(latitude)) * 60)
    latitudeSeconds := (lat - float64(latitude) - float64(latitudeMinutes) / 60) * 3600

    longitude := uint8(lon)
    longitudeMinutes := uint8((lon - float64(longitude)) * 60)
    longitudeSeconds := (lon - float64(longitude) - float64(longitudeMinutes) / 60) * 3600

    return &DMS{degrees: latitude, minutes: latitudeMinutes, seconds: latitudeSeconds,}, 
           &DMS{degrees: longitude, minutes: longitudeMinutes, seconds: longitudeSeconds,}, 
           nil
}
