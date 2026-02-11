// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package coaching_principle

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CoachingPrincipleDataSourceModel struct {
	PrincipleID     types.String `tfsdk:"principle_id" path:"principle_id,required"`
	Application     types.String `tfsdk:"application" json:"application,computed"`
	ExampleFromShow types.String `tfsdk:"example_from_show" json:"example_from_show,computed"`
	Explanation     types.String `tfsdk:"explanation" json:"explanation,computed"`
	ID              types.String `tfsdk:"id" json:"id,computed"`
	Principle       types.String `tfsdk:"principle" json:"principle,computed"`
	TedQuote        types.String `tfsdk:"ted_quote" json:"ted_quote,computed"`
}
