package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var stderr = os.Stderr

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct{}

func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"dummy": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {

}

// GetResources - Defines provider resources
func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{}, nil
}

// GetDataSources - Defines provider data sources
func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"paths_components": dataSourceComponentsType{},
		"paths_parents":    dataSourceParentsType{},
	}, nil
}
