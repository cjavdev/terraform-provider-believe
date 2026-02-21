// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package coaching_principle

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/option"
	"github.com/cjavdev/terraform-provider-believe/internal/apijson"
	"github.com/cjavdev/terraform-provider-believe/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type CoachingPrincipleDataSource struct {
	client *believe.Client
}

var _ datasource.DataSourceWithConfigure = (*CoachingPrincipleDataSource)(nil)

func NewCoachingPrincipleDataSource() datasource.DataSource {
	return &CoachingPrincipleDataSource{}
}

func (d *CoachingPrincipleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_coaching_principle"
}

func (d *CoachingPrincipleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *CoachingPrincipleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CoachingPrincipleDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := d.client.Coaching.Principles.Get(
		ctx,
		data.PrincipleID.ValueString(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
