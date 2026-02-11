// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package coaching_principle

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CoachingPrinciplesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CoachingPrinciplesItemsDataSourceModel] `json:"data,computed"`
}

type CoachingPrinciplesDataSourceModel struct {
	MaxItems types.Int64                                                          `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[CoachingPrinciplesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CoachingPrinciplesDataSourceModel) toListParams(_ context.Context) (params believe.CoachingPrincipleListParams, diags diag.Diagnostics) {
	params = believe.CoachingPrincipleListParams{}

	return
}

type CoachingPrinciplesItemsDataSourceModel struct {
	ID              types.String `tfsdk:"id" json:"id,computed"`
	Application     types.String `tfsdk:"application" json:"application,computed"`
	ExampleFromShow types.String `tfsdk:"example_from_show" json:"example_from_show,computed"`
	Explanation     types.String `tfsdk:"explanation" json:"explanation,computed"`
	Principle       types.String `tfsdk:"principle" json:"principle,computed"`
	TedQuote        types.String `tfsdk:"ted_quote" json:"ted_quote,computed"`
}
