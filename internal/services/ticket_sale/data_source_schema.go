// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ticket_sale

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*TicketSaleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Ticket sales with 300 records for practicing pagination, filtering, and financial data",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ticket_sale_id": schema.StringAttribute{
				Optional: true,
			},
			"buyer_email": schema.StringAttribute{
				Description: "Email of the ticket buyer",
				Computed:    true,
			},
			"buyer_name": schema.StringAttribute{
				Description: "Name of the ticket buyer",
				Computed:    true,
			},
			"coupon_code": schema.StringAttribute{
				Description: "Coupon code applied, if any",
				Computed:    true,
			},
			"currency": schema.StringAttribute{
				Description: "Currency code (GBP, USD, or EUR)",
				Computed:    true,
			},
			"discount": schema.StringAttribute{
				Description: "Discount amount applied from coupon",
				Computed:    true,
			},
			"match_id": schema.StringAttribute{
				Description: "ID of the match",
				Computed:    true,
			},
			"purchase_method": schema.StringAttribute{
				Description: "How the ticket was purchased\nAvailable values: \"online\", \"box_office\", \"will_call\", \"phone\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"online",
						"box_office",
						"will_call",
						"phone",
					),
				},
			},
			"quantity": schema.Int64Attribute{
				Description: "Number of tickets purchased",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 20),
				},
			},
			"subtotal": schema.StringAttribute{
				Description: "Subtotal before discount and tax (unit_price * quantity)",
				Computed:    true,
			},
			"tax": schema.StringAttribute{
				Description: "Tax amount (20% UK VAT on discounted subtotal)",
				Computed:    true,
			},
			"total": schema.StringAttribute{
				Description: "Final total (subtotal - discount + tax)",
				Computed:    true,
			},
			"unit_price": schema.StringAttribute{
				Description: "Price per ticket (decimal string)",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"coupon_code": schema.StringAttribute{
						Description: "Filter by coupon code (use 'none' for sales without coupons)",
						Optional:    true,
					},
					"currency": schema.StringAttribute{
						Description: "Filter by currency (GBP, USD, EUR)",
						Optional:    true,
					},
					"match_id": schema.StringAttribute{
						Description: "Filter by match ID",
						Optional:    true,
					},
					"purchase_method": schema.StringAttribute{
						Description: "Filter by purchase method\nAvailable values: \"online\", \"box_office\", \"will_call\", \"phone\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"online",
								"box_office",
								"will_call",
								"phone",
							),
						},
					},
				},
			},
		},
	}
}

func (d *TicketSaleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TicketSaleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("ticket_sale_id"), path.MatchRoot("find_one_by")),
	}
}
