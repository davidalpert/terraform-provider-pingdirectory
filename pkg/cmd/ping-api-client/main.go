/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
	"os"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	commit string = ""
)

func main() {
	client, err := apiclient.NewConfig(os.Getenv("PING_BASE_URL")).
		WithBasicAuth(os.Getenv("PING_USERNAME"), os.Getenv("PING_PASSWORD")).
		BuildPingApiClient()

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	err = cmd.NewRootCmd(printers.DefaultOSStreams(), "ping-api-client", version, commit, client).Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
