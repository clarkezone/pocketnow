// gpxconverter_test.go
package gpxconverter

import (
	"encoding/xml"
	"testing"
	"time"
)

type PointStruct struct {
	ID        string    `json:"id"`
	Lat       float64   `json:"Lat"`
	Lon       float64   `json:"Lon"`
	Timestamp time.Time `json:"Timestamp"`
}

func (p PointStruct) GetLat() float64 {
	return p.Lat
}

func (p PointStruct) GetLon() float64 {
	return p.Lon
}

func (p PointStruct) GetTimestamp() time.Time {
	return p.Timestamp
}

func TestConvertToGPX(t *testing.T) {
	points := []Point{
		PointStruct{
			ID:        "a48eff9f-ca51-42e5-9f46-8b6d152aa950",
			Lat:       37.4219999,
			Lon:       -122.0840575,
			Timestamp: time.Now(),
		},
	}

	gpxStr, err := ConvertToGPX(points)
	if err != nil {
		t.Error("Failed to convert to GPX:", err)
	}

	if gpxStr == "" {
		t.Error("Got empty GPX string")
	}

	// Now parse the XML back into a GPX object
	var gpx GPX
	err = xml.Unmarshal([]byte(gpxStr), &gpx)
	if err != nil {
		t.Error("Failed to parse GPX:", err)
	}

	// Check if the parsed GPX has the correct data
	if len(gpx.Trk.Trkseg.Trkpts) != 1 {
		t.Errorf("Expected 1 track point, got %d", len(gpx.Trk.Trkseg.Trkpts))
	}

	if gpx.Trk.Trkseg.Trkpts[0].Lat != points[0].GetLat() {
		t.Errorf("Expected lat %f, got %f", points[0].GetLat(), gpx.Trk.Trkseg.Trkpts[0].Lat)
	}

	if gpx.Trk.Trkseg.Trkpts[0].Lon != points[0].GetLon() {
		t.Errorf("Expected lon %f, got %f", points[0].GetLon(), gpx.Trk.Trkseg.Trkpts[0].Lon)
	}
}
