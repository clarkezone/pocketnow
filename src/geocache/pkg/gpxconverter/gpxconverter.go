// gpxconverter.go
package gpxconverter

import (
	"encoding/xml"
	"time"
)

type Point interface {
	GetLat() float64
	GetLon() float64
	GetTimestamp() time.Time
}

type GPX struct {
	XMLName xml.Name `xml:"http://www.topografix.com/GPX/1/1 gpx"`
	Version string   `xml:"version,attr"`
	Creator string   `xml:"creator,attr"`
	Trk     Trk      `xml:"trk"`
}

type Trk struct {
	Name   string `xml:"name"`
	Type   string `xml:"type"`
	Trkseg Trkseg `xml:"trkseg"`
}

type Trkseg struct {
	Trkpts []Trkpt `xml:"trkpt"`
}

type Trkpt struct {
	Lat  float64   `xml:"lat,attr"`
	Lon  float64   `xml:"lon,attr"`
	Ele  float64   `xml:"ele"`
	Time time.Time `xml:"time"`
}

func ConvertToGPX(points []Point) (string, error) {
	gpx := GPX{
		Version: "1.1",
		Creator: "Your app name",
		Trk: Trk{
			Name: "new",
			Type: "Cycling",
		},
	}

	for _, point := range points {
		gpx.Trk.Trkseg.Trkpts = append(gpx.Trk.Trkseg.Trkpts, Trkpt{
			Lat:  point.GetLat(),
			Lon:  point.GetLon(),
			Ele:  0.0,
			Time: point.GetTimestamp(),
		})
	}

	gpxXML, err := xml.MarshalIndent(gpx, "", "  ")
	if err != nil {
		return "", err
	}

	gpxXML = []byte(xml.Header + string(gpxXML))
	return string(gpxXML), nil
}
