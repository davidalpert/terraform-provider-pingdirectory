package cmd

import (
	"context"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/spf13/cobra"
)

type BackendsGetOptions struct {
	Global *global.Options
	*printers.PrinterOptions
}

func NewBackendsGetOptions(globalOpt *global.Options) *BackendsGetOptions {
	return &BackendsGetOptions{
		PrinterOptions: printers.NewPrinterOptions().WithDefaultOutput("text"),
		Global:         globalOpt,
	}
}

func NewCmdBackendsGet(globalOpt *global.Options) *cobra.Command {
	o := NewBackendsGetOptions(globalOpt)
	var cmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "get information about data sync backends",
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
func (o *BackendsGetOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *BackendsGetOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *BackendsGetOptions) Run() error {
	result := o.Global.Client.DataSync.GetBackendConfig(context.Background())
	return o.Global.Streams.WriteOutput(result, o.PrinterOptions)
}
