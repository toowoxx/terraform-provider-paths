package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/toowoxx/go-lib-fs/filepaths"
)

type dataSourceComponentsType struct {
	Path       string   `tfsdk:"path"`
	Components []string `tfsdk:"components"`
}

func (r dataSourceComponentsType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"path": {
				Description: "The path which to return components for",
				Required:    true,
				Type:        types.StringType,
			},
			"components": {
				Description: "Components of the path",
				Computed:    true,
				Type:        types.SetType{ElemType: types.StringType},
			},
		},
	}, nil
}

func (r dataSourceComponentsType) NewDataSource(ctx context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceComponents{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceComponents struct {
	p provider
}

func (r dataSourceComponents) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	resourceState := dataSourceComponentsType{}
	diags := req.Config.Get(ctx, &resourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceState.Components = filepaths.Components(resourceState.Path)

	diags = resp.State.Set(ctx, &resourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
