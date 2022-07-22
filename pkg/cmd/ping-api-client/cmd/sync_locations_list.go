package cmd

import (
	"context"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/spf13/cobra"
)

type LocationsListOptions struct {
	Global *global.Options
	*printers.PrinterOptions
}

func NewLocationsListOptions(globalOpt *global.Options) *LocationsListOptions {
	return &LocationsListOptions{
		PrinterOptions: printers.NewPrinterOptions().WithDefaultTableWriter(),
		Global:         globalOpt,
	}
}

func NewCmdLocationsList(globalOpt *global.Options) *cobra.Command {
	o := NewLocationsListOptions(globalOpt)
	var cmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "get-all"},
		Short:   "list all data sync locations",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run()
		},
	}

	o.PrinterOptions.AddPrinterFlags(cmd.Flags())

	return cmd
}

// Complete the options
func (o *LocationsListOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *LocationsListOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *LocationsListOptions) Run() error {
	result, _, err := o.Global.Client.DataSync.LocationsGetAll(context.Background())
	if err != nil {
		return err
	}
	return printLocations(o.PrinterOptions, o.Global.Client.Config.BaseURL, result)
}
