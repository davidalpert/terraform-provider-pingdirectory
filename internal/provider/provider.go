package provider

import (
	"context"
	"fmt"
	"github.com/davidalpert/terraform-provider-pingdirectory/pkg/ping/directory/apiclient/v1"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.Provider = &provider{}

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type provider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	//
	client *apiclient.Client

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	SyncBaseUrL types.String `tfsdk:"sync_base_url"`
	SyncDNUser  types.String `tfsdk:"sync_dn_user"`
	SyncDNPass  types.String `tfsdk:"sync_dn_pass"`
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Unmarshal provider schema config data
	var data providerData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the upstream provider SDK or HTTP client requires configuration, such
	// as authentication or logging, this is a great opportunity to do so.

	var syncBaseURL string
	if data.SyncBaseUrL.Null {
		syncBaseURL = os.Getenv("PING_SYNC_BASE_URL")
	} else {
		syncBaseURL = data.SyncBaseUrL.Value
	}
	if syncBaseURL == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find Sync BaseURL",
			"SyncBaseURL cannot be an empty string",
		)
		return
	}

	var syncDNUser string
	if data.SyncDNUser.Null {
		syncDNUser = os.Getenv("PING_SYNC_DN_USER")
	} else {
		syncDNUser = data.SyncDNUser.Value
	}
	if syncDNUser == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find Sync DN User",
			"SyncDNUser cannot be an empty string",
		)
		return
	}

	var syncDNPass string
	if data.SyncDNPass.Null {
		syncDNPass = os.Getenv("PING_SYNC_DN_PASS")
	} else {
		syncDNPass = data.SyncDNPass.Value
	}
	if syncDNPass == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find Sync DN Pass",
			"SyncDNPass cannot be an empty string",
		)
		return
	}

	c, err := apiclient.NewConfig(syncBaseURL).WithBasicAuth(syncDNUser, syncDNPass).BuildPingApiClient()
	if err != nil {
		resp.Diagnostics.AddError("Unable to build a Ping Directory API Client", err.Error())
	}
	p.client = c

	p.configured = true
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"pingdirectory_sync_location": syncLocationResourceType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"pingdirectory_sync_location": syncLocationDataSourceType{},
	}, nil
}

func (p *provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"sync_base_url": {
				MarkdownDescription: "DataSync BaseURL (e.g. https://datasync.my.org:8443)",
				Optional:            true,
				Type:                types.StringType,
			},
			"sync_dn_user": {
				MarkdownDescription: "DataSync auth dn",
				Optional:            true,
				Type:                types.StringType,
			},
			"sync_dn_pass": {
				MarkdownDescription: "DataSync auth pass",
				Optional:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &provider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func convertProviderType(in tfsdk.Provider) (provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*provider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return provider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return provider{}, diags
	}

	return *p, diags
}
