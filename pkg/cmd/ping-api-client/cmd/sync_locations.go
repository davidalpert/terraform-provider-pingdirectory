package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"strings"
)

func NewCmdLocations(globalOpt *global.Options) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "locations",
		Aliases: []string{"l", "loc", "locations"},
		Short:   "subcommands for data sync locations",
		Args:    cobra.NoArgs,
	}

	cmd.AddCommand(NewCmdLocationsCreate(globalOpt))
	cmd.AddCommand(NewCmdLocationsRead(globalOpt))
	cmd.AddCommand(NewCmdLocationsList(globalOpt))
	cmd.AddCommand(NewCmdLocationsUpdate(globalOpt))
	cmd.AddCommand(NewCmdLocationsDelete(globalOpt))

	return cmd
}

func printLocations(o *printers.PrinterOptions, baseUrl string, result []apiclient.Location) error {
	caption := fmt.Sprintf("locations @ %s", baseUrl)
	return o.WithTableWriter(caption, func(t *tablewriter.Table) {
		t.SetHeader([]string{"Name", "Description", "Preferred Failover Locations"})
		for _, r := range result {
			t.Append([]string{r.Name, r.Description, strings.Join(r.PreferredFailoverLocation, ",")})
		}
	}).WriteOutput(result)
}
