// Package cmd contains the cli command definitions for pocketshorten
package cmd

/*
Copyright © 2022 James Clarke james@clarkezone.net

*/

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/clarkezone/boosted-go/basicserverhttp"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"

	"github.com/clarkezone/geocache/internal"
	"github.com/clarkezone/geocache/pkg/config"
	"github.com/clarkezone/geocache/pkg/cosmosdbbackend"
	"github.com/clarkezone/geocache/pkg/gpxconverter"

	"tailscale.com/tsnet"
	"tailscale.com/types/logger"
)

const (
	prefix string = "pocketshorten_frontend"
)

// QueryServerFrontendCmdState object
type QueryServerFrontendCmdState struct {
	webserver *basicserverhttp.BasicServer
	// shortener *shortener.ShortenHandler
}

func newQueryServerFrontend(parent *cobra.Command) (*QueryServerFrontendCmdState, error) {
	ss := basicserverhttp.CreateBasicServer()
	cmdstate := &QueryServerFrontendCmdState{webserver: ss}

	// shortenservercmd represents the testserver command
	shortenservercmd := &cobra.Command{
		Use:   "servequeryfrontend",
		Short: "Starts a frontend query server",
		Long: `Starts a frontend query server that will
listen for queries and return results in different formats:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("Geo frontend query server started %v,%v \n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)

			//			ss.StartListen("", wrappedmux)
			//		return ss.WaitforInterupt()

			hostname := "queryserver"

			s := &tsnet.Server{
				Hostname: hostname,
				Logf:     logger.Discard,
			}

			defer s.Close()

			ln, err := s.Listen("tcp", ":80")
			if err != nil {
				log.Fatal(err)
			}

			defer ln.Close()
			lc, err := s.LocalClient()
			if err != nil {
				log.Fatal(err)
			}

			log.Fatal(http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				who, err := lc.WhoIs(r.Context(), r.RemoteAddr)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}

				fmt.Fprintf(w, "<html><body><h1>Hello, world!</h1>\n")
				fmt.Fprintf(w, "<p>You are <b>%s</b> from <b>%s</b> (%s)</p>",
					html.EscapeString(who.UserProfile.LoginName),
					html.EscapeString(firstLabel(who.Node.ComputedName)),
					r.RemoteAddr)

				endpoint := os.Getenv("COSMOSDB_URL")
				if endpoint == "" {
					log.Fatal("COSMOSDB_ENDPOINT could not be found")
				}
				key := os.Getenv("COSMOSDB_KEY")
				if key == "" {
					log.Fatal("COSMOSDB_KEY could not be found")
				}

				cosmosdal, err := cosmosdbbackend.CreateCosmosDAL(endpoint, key)
				if err != nil {
					log.Fatal(err)
				}

				var databaseName = "pocketnow"
				var containerName = "geocache"
				var partitionKey = "1"

				sql := "SELECT * FROM geocache c where c.Timestamp >= '2023-05-13T14:45:59Z' AND c.Timestamp <= '2023-05-13T15:15:59Z'"
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

				returned, err := cosmosdal.Query(databaseName, containerName, partitionKey, sql, ctx)

				if err != nil {
					log.Fatalf("createItem failed: %s\n", err)
				}

				log.Printf("returned: %+v\n", len(returned))

				b := make([]gpxconverter.Point, len(returned))
				for i := range returned {
					b[i] = returned[i]
				}
				fmt.Fprintf(w, "<textarea>")
				err = gpxconverter.ConvertToGPXWriter(b, w)
				if err != nil {
					log.Fatalf("createItem failed: %s\n", err)
				}
				fmt.Fprintf(w, "</textarea>")
			})))

			return nil
		},
	}
	err := cmdstate.configFlags(shortenservercmd)
	if err != nil {
		return nil, err
	}
	parent.AddCommand(shortenservercmd)
	return cmdstate, nil
}

func firstLabel(s string) string {
	if hostname, _, ok := strings.Cut(s, "."); ok {
		return hostname
	}

	return s
}

func (state *QueryServerFrontendCmdState) configFlags(cmd *cobra.Command) error {
	cmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, cmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		return err
	}
	return nil
}
