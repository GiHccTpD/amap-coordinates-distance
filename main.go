package amap_coordinates_distance

import (
	"math"
)

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

const constant = 0.01745329251994329

func (origin Coordinates) Distance(destination Coordinates) float64 {
	// https://lbs.amap.com/demo/javascript-api/example/calcutation/calculate-distance-between-two-markers

	oLongitude := origin.Longitude
	oLatitude := origin.Latitude

	dLongitude := destination.Longitude
	dLatitude := destination.Latitude

	oLongitude *= constant
	oLatitude *= constant
	dLongitude *= constant
	dLatitude *= constant

	oLonSin := math.Sin(oLongitude)
	oLatSin := math.Sin(oLatitude)
	oLonCos := math.Cos(oLongitude)
	oLatCos := math.Cos(oLatitude)

	dLonSin := math.Sin(dLongitude)
	dLatSin := math.Sin(dLatitude)
	dLonCos := math.Cos(dLongitude)
	dLatCos := math.Cos(dLatitude)

	oArr := make([]float64, 0, 3)
	dArr := make([]float64, 0, 3)

	oArr = append(oArr, oLonCos*oLatCos)
	oArr = append(oArr, oLatCos*oLonSin)
	oArr = append(oArr, oLatSin)

	dArr = append(dArr, dLatCos*dLonCos)
	dArr = append(dArr, dLatCos*dLonSin)
	dArr = append(dArr, dLatSin)

	d := math.Sqrt((oArr[0]-dArr[0])*(oArr[0]-dArr[0]) +
		(oArr[1]-dArr[1])*(oArr[1]-dArr[1]) +
		(oArr[2]-dArr[2])*(oArr[2]-dArr[2]))

	return math.Asin(d/2.0) * 12742001.579854401
}
