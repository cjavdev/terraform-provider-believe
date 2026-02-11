// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package biscuit

import (
	"context"

	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*BiscuitsDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[BiscuitsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Biscuit identifier",
							Computed:    true,
						},
						"message": schema.StringAttribute{
							Description: "Message that comes with the biscuit",
							Computed:    true,
						},
						"pairs_well_with": schema.StringAttribute{
							Description: "What this biscuit pairs well with",
							Computed:    true,
						},
						"ted_note": schema.StringAttribute{
							Description: "A handwritten note from Ted",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Type of biscuit\nAvailable values: \"classic\", \"shortbread\", \"chocolate_chip\", \"oatmeal_raisin\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"classic",
									"shortbread",
									"chocolate_chip",
									"oatmeal_raisin",
								),
							},
						},
						"warmth_level": schema.Int64Attribute{
							Description: "How warm and fresh (1-10)",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
					},
				},
			},
		},
	}
}

func (d *BiscuitsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *BiscuitsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
