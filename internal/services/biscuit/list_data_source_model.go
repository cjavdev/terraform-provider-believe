// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package biscuit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type BiscuitsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[BiscuitsItemsDataSourceModel] `json:"data,computed"`
}

type BiscuitsDataSourceModel struct {
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[BiscuitsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *BiscuitsDataSourceModel) toListParams(_ context.Context) (params believe.BiscuitListParams, diags diag.Diagnostics) {
	params = believe.BiscuitListParams{}

	return
}

type BiscuitsItemsDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	Message       types.String `tfsdk:"message" json:"message,computed"`
	PairsWellWith types.String `tfsdk:"pairs_well_with" json:"pairs_well_with,computed"`
	TedNote       types.String `tfsdk:"ted_note" json:"ted_note,computed"`
	Type          types.String `tfsdk:"type" json:"type,computed"`
	WarmthLevel   types.Int64  `tfsdk:"warmth_level" json:"warmth_level,computed"`
}
