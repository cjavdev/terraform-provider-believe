// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package coaching_principle

import (
	"context"
	"fmt"

	"github.com/cjavdev/believe-go"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stainless-sdks/believe-terraform/internal/apijson"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type CoachingPrinciplesDataSource struct {
	client *believe.Client
}

var _ datasource.DataSourceWithConfigure = (*CoachingPrinciplesDataSource)(nil)

func NewCoachingPrinciplesDataSource() datasource.DataSource {
	return &CoachingPrinciplesDataSource{}
}

func (d *CoachingPrinciplesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_coaching_principles"
}

func (d *CoachingPrinciplesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*believe.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *believe.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *CoachingPrinciplesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CoachingPrinciplesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toListParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	env := CoachingPrinciplesDataListDataSourceEnvelope{}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []attr.Value{}
	if maxItems <= 0 {
		maxItems = 1000
	}
	page, err := d.client.Coaching.Principles.List(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Data) > 0 {
		bytes := []byte(page.RawJSON())
		err = apijson.UnmarshalComputed(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, env.Data.Elements()...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

	acc = acc[:min(len(acc), maxItems)]
	result, diags := customfield.NewObjectListFromAttributes[CoachingPrinciplesItemsDataSourceModel](ctx, acc)
	resp.Diagnostics.Append(diags...)
	data.Items = result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
