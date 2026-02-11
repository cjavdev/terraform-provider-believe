// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/apijson"
)

type QuoteModel struct {
	ID              types.String    `tfsdk:"id" json:"id,computed"`
	CharacterID     types.String    `tfsdk:"character_id" json:"character_id,required"`
	Context         types.String    `tfsdk:"context" json:"context,required"`
	MomentType      types.String    `tfsdk:"moment_type" json:"moment_type,required"`
	Text            types.String    `tfsdk:"text" json:"text,required"`
	Theme           types.String    `tfsdk:"theme" json:"theme,required"`
	EpisodeID       types.String    `tfsdk:"episode_id" json:"episode_id,optional"`
	PopularityScore types.Float64   `tfsdk:"popularity_score" json:"popularity_score,optional"`
	TimesShared     types.Int64     `tfsdk:"times_shared" json:"times_shared,optional"`
	SecondaryThemes *[]types.String `tfsdk:"secondary_themes" json:"secondary_themes,optional"`
	IsFunny         types.Bool      `tfsdk:"is_funny" json:"is_funny,computed_optional"`
	IsInspirational types.Bool      `tfsdk:"is_inspirational" json:"is_inspirational,computed_optional"`
}

func (m QuoteModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m QuoteModel) MarshalJSONForUpdate(state QuoteModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
