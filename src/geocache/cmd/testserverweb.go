// Package cmd contains the cli command definitions for geocache:w
package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"fmt"
	"net/http"

	"github.com/clarkezone/boosted-go/basicserverhttp"
	clarkezoneLog "github.com/clarkezone/boosted-go/log"
	"github.com/clarkezone/boosted-go/middlewarehttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/clarkezone/geocache/internal"
	"github.com/clarkezone/geocache/pkg/config"
)

var bsweb = basicserverhttp.CreateBasicServer()

var (
	// testserverWebCmd represents the testserver command
	testserverWebCmd = &cobra.Command{
		Use:   "testserverweb",
		Short: "Starts a test http server to test logging and metrics",
		Long: `Starts a listener that will
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("geocache version %v,%v started in testserver mode\n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)
			mux := basicserverhttp.DefaultMux()
			mux.HandleFunc("/", getHelloHandler())

			var wrappedmux http.Handler
			wrappedmux = middlewarehttp.NewLoggingMiddleware(mux)
			wrappedmux = middlewarehttp.NewPromMetricsMiddlewareWeb("geocache_testWebservice", wrappedmux)

			if viper.GetString(internal.ServiceURLVar) != "" {
				clarkezoneLog.Successf("Delegating to %v", internal.ServiceURL)
			} else {
				clarkezoneLog.Debugf("Not delegating to %v", internal.ServiceURL)
			}

			clarkezoneLog.Successf("Starting web server on port %v", internal.Port)
			bsweb.StartMetrics()
			clarkezoneLog.Successf("Starting metrics on port %v", internal.MetricsPort)
			bsweb.StartListen("", wrappedmux)
			return bsweb.WaitforInterupt()
		},
	}
)

func getHelloHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var message string

		message = fmt.Sprintln("Hello World<br>")

		_, err := w.Write([]byte(message))
		if err != nil {
			clarkezoneLog.Debugf("Failed to write bytes %v\n", err)
			panic(err)
		}

	}
}

func init() {
	rootCmd.AddCommand(testserverWebCmd)
	testserverWebCmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, testserverWebCmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		panic(err)
	}
}
