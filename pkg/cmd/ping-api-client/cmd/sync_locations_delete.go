package cmd

import (
	"context"
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/spf13/cobra"
)

type LocationsDeleteOptions struct {
	Global *global.Options
	*printers.PrinterOptions
	Name string
}

func NewLocationsDeleteOptions(globalOpt *global.Options) *LocationsDeleteOptions {
	return &LocationsDeleteOptions{
		PrinterOptions: printers.NewPrinterOptions().WithDefaultOutput("text"),
		Global:         globalOpt,
	}
}

func NewCmdLocationsDelete(globalOpt *global.Options) *cobra.Command {
	o := NewLocationsDeleteOptions(globalOpt)
	var cmd = &cobra.Command{
		Use:     "delete <name>",
		Aliases: []string{"d"},
		Short:   "delete a data sync location",
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
func (o *LocationsDeleteOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Name = args[0]
	return nil
}

// Validate the options
func (o *LocationsDeleteOptions) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("location name is required")
	}
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *LocationsDeleteOptions) Run() error {
	_, err := o.Global.Client.DataSync.LocationDeleteByName(context.Background(), o.Name)
	return err
}
