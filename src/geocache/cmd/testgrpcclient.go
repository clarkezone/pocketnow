// Package cmd contains the cli command definitions for geocache:w
package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"context"

	"github.com/clarkezone/geocache/internal"
	"github.com/clarkezone/geocache/pkg/config"
	"github.com/clarkezone/geocache/pkg/geocacheservice"
	clarkezoneLog "github.com/clarkezone/geocache/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TestClientGrpcCmd is the command to start a test grpc client
type TestClientGrpcCmd struct {
}

func newTestClientGrpcCmd(partent *cobra.Command) (*TestClientGrpcCmd, error) {
	cmd := &cobra.Command{
		Use:   "testclientgrpc",
		Short: "Starts a client",
		Long: `Starts a client that will call the testservergrpc which must be running
have already been started with the testservergrpc command`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("geocache version %v,%v started in testclientgrpc mode\n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)

			clarkezoneLog.Successf("ServiceURL %v", internal.ServiceURL)

			conn, err := grpc.Dial(internal.ServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				clarkezoneLog.Errorf("fail to dial: %v", err)
			}
			defer conn.Close()

			if err == nil {
				client := geocacheservice.NewGeoCacheServiceClient(conn)
				_, err := client.SaveLocations(context.Background(), &geocacheservice.Locations{})
				if err != nil {
					clarkezoneLog.Errorf("Error %v", err)
				} else {
					clarkezoneLog.Successf("Result received")
				}
			}

			return err
		},
	}
	partent.AddCommand(cmd)
	cmd.PersistentFlags().StringVarP(&internal.ServiceURL, internal.ServiceURLVar, "",
		viper.GetString(internal.ServiceURLVar), "If value passed, testserverweb will delegate to this service")
	err := viper.BindPFlag(internal.ServiceURLVar, cmd.PersistentFlags().Lookup(internal.ServiceURLVar))
	if err != nil {
		panic(err)
	}
	return nil, nil
}
