// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package coaching_principle

import (
	"context"

	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*CoachingPrinciplesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CoachingPrinciplesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Principle identifier",
							Computed:    true,
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
						"principle": schema.StringAttribute{
							Description: "The coaching principle",
							Computed:    true,
						},
						"ted_quote": schema.StringAttribute{
							Description: "Related Ted quote",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *CoachingPrinciplesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CoachingPrinciplesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
