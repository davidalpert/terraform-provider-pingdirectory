package global

import (
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
)

type Options struct {
	AppName   string
	CommitSHA string
	Streams   printers.IOStreams
	Version   string
	Client    *apiclient.Client
}
