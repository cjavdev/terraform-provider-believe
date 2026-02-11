// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EpisodeDataSourceModel struct {
	ID                     types.String                     `tfsdk:"id" path:"episode_id,computed"`
	EpisodeID              types.String                     `tfsdk:"episode_id" path:"episode_id,optional"`
	AirDate                timetypes.RFC3339                `tfsdk:"air_date" json:"air_date,computed" format:"date"`
	BiscuitsWithBossMoment types.String                     `tfsdk:"biscuits_with_boss_moment" json:"biscuits_with_boss_moment,computed"`
	Director               types.String                     `tfsdk:"director" json:"director,computed"`
	EpisodeNumber          types.Int64                      `tfsdk:"episode_number" json:"episode_number,computed"`
	MainTheme              types.String                     `tfsdk:"main_theme" json:"main_theme,computed"`
	RuntimeMinutes         types.Int64                      `tfsdk:"runtime_minutes" json:"runtime_minutes,computed"`
	Season                 types.Int64                      `tfsdk:"season" json:"season,computed"`
	Synopsis               types.String                     `tfsdk:"synopsis" json:"synopsis,computed"`
	TedWisdom              types.String                     `tfsdk:"ted_wisdom" json:"ted_wisdom,computed"`
	Title                  types.String                     `tfsdk:"title" json:"title,computed"`
	UsViewersMillions      types.Float64                    `tfsdk:"us_viewers_millions" json:"us_viewers_millions,computed"`
	ViewerRating           types.Float64                    `tfsdk:"viewer_rating" json:"viewer_rating,computed"`
	Writer                 types.String                     `tfsdk:"writer" json:"writer,computed"`
	CharacterFocus         customfield.List[types.String]   `tfsdk:"character_focus" json:"character_focus,computed"`
	MemorableMoments       customfield.List[types.String]   `tfsdk:"memorable_moments" json:"memorable_moments,computed"`
	FindOneBy              *EpisodeFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *EpisodeDataSourceModel) toListParams(_ context.Context) (params believe.EpisodeListParams, diags diag.Diagnostics) {
	params = believe.EpisodeListParams{}

	if !m.FindOneBy.CharacterFocus.IsNull() {
		params.CharacterFocus = param.NewOpt(m.FindOneBy.CharacterFocus.ValueString())
	}
	if !m.FindOneBy.Season.IsNull() {
		params.Season = param.NewOpt(m.FindOneBy.Season.ValueInt64())
	}

	return
}

type EpisodeFindOneByDataSourceModel struct {
	CharacterFocus types.String `tfsdk:"character_focus" query:"character_focus,optional"`
	Season         types.Int64  `tfsdk:"season" query:"season,optional"`
}
