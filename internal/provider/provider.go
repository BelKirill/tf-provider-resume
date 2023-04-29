package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure resumeProvider satisfies various provider interfaces.
var _ provider.Provider = &resumeProvider{}

// resumeProvider defines the provider implementation.
type resumeProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// resumeProviderModel describes the provider data model.
type resumeProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *resumeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "resume"
	resp.Version = p.version
}

func (p *resumeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Endpoint for the resume builder service",
				Optional:            true,
			},
			// Need to secure the end point
		},
	}
}

func (p *resumeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	endpoint := "127.0.0.1:8000"

	var data resumeProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if data.Endpoint.ValueString() != "" {
		endpoint = data.Endpoint.ValueString()
	}

	if endpoint == "" {
		resp.Diagnostics.AddError(
			"Missing endpoint data",
			"No Endpoint for the resume builder API was found.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *resumeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewIntroResource,
	}
}

func (p *resumeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewIntroDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &resumeProvider{
			version: version,
		}
	}
}
