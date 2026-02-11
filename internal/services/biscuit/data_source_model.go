// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package biscuit

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BiscuitDataSourceModel struct {
	BiscuitID     types.String `tfsdk:"biscuit_id" path:"biscuit_id,required"`
	ID            types.String `tfsdk:"id" json:"id,computed"`
	Message       types.String `tfsdk:"message" json:"message,computed"`
	PairsWellWith types.String `tfsdk:"pairs_well_with" json:"pairs_well_with,computed"`
	TedNote       types.String `tfsdk:"ted_note" json:"ted_note,computed"`
	Type          types.String `tfsdk:"type" json:"type,computed"`
	WarmthLevel   types.Int64  `tfsdk:"warmth_level" json:"warmth_level,computed"`
}
