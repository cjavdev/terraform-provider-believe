// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_ticket_sale

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ClientTicketSaleDataSourceModel struct {
	ID             types.String                              `tfsdk:"id" path:"ticket_sale_id,computed"`
	TicketSaleID   types.String                              `tfsdk:"ticket_sale_id" path:"ticket_sale_id,optional"`
	BuyerEmail     types.String                              `tfsdk:"buyer_email" json:"buyer_email,computed"`
	BuyerName      types.String                              `tfsdk:"buyer_name" json:"buyer_name,computed"`
	CouponCode     types.String                              `tfsdk:"coupon_code" json:"coupon_code,computed"`
	Currency       types.String                              `tfsdk:"currency" json:"currency,computed"`
	Discount       types.String                              `tfsdk:"discount" json:"discount,computed"`
	MatchID        types.String                              `tfsdk:"match_id" json:"match_id,computed"`
	PurchaseMethod types.String                              `tfsdk:"purchase_method" json:"purchase_method,computed"`
	Quantity       types.Int64                               `tfsdk:"quantity" json:"quantity,computed"`
	Subtotal       types.String                              `tfsdk:"subtotal" json:"subtotal,computed"`
	Tax            types.String                              `tfsdk:"tax" json:"tax,computed"`
	Total          types.String                              `tfsdk:"total" json:"total,computed"`
	UnitPrice      types.String                              `tfsdk:"unit_price" json:"unit_price,computed"`
	FindOneBy      *ClientTicketSaleFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *ClientTicketSaleDataSourceModel) toListParams(_ context.Context) (params believe.ClientTicketSaleListParams, diags diag.Diagnostics) {
	params = believe.ClientTicketSaleListParams{}

	if !m.FindOneBy.CouponCode.IsNull() {
		params.CouponCode = param.NewOpt(m.FindOneBy.CouponCode.ValueString())
	}
	if !m.FindOneBy.Currency.IsNull() {
		params.Currency = param.NewOpt(m.FindOneBy.Currency.ValueString())
	}
	if !m.FindOneBy.MatchID.IsNull() {
		params.MatchID = param.NewOpt(m.FindOneBy.MatchID.ValueString())
	}
	if !m.FindOneBy.PurchaseMethod.IsNull() {
		params.PurchaseMethod = believe.ClientTicketSaleListParamsPurchaseMethod(m.FindOneBy.PurchaseMethod.ValueString())
	}

	return
}

type ClientTicketSaleFindOneByDataSourceModel struct {
	CouponCode     types.String `tfsdk:"coupon_code" query:"coupon_code,optional"`
	Currency       types.String `tfsdk:"currency" query:"currency,optional"`
	MatchID        types.String `tfsdk:"match_id" query:"match_id,optional"`
	PurchaseMethod types.String `tfsdk:"purchase_method" query:"purchase_method,optional"`
}
