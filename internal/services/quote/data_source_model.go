// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type QuoteDataSourceModel struct {
	ID              types.String                   `tfsdk:"id" path:"quote_id,computed"`
	QuoteID         types.String                   `tfsdk:"quote_id" path:"quote_id,optional"`
	CharacterID     types.String                   `tfsdk:"character_id" json:"character_id,computed"`
	Context         types.String                   `tfsdk:"context" json:"context,computed"`
	EpisodeID       types.String                   `tfsdk:"episode_id" json:"episode_id,computed"`
	IsFunny         types.Bool                     `tfsdk:"is_funny" json:"is_funny,computed"`
	IsInspirational types.Bool                     `tfsdk:"is_inspirational" json:"is_inspirational,computed"`
	MomentType      types.String                   `tfsdk:"moment_type" json:"moment_type,computed"`
	PopularityScore types.Float64                  `tfsdk:"popularity_score" json:"popularity_score,computed"`
	Text            types.String                   `tfsdk:"text" json:"text,computed"`
	Theme           types.String                   `tfsdk:"theme" json:"theme,computed"`
	TimesShared     types.Int64                    `tfsdk:"times_shared" json:"times_shared,computed"`
	SecondaryThemes customfield.List[types.String] `tfsdk:"secondary_themes" json:"secondary_themes,computed"`
	FindOneBy       *QuoteFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *QuoteDataSourceModel) toListParams(_ context.Context) (params believe.QuoteListParams, diags diag.Diagnostics) {
	params = believe.QuoteListParams{}

	if !m.FindOneBy.CharacterID.IsNull() {
		params.CharacterID = param.NewOpt(m.FindOneBy.CharacterID.ValueString())
	}
	if !m.FindOneBy.Funny.IsNull() {
		params.Funny = param.NewOpt(m.FindOneBy.Funny.ValueBool())
	}
	if !m.FindOneBy.Inspirational.IsNull() {
		params.Inspirational = param.NewOpt(m.FindOneBy.Inspirational.ValueBool())
	}
	if !m.FindOneBy.MomentType.IsNull() {
		params.MomentType = believe.QuoteMoment(m.FindOneBy.MomentType.ValueString())
	}
	if !m.FindOneBy.Theme.IsNull() {
		params.Theme = believe.QuoteTheme(m.FindOneBy.Theme.ValueString())
	}

	return
}

type QuoteFindOneByDataSourceModel struct {
	CharacterID   types.String `tfsdk:"character_id" query:"character_id,optional"`
	Funny         types.Bool   `tfsdk:"funny" query:"funny,optional"`
	Inspirational types.Bool   `tfsdk:"inspirational" query:"inspirational,optional"`
	MomentType    types.String `tfsdk:"moment_type" query:"moment_type,optional"`
	Theme         types.String `tfsdk:"theme" query:"theme,optional"`
}
