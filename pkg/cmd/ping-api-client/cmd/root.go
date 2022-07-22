package cmd

import (
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
	"github.com/spf13/cobra"
)

func NewRootCmd(str printers.IOStreams, appName, version, commit string, client *apiclient.Client) *cobra.Command {
	o := &global.Options{
		AppName:   appName,
		Version:   version,
		CommitSHA: commit,
		Streams:   str,
		Client:    client,
	}

	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "ping-api-client",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//RunE: func(cmd *cobra.Command, args []string) error {
		//	if err := o.Complete(cmd, args); err != nil {
		//		return err
		//	}
		//	if err := o.Validate(); err != nil {
		//		return err
		//	}
		//	return o.Run()
		//},
	}

	rootCmd.AddCommand(NewCmdSync(o))
	rootCmd.AddCommand(NewCmdVersion(o))

	return rootCmd
}
