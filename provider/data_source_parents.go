package provider

import (
	"context"
	"path"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type dataSourceParentsType struct {
	Path    string   `tfsdk:"path"`
	Parents []string `tfsdk:"parents"`
}

func (r dataSourceParentsType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"path": {
				Description: "The path which to return parents for",
				Required:    true,
				Type:        types.StringType,
			},
			"parents": {
				Description: "Parents of the path",
				Computed:    true,
				Type:        types.SetType{ElemType: types.StringType},
			},
		},
	}, nil
}

func (r dataSourceParentsType) NewDataSource(ctx context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceParents{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceParents struct {
	p provider
}

func (r dataSourceParents) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	resourceState := dataSourceParentsType{}
	diags := req.Config.Get(ctx, &resourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	remainingPath := strings.TrimSuffix(resourceState.Path, string(filepath.Separator))
	var lastPath string
	for {
		lastPath = remainingPath
		remainingPath, _ = path.Split(remainingPath)
		remainingPath = strings.TrimSuffix(remainingPath, string(filepath.Separator))
		if remainingPath == "" {
			break
		}
		if lastPath == remainingPath {
			resp.Diagnostics.AddError("Bug detected", "While traversing the path, an endless loop was caused")
			return
		}
		resourceState.Parents = append(resourceState.Parents, remainingPath)
	}

	diags = resp.State.Set(ctx, &resourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
