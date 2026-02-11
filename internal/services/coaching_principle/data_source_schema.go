// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package coaching_principle

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*CoachingPrincipleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"principle_id": schema.StringAttribute{
				Required: true,
			},
			"application": schema.StringAttribute{
				Description: "How to apply this principle",
				Computed:    true,
			},
			"example_from_show": schema.StringAttribute{
				Description: "Example from the show",
				Computed:    true,
			},
			"explanation": schema.StringAttribute{
				Description: "What this principle means",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Principle identifier",
				Computed:    true,
			},
			"principle": schema.StringAttribute{
				Description: "The coaching principle",
				Computed:    true,
			},
			"ted_quote": schema.StringAttribute{
				Description: "Related Ted quote",
				Computed:    true,
			},
		},
	}
}

func (d *CoachingPrincipleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CoachingPrincipleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
