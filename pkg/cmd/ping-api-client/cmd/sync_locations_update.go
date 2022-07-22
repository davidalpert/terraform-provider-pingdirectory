package cmd

import (
	"context"
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
	"github.com/spf13/cobra"
)

type LocationsUpdateOptions struct {
	Global *global.Options
	*printers.PrinterOptions
	Location apiclient.Location
}

func NewLocationsUpdateOptions(globalOpt *global.Options) *LocationsUpdateOptions {
	return &LocationsUpdateOptions{
		PrinterOptions: printers.NewPrinterOptions().WithDefaultOutput("text"),
		Global:         globalOpt,
	}
}

func NewCmdLocationsUpdate(globalOpt *global.Options) *cobra.Command {
	o := NewLocationsUpdateOptions(globalOpt)
	var cmd = &cobra.Command{
		Use:     "update <name>",
		Aliases: []string{"u"},
		Short:   "update a data sync location",
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
func (o *LocationsUpdateOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Location.Name = args[0]
	return nil
}

// Validate the options
func (o *LocationsUpdateOptions) Validate() error {
	if o.Location.Name == "" {
		return fmt.Errorf("location name is required")
	}
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *LocationsUpdateOptions) Run() error {
	_, err := o.Global.Client.DataSync.LocationUpdate(context.Background(), o.Location)
	return err
}
