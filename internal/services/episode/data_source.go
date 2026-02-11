// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-go/option"
	"github.com/stainless-sdks/believe-terraform/internal/apijson"
	"github.com/stainless-sdks/believe-terraform/internal/logging"
)

type EpisodeDataSource struct {
	client *believe.Client
}

var _ datasource.DataSourceWithConfigure = (*EpisodeDataSource)(nil)

func NewEpisodeDataSource() datasource.DataSource {
	return &EpisodeDataSource{}
}

func (d *EpisodeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_episode"
}

func (d *EpisodeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *EpisodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *EpisodeDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.FindOneBy != nil {
		params, diags := data.toListParams(ctx)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		env := EpisodesDataListDataSourceEnvelope{}
		page, err := d.client.Episodes.List(ctx, params)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		bytes := []byte(page.RawJSON())
		err = apijson.UnmarshalComputed(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}

		if count := len(env.Data.Elements()); count != 1 {
			resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
			return
		}
		ts, diags := env.Data.AsStructSliceT(ctx)
		resp.Diagnostics.Append(diags...)
		data.EpisodeID = ts[0].ID
	}

	res := new(http.Response)
	_, err := d.client.Episodes.Get(
		ctx,
		data.EpisodeID.ValueString(),
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
	data.ID = data.EpisodeID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
