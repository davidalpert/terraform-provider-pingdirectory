package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/cmd/ping-api-client/cmd/global"
	"github.com/spf13/cobra"
	"strings"
)

type VersionOptions struct {
	Global *global.Options
	*printers.PrinterOptions
}

func NewVersionOptions(globalOpt *global.Options) *VersionOptions {
	return &VersionOptions{
		PrinterOptions: printers.NewPrinterOptions().WithDefaultOutput("text"),
		Global:         globalOpt,
	}
}

func NewCmdVersion(globalOpt *global.Options) *cobra.Command {
	o := NewVersionOptions(globalOpt)
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "show version information",
		Args:  cobra.NoArgs,
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
func (o *VersionOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *VersionOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

type VersionOutput struct {
	CLIName       string `json:"cli_name"`
	CLIVersion    string `json:"cli_version"`
	CLICommitHash string `json:"commit_hash"`
}

func (v VersionOutput) String() string {
	s := fmt.Sprintf("%s %s", v.CLIName, v.CLIVersion)
	if v.CLICommitHash != "" {
		s += "+" + v.CLICommitHash
	}
	return s
}

// Run the command
func (o *VersionOptions) Run() error {
	v := VersionOutput{
		CLIName:       o.Global.AppName,
		CLIVersion:    o.Global.Version,
		CLICommitHash: o.Global.CommitSHA,
	}
	if strings.EqualFold(*o.OutputFormat, "text") {
		_, err := fmt.Fprintf(o.Global.Streams.Out, "%s\n", v)
		return err
	}
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.OutputFormat = printers.StringPointer("json")
	}
	if v.CLICommitHash == "" {
		v.CLICommitHash = "n/a"
	}

	return o.Global.Streams.WriteOutput(v, o.PrinterOptions)
}
