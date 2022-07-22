package cmd

import (
	"context"
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
	"github.com/spf13/cobra"
)

type LocationsReadOptions struct {
	Global *global.Options
	*printers.PrinterOptions
	Name string
}

func NewLocationsReadOptions(globalOpt *global.Options) *LocationsReadOptions {
	return &LocationsReadOptions{
		PrinterOptions: printers.NewPrinterOptions().WithDefaultTableWriter(),
		Global:         globalOpt,
	}
}

func NewCmdLocationsRead(globalOpt *global.Options) *cobra.Command {
	o := NewLocationsReadOptions(globalOpt)
	var cmd = &cobra.Command{
		Use:     "get <name>",
		Aliases: []string{"g", "read", "describe"},
		Short:   "get information about one data sync location",
		Args:    cobra.ExactArgs(1),
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
func (o *LocationsReadOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Name = args[0]
	return nil
}

// Validate the options
func (o *LocationsReadOptions) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("location name is required")
	}
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *LocationsReadOptions) Run() error {
	result, _, err := o.Global.Client.DataSync.LocationsGet(context.Background(), o.Name)
	if err != nil {
		return err
	}
	return printLocations(o.PrinterOptions, o.Global.Client.Config.BaseURL, []apiclient.Location{*result})
}
