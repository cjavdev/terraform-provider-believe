// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ticket_sale

import (
  "context"

  "github.com/cjavdev/believe-go"
  "github.com/cjavdev/believe-go/packages/param"
  "github.com/cjavdev/terraform-provider-believe/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type TicketSalesDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[TicketSalesItemsDataSourceModel] `json:"data,computed"`
}

type TicketSalesDataSourceModel struct {
CouponCode types.String `tfsdk:"coupon_code" query:"coupon_code,optional"`
Currency types.String `tfsdk:"currency" query:"currency,optional"`
MatchID types.String `tfsdk:"match_id" query:"match_id,optional"`
PurchaseMethod types.String `tfsdk:"purchase_method" query:"purchase_method,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[TicketSalesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *TicketSalesDataSourceModel) toListParams(_ context.Context) (params believe.TicketSaleListParams, diags diag.Diagnostics) {
  params = believe.TicketSaleListParams{

  }

  if !m.CouponCode.IsNull() {
    params.CouponCode = param.NewOpt(m.CouponCode.ValueString())
  }
  if !m.Currency.IsNull() {
    params.Currency = param.NewOpt(m.Currency.ValueString())
  }
  if !m.MatchID.IsNull() {
    params.MatchID = param.NewOpt(m.MatchID.ValueString())
  }
  if !m.PurchaseMethod.IsNull() {
    params.PurchaseMethod = believe.PurchaseMethod(m.PurchaseMethod.ValueString())
  }

  return
}

type TicketSalesItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
BuyerName types.String `tfsdk:"buyer_name" json:"buyer_name,computed"`
Currency types.String `tfsdk:"currency" json:"currency,computed"`
Discount types.String `tfsdk:"discount" json:"discount,computed"`
MatchID types.String `tfsdk:"match_id" json:"match_id,computed"`
PurchaseMethod types.String `tfsdk:"purchase_method" json:"purchase_method,computed"`
Quantity types.Int64 `tfsdk:"quantity" json:"quantity,computed"`
Subtotal types.String `tfsdk:"subtotal" json:"subtotal,computed"`
Tax types.String `tfsdk:"tax" json:"tax,computed"`
Total types.String `tfsdk:"total" json:"total,computed"`
UnitPrice types.String `tfsdk:"unit_price" json:"unit_price,computed"`
BuyerEmail types.String `tfsdk:"buyer_email" json:"buyer_email,computed"`
CouponCode types.String `tfsdk:"coupon_code" json:"coupon_code,computed"`
}
