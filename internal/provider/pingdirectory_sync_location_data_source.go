package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"regexp"
	"strings"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.DataSourceType = syncLocationDataSourceType{}
var _ tfsdk.DataSource = syncLocationDataSource{}

type syncLocationDataSourceType struct{}

func (t syncLocationDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example data source",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Required:            true,
				Type:                types.StringType,
				MarkdownDescription: "Sync Location uri",
				Validators: []tfsdk.AttributeValidator{
					// These are example validators from terraform-plugin-framework-validators
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-z0-9-]+$`),
						"must contain only lowercase characters",
					),
				},
			},
			"name": {
				Computed:            true,
				MarkdownDescription: "Sync Location identifier",
				Type:                types.StringType,
			},
			"description": {
				MarkdownDescription: "Sync Location description",
				Computed:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func (t syncLocationDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return syncLocationDataSource{
		provider: provider,
	}, diags
}

type exampleDataSourceData struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

type syncLocationDataSource struct {
	provider provider
}

func (d syncLocationDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data exampleDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read Data Source", map[string]interface{}{"resource": "Location", "id": data.ID.Value})

	l, _, err := d.provider.client.DataSync.LocationsGet(ctx, data.ID.Value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Sync Location, got error: %v", err))
		return
	}

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ID = types.String{Value: strings.ToLower(l.Name)}
	data.Name = types.String{Value: strings.ToLower(l.Name)}
	data.Description = types.String{Value: l.Description}

	tflog.Debug(ctx, "Read Data Source Complete", map[string]interface{}{"resource": "Location", "id": data.ID.Value, "name": data.Name.Value, "description": data.Description.Value})

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
