package cmd

import (
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/spf13/cobra"
)

func NewCmdSync(globalOpt *global.Options) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "data-sync",
		Aliases: []string{"ds", "sync"},
		Short:   "subcommands for Ping Data Sync",
		Args:    cobra.NoArgs,
	}

	cmd.AddCommand(NewCmdBackends(globalOpt))
	cmd.AddCommand(NewCmdLocations(globalOpt))

	return cmd
}
