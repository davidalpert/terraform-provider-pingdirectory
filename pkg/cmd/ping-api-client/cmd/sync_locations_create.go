package cmd

import (
	"context"
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
	"github.com/spf13/cobra"
)

type LocationsCreateOptions struct {
	Global *global.Options
	*printers.PrinterOptions
	Location apiclient.Location
}

func NewLocationsCreateOptions(gOpt *global.Options) *LocationsCreateOptions {
	return &LocationsCreateOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(gOpt.Streams).WithDefaultOutput("text"),
		Global:         gOpt,
	}
}

func NewCmdLocationsCreate(globalOpt *global.Options) *cobra.Command {
	o := NewLocationsCreateOptions(globalOpt)
	var cmd = &cobra.Command{
		Use:     "create <name>",
		Aliases: []string{"c", "new"},
		Short:   "create a data sync location",
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
	cmd.Flags().StringVarP(&o.Location.Description, "description", "d", "", "description")
	cmd.Flags().StringSliceVarP(&o.Location.PreferredFailoverLocation, "preferred-failover-location", "f", []string{}, "preferred failover locations")

	return cmd
}

// Complete the options
func (o *LocationsCreateOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Location.Name = args[0]
	return nil
}

// Validate the options
func (o *LocationsCreateOptions) Validate() error {
	if o.Location.Name == "" {
		return fmt.Errorf("location name is required")
	}
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *LocationsCreateOptions) Run() error {
	_, err := o.Global.Client.DataSync.LocationsCreate(context.Background(), o.Location)
	return err
}
