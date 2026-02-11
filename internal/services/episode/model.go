// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode

import (
	"github.com/cjavdev/terraform-provider-believe/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EpisodeModel struct {
	ID                     types.String      `tfsdk:"id" json:"id,computed"`
	AirDate                timetypes.RFC3339 `tfsdk:"air_date" json:"air_date,required" format:"date"`
	Director               types.String      `tfsdk:"director" json:"director,required"`
	EpisodeNumber          types.Int64       `tfsdk:"episode_number" json:"episode_number,required"`
	MainTheme              types.String      `tfsdk:"main_theme" json:"main_theme,required"`
	RuntimeMinutes         types.Int64       `tfsdk:"runtime_minutes" json:"runtime_minutes,required"`
	Season                 types.Int64       `tfsdk:"season" json:"season,required"`
	Synopsis               types.String      `tfsdk:"synopsis" json:"synopsis,required"`
	TedWisdom              types.String      `tfsdk:"ted_wisdom" json:"ted_wisdom,required"`
	Title                  types.String      `tfsdk:"title" json:"title,required"`
	Writer                 types.String      `tfsdk:"writer" json:"writer,required"`
	CharacterFocus         *[]types.String   `tfsdk:"character_focus" json:"character_focus,required"`
	BiscuitsWithBossMoment types.String      `tfsdk:"biscuits_with_boss_moment" json:"biscuits_with_boss_moment,optional"`
	UsViewersMillions      types.Float64     `tfsdk:"us_viewers_millions" json:"us_viewers_millions,optional"`
	ViewerRating           types.Float64     `tfsdk:"viewer_rating" json:"viewer_rating,optional"`
	MemorableMoments       *[]types.String   `tfsdk:"memorable_moments" json:"memorable_moments,optional"`
}

func (m EpisodeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EpisodeModel) MarshalJSONForUpdate(state EpisodeModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
