package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/davidalpert/terraform-provider-pingdirectory/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	commit string = ""
)

func main() {
	var debug bool
	var showVersion bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.BoolVar(&showVersion, "version", false, "show the plugin version")
	flag.Parse()

	if showVersion {
		v := version
		if commit != "" {
			v += "+" + commit
		}
		fmt.Printf("terraform-provider-pingdirectory %s\n", v)
		os.Exit(0)
	}

	opts := providerserver.ServeOpts{
		// TODO: Update this string with the published name of your provider.
		Address: "registry.terraform.io/davidalpert/pingdirectory",
		Debug: debug,
	}

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
