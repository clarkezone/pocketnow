// Package cmd contains the cli command definitions for geocache:w
package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/clarkezone/geocache/internal"
	"github.com/clarkezone/geocache/pkg/basicserver"
	"github.com/clarkezone/geocache/pkg/config"
	"github.com/clarkezone/geocache/pkg/geocacheservice"
	clarkezoneLog "github.com/clarkezone/geocache/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var geoweb = basicserver.CreateBasicServer()

var (
	// testserverWebCmd represents the testserver command
	geoserverWebCmd = &cobra.Command{
		Use:   "geoserverweb",
		Short: "Starts a geoserver endpoint to receive data from overland",
		Long: `Starts a listener that will
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("geocache version %v,%v started in geoserverweb mode\n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)
			mux := basicserver.DefaultMux()
			mux.HandleFunc("/postgeo", getGeoHandler())

			var wrappedmux http.Handler
			wrappedmux = basicserver.NewLoggingMiddleware(mux)
			wrappedmux = basicserver.NewPromMetricsMiddlewareWeb("geocache_geoWebservice", wrappedmux)

			if viper.GetString(internal.ServiceURLVar) != "" {
				clarkezoneLog.Successf("Delegating to %v", internal.ServiceURL)
			} else {
				clarkezoneLog.Debugf("Not delegating to %v", internal.ServiceURL)
			}

			clarkezoneLog.Successf("Starting web server on port %v", internal.Port)
			geoweb.StartMetrics()
			clarkezoneLog.Successf("Starting metrics on port %v", internal.MetricsPort)
			geoweb.StartListen("", wrappedmux)
			return geoweb.WaitforInterupt()
		},
	}
)

type GeoStruct struct {
	Locations []struct {
		Type     string `json:"type"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			Motion             []string  `json:"motion"`
			Speed              int       `json:"speed"`
			BatteryLevel       float64   `json:"battery_level"`
			Altitude           int       `json:"altitude"`
			BatteryState       string    `json:"battery_state"`
			HorizontalAccuracy int       `json:"horizontal_accuracy"`
			VerticalAccuracy   int       `json:"vertical_accuracy"`
			Timestamp          time.Time `json:"timestamp"`
			Wifi               string    `json:"wifi"`
		} `json:"properties"`
	} `json:"locations"`
}

func getGeoHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		dresp := GeoStruct{}
		err := json.NewDecoder(r.Body).Decode(&dresp)

		if err != nil {
			clarkezoneLog.Errorf("unable to unmarshal json %v", err)
		}

		clarkezoneLog.Debugf("Got a geocoordinate %v", dresp.Locations[0].Geometry.Coordinates[0])

		if viper.GetString(internal.ServiceURLVar) != "" {
			newFunction(w)
		} else {
			clarkezoneLog.Debugf("Envalid ServiceURL, unable to write data %v", internal.ServiceURL)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("500 - Something bad happened!"))
			if err != nil {
				clarkezoneLog.Errorf("Error unable to write to connection %v", err)
			}
		}

		//ctx.SuccessString("text/json",
		w.Header().Add("Content-Type", "application/json")
		message := `
{
  "result": "ok"
}`

		_, err = w.Write([]byte(message))
		if err != nil {
			clarkezoneLog.Debugf("Failed to write response %v\n", err)
		}

	}
}

func newFunction(w http.ResponseWriter) {
	conn, err := grpc.Dial(internal.ServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		clarkezoneLog.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()

	if err == nil {
		client := geocacheservice.NewGeoCacheServiceClient(conn)
		locs := &geocacheservice.Locations{}
		for i := range dresp.Locations {
			clarkezoneLog.Debugf("Got a geocoordinate %v", i.Geometry.Coordinates[0])
			locs.Locations = append(locs.Locations, &geocacheservice.Location{
				Latitude: dresp.Locations[i].Geometry.Coordinates[0]})
		}
		_, err := client.SaveLocations(context.Background())
		if err != nil {
			clarkezoneLog.Errorf("Error %v", err)
		} else {
			clarkezoneLog.Successf("Result received")
		}
	} else {
		clarkezoneLog.Errorf("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("500 - Something bad happened!"))
		if err != nil {
			clarkezoneLog.Errorf("Error %v", err)
		}
	}
}

func init() {
	rootCmd.AddCommand(geoserverWebCmd)
	geoserverWebCmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, geoserverWebCmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		panic(err)
	}
}
