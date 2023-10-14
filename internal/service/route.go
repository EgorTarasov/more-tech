package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"more-tech/internal/logging"
	"net/http"
	"time"
)

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

type RouteResponse struct {
	Type     string     `json:"type"`
	Metadata Metadata   `json:"metadata"`
	Bbox     []float64  `json:"bbox"`
	Features []Features `json:"features"`
}

type Query struct {
	Coordinates any    `json:"coordinates"`
	Profile     string `json:"profile"`
	Format      string `json:"format"`
}

type Engine struct {
	Version   string    `json:"version"`
	BuildDate time.Time `json:"build_date"`
	GraphDate time.Time `json:"graph_date"`
}

type Metadata struct {
	Attribution string `json:"attribution"`
	Service     string `json:"service"`
	Timestamp   int64  `json:"timestamp"`
	Query       Query  `json:"query"`
	Engine      Engine `json:"engine"`
}

type Steps struct {
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
	Type        int     `json:"type"`
	Instruction string  `json:"instruction"`
	Name        string  `json:"name"`
	WayPoints   []int   `json:"way_points"`
}

type Segments struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Steps    []Steps `json:"steps"`
}

type Summary struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
}

type Properties struct {
	Transfers int        `json:"transfers"`
	Fare      int        `json:"fare"`
	Segments  []Segments `json:"segments"`
	Summary   Summary    `json:"summary"`
	WayPoints []int      `json:"way_points"`
}

type Geometry struct {
	Coordinates any    `json:"coordinates"`
	Type        string `json:"type"`
}

type Features struct {
	Bbox       []float64  `json:"bbox"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

func GetEstimatedTime(startLongitude, startLatitude, endLongitude, endLatitude float64, routeHost string) (timeCar, timeWalk float64, err error) {
	var carResponse, walkResponse RouteResponse
	logging.Log.Debugf("sending car request to %s", fmt.Sprintf("http://%s/ors/v2/directions/driving-car?start=%f,%f&end=%f,%f", routeHost, startLongitude, startLatitude, endLongitude, endLatitude))

	resp, err := http.Get(fmt.Sprintf("http://%s/ors/v2/directions/driving-car?start=%f,%f&end=%f,%f", routeHost, startLongitude, startLatitude, endLongitude, endLatitude))
	if err != nil {
		return 0, 0, err
	} else if resp.StatusCode != http.StatusOK {
		return 0, 0, errors.New("route service error")
	}

	if err := json.NewDecoder(resp.Body).Decode(&carResponse); err != nil {
		return 0, 0, err
	}

	logging.Log.Debugf("sending walk request to %s", fmt.Sprintf("http://%s/ors/v2/directions/foot-walking?start=%f,%f&end=%f,%f", routeHost, startLongitude, startLatitude, endLongitude, endLatitude))

	resp, err = http.Get(fmt.Sprintf("http://%s/ors/v2/directions/foot-walking?start=%f,%f&end=%f,%f", routeHost, startLongitude, startLatitude, endLongitude, endLatitude))
	if err != nil {
		return 0, 0, err
	} else if resp.StatusCode != http.StatusOK {
		return 0, 0, errors.New("route service error")
	}

	if err := json.NewDecoder(resp.Body).Decode(&walkResponse); err != nil {
		return 0, 0, err
	}

	return carResponse.Features[0].Properties.Segments[0].Duration, walkResponse.Features[0].Properties.Segments[0].Duration, nil
}
