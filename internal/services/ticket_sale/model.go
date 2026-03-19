// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ticket_sale

import (
	"github.com/cjavdev/terraform-provider-believe/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TicketSaleModel struct {
	ID             types.String `tfsdk:"id" json:"id,computed"`
	BuyerName      types.String `tfsdk:"buyer_name" json:"buyer_name,required"`
	Currency       types.String `tfsdk:"currency" json:"currency,required"`
	Discount       types.String `tfsdk:"discount" json:"discount,required"`
	MatchID        types.String `tfsdk:"match_id" json:"match_id,required"`
	PurchaseMethod types.String `tfsdk:"purchase_method" json:"purchase_method,required"`
	Quantity       types.Int64  `tfsdk:"quantity" json:"quantity,required"`
	Subtotal       types.String `tfsdk:"subtotal" json:"subtotal,required"`
	Tax            types.String `tfsdk:"tax" json:"tax,required"`
	Total          types.String `tfsdk:"total" json:"total,required"`
	UnitPrice      types.String `tfsdk:"unit_price" json:"unit_price,required"`
	BuyerEmail     types.String `tfsdk:"buyer_email" json:"buyer_email,optional"`
	CouponCode     types.String `tfsdk:"coupon_code" json:"coupon_code,optional"`
}

func (m TicketSaleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TicketSaleModel) MarshalJSONForUpdate(state TicketSaleModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
