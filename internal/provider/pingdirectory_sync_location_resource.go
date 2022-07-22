package provider

import (
	"context"
	"fmt"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.ResourceType = syncLocationResourceType{}
var _ tfsdk.Resource = syncLocationResource{}
var _ tfsdk.ResourceWithImportState = syncLocationResource{}

type syncLocationResourceType struct{}

func (t syncLocationResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Data Sync Location resource",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Computed:            true,
				MarkdownDescription: "Sync Location identifier",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"name": {
				Type:                types.StringType,
				MarkdownDescription: "Sync Location name",
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					// These are example validators from terraform-plugin-framework-validators
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-z0-9-]+$`),
						"must contain only lowercase characters",
					),
				},
			},
			"description": {
				MarkdownDescription: "Sync Location description",
				Optional:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func (t syncLocationResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return syncLocationResource{
		provider: provider,
	}, diags
}

type syncLocationResourceData struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

type syncLocationResource struct {
	provider provider
}

func (r syncLocationResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var data syncLocationResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	l := apiclient.Location{Name: data.Name.Value}
	if !data.Description.IsNull() {
		l.Description = data.Description.Value
	}

	tflog.Debug(ctx, "Create", map[string]interface{}{"resource": "Location", "name": data.Name.Value, "value": l})

	_, err := r.provider.client.DataSync.LocationsCreate(ctx, l)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Sync Location: %v", err))
		return
	}

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ID = types.String{Value: strings.ToLower(l.Name)}

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Debug(ctx, "created a location", map[string]interface{}{"id": data.ID, "name": data.Name, "description": data.Description})

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r syncLocationResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data syncLocationResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read", map[string]interface{}{"resource": "Location", "id": data.ID.Value})

	l, _, err := r.provider.client.DataSync.LocationsGet(ctx, strings.ToLower(data.ID.Value))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Sync Location, got error: %v", err))
		return
	}

	data.Name = types.String{Value: strings.ToLower(l.Name)}
	data.Description = types.String{Value: l.Description}

	tflog.Debug(ctx, "Read Result", map[string]interface{}{"resource": "Location", "id": data.ID.Value, "name": data.Name.Value, "description": data.Description.Value})

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r syncLocationResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data syncLocationResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	l := apiclient.Location{Name: data.Name.Value}
	if !data.Description.IsNull() {
		l.Description = data.Description.Value
	}

	tflog.Debug(ctx, "Update", map[string]interface{}{"resource": "Location", "id": data.ID.Value, "value": l})

	_, err := r.provider.client.DataSync.LocationUpdate(ctx, l)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Sync Location, got error: %v", err))
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r syncLocationResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data syncLocationResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Delete", map[string]interface{}{"resource": "Location", "id": data.ID.Value})

	_, err := r.provider.client.DataSync.LocationDeleteByName(ctx, data.Name.Value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Sync Location, got error: %v", err))
		return
	}
}

func (r syncLocationResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
