// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package match

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/apijson"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type MatchModel struct {
	ID                   types.String                       `tfsdk:"id" json:"id,computed"`
	AwayTeamID           types.String                       `tfsdk:"away_team_id" json:"away_team_id,required"`
	Date                 timetypes.RFC3339                  `tfsdk:"date" json:"date,required" format:"date-time"`
	HomeTeamID           types.String                       `tfsdk:"home_team_id" json:"home_team_id,required"`
	MatchType            types.String                       `tfsdk:"match_type" json:"match_type,required"`
	Attendance           types.Int64                        `tfsdk:"attendance" json:"attendance,optional"`
	EpisodeID            types.String                       `tfsdk:"episode_id" json:"episode_id,optional"`
	LessonLearned        types.String                       `tfsdk:"lesson_learned" json:"lesson_learned,optional"`
	PossessionPercentage types.Float64                      `tfsdk:"possession_percentage" json:"possession_percentage,optional"`
	Result               types.String                       `tfsdk:"result" json:"result,optional"`
	TedHalftimeSpeech    types.String                       `tfsdk:"ted_halftime_speech" json:"ted_halftime_speech,optional"`
	WeatherTempCelsius   types.Float64                      `tfsdk:"weather_temp_celsius" json:"weather_temp_celsius,optional"`
	TurningPoints        *[]*MatchTurningPointsModel        `tfsdk:"turning_points" json:"turning_points,optional"`
	TicketRevenueGbp     customfield.NormalizedDynamicValue `tfsdk:"ticket_revenue_gbp" json:"ticket_revenue_gbp,optional,no_refresh"`
	AwayScore            types.Int64                        `tfsdk:"away_score" json:"away_score,computed_optional"`
	HomeScore            types.Int64                        `tfsdk:"home_score" json:"home_score,computed_optional"`
}

func (m MatchModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MatchModel) MarshalJSONForUpdate(state MatchModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type MatchTurningPointsModel struct {
	Description       types.String `tfsdk:"description" json:"description,required"`
	EmotionalImpact   types.String `tfsdk:"emotional_impact" json:"emotional_impact,required"`
	Minute            types.Int64  `tfsdk:"minute" json:"minute,required"`
	CharacterInvolved types.String `tfsdk:"character_involved" json:"character_involved,optional"`
}
