// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QuotesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[QuotesItemsDataSourceModel] `json:"data,computed"`
}

type QuotesDataSourceModel struct {
	CharacterID   types.String                                             `tfsdk:"character_id" query:"character_id,optional"`
	Funny         types.Bool                                               `tfsdk:"funny" query:"funny,optional"`
	Inspirational types.Bool                                               `tfsdk:"inspirational" query:"inspirational,optional"`
	MomentType    types.String                                             `tfsdk:"moment_type" query:"moment_type,optional"`
	Theme         types.String                                             `tfsdk:"theme" query:"theme,optional"`
	MaxItems      types.Int64                                              `tfsdk:"max_items"`
	Items         customfield.NestedObjectList[QuotesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *QuotesDataSourceModel) toListParams(_ context.Context) (params believe.QuoteListParams, diags diag.Diagnostics) {
	params = believe.QuoteListParams{}

	if !m.CharacterID.IsNull() {
		params.CharacterID = param.NewOpt(m.CharacterID.ValueString())
	}
	if !m.Funny.IsNull() {
		params.Funny = param.NewOpt(m.Funny.ValueBool())
	}
	if !m.Inspirational.IsNull() {
		params.Inspirational = param.NewOpt(m.Inspirational.ValueBool())
	}
	if !m.MomentType.IsNull() {
		params.MomentType = believe.QuoteMoment(m.MomentType.ValueString())
	}
	if !m.Theme.IsNull() {
		params.Theme = believe.QuoteTheme(m.Theme.ValueString())
	}

	return
}

type QuotesItemsDataSourceModel struct {
	ID              types.String                   `tfsdk:"id" json:"id,computed"`
	CharacterID     types.String                   `tfsdk:"character_id" json:"character_id,computed"`
	Context         types.String                   `tfsdk:"context" json:"context,computed"`
	MomentType      types.String                   `tfsdk:"moment_type" json:"moment_type,computed"`
	Text            types.String                   `tfsdk:"text" json:"text,computed"`
	Theme           types.String                   `tfsdk:"theme" json:"theme,computed"`
	EpisodeID       types.String                   `tfsdk:"episode_id" json:"episode_id,computed"`
	IsFunny         types.Bool                     `tfsdk:"is_funny" json:"is_funny,computed"`
	IsInspirational types.Bool                     `tfsdk:"is_inspirational" json:"is_inspirational,computed"`
	PopularityScore types.Float64                  `tfsdk:"popularity_score" json:"popularity_score,computed"`
	SecondaryThemes customfield.List[types.String] `tfsdk:"secondary_themes" json:"secondary_themes,computed"`
	TimesShared     types.Int64                    `tfsdk:"times_shared" json:"times_shared,computed"`
}
