// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package biscuit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*BiscuitDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"biscuit_id": schema.StringAttribute{
				Required: true,
			},
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
	}
}

func (d *BiscuitDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BiscuitDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
