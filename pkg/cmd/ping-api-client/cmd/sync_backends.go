package cmd

import (
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/spf13/cobra"
)

func NewCmdBackends(globalOpt *global.Options) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "backends",
		Short: "subcommands for data sync backends",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(NewCmdBackendsGet(globalOpt))

	return cmd
}
