package service

import "math"

const (
	earthRadius = 6371 // km
)

func Haversine(lon1, lat1, lon2, lat2 float64) float64 {
	lat1 = lat1 * math.Pi / 180
	lon1 = lon1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180
	lon2 = lon2 * math.Pi / 180

	// Calculate the differences
	dlat := lat2 - lat1
	dlon := lon2 - lon1

	// Calculate the haversine formula
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}
