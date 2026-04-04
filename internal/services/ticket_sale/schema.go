// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ticket_sale

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*TicketSaleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Ticket sales with 300 records for practicing pagination, filtering, and financial data",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"buyer_name": schema.StringAttribute{
				Description: "Name of the ticket buyer",
				Required:    true,
			},
			"currency": schema.StringAttribute{
				Description: "Currency code (GBP, USD, or EUR)",
				Required:    true,
			},
			"discount": schema.StringAttribute{
				Description: "Discount amount applied from coupon",
				Required:    true,
			},
			"match_id": schema.StringAttribute{
				Description: "ID of the match",
				Required:    true,
			},
			"purchase_method": schema.StringAttribute{
				Description: "How the ticket was purchased\nAvailable values: \"online\", \"box_office\", \"will_call\", \"phone\".",
				Required:    true,
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
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 20),
				},
			},
			"subtotal": schema.StringAttribute{
				Description: "Subtotal before discount and tax (unit_price * quantity)",
				Required:    true,
			},
			"tax": schema.StringAttribute{
				Description: "Tax amount (20% UK VAT on discounted subtotal)",
				Required:    true,
			},
			"total": schema.StringAttribute{
				Description: "Final total (subtotal - discount + tax)",
				Required:    true,
			},
			"unit_price": schema.StringAttribute{
				Description: "Price per ticket (decimal string)",
				Required:    true,
			},
			"buyer_email": schema.StringAttribute{
				Description: "Email of the ticket buyer",
				Optional:    true,
			},
			"coupon_code": schema.StringAttribute{
				Description: "Coupon code applied, if any",
				Optional:    true,
			},
		},
	}
}

func (r *TicketSaleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TicketSaleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
