// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package match

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-go/packages/param"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type MatchDataSourceModel struct {
	ID                   types.String                                                    `tfsdk:"id" path:"match_id,computed"`
	MatchID              types.String                                                    `tfsdk:"match_id" path:"match_id,optional"`
	Attendance           types.Int64                                                     `tfsdk:"attendance" json:"attendance,computed"`
	AwayScore            types.Int64                                                     `tfsdk:"away_score" json:"away_score,computed"`
	AwayTeamID           types.String                                                    `tfsdk:"away_team_id" json:"away_team_id,computed"`
	Date                 timetypes.RFC3339                                               `tfsdk:"date" json:"date,computed" format:"date-time"`
	EpisodeID            types.String                                                    `tfsdk:"episode_id" json:"episode_id,computed"`
	HomeScore            types.Int64                                                     `tfsdk:"home_score" json:"home_score,computed"`
	HomeTeamID           types.String                                                    `tfsdk:"home_team_id" json:"home_team_id,computed"`
	LessonLearned        types.String                                                    `tfsdk:"lesson_learned" json:"lesson_learned,computed"`
	MatchType            types.String                                                    `tfsdk:"match_type" json:"match_type,computed"`
	PossessionPercentage types.Float64                                                   `tfsdk:"possession_percentage" json:"possession_percentage,computed"`
	Result               types.String                                                    `tfsdk:"result" json:"result,computed"`
	TedHalftimeSpeech    types.String                                                    `tfsdk:"ted_halftime_speech" json:"ted_halftime_speech,computed"`
	TicketRevenueGbp     types.String                                                    `tfsdk:"ticket_revenue_gbp" json:"ticket_revenue_gbp,computed"`
	WeatherTempCelsius   types.Float64                                                   `tfsdk:"weather_temp_celsius" json:"weather_temp_celsius,computed"`
	TurningPoints        customfield.NestedObjectList[MatchTurningPointsDataSourceModel] `tfsdk:"turning_points" json:"turning_points,computed"`
	FindOneBy            *MatchFindOneByDataSourceModel                                  `tfsdk:"find_one_by"`
}

func (m *MatchDataSourceModel) toListParams(_ context.Context) (params believe.MatchListParams, diags diag.Diagnostics) {
	params = believe.MatchListParams{}

	if !m.FindOneBy.MatchType.IsNull() {
		params.MatchType = believe.MatchType(m.FindOneBy.MatchType.ValueString())
	}
	if !m.FindOneBy.Result.IsNull() {
		params.Result = believe.MatchResult(m.FindOneBy.Result.ValueString())
	}
	if !m.FindOneBy.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.FindOneBy.TeamID.ValueString())
	}

	return
}

type MatchTurningPointsDataSourceModel struct {
	Description       types.String `tfsdk:"description" json:"description,computed"`
	EmotionalImpact   types.String `tfsdk:"emotional_impact" json:"emotional_impact,computed"`
	Minute            types.Int64  `tfsdk:"minute" json:"minute,computed"`
	CharacterInvolved types.String `tfsdk:"character_involved" json:"character_involved,computed"`
}

type MatchFindOneByDataSourceModel struct {
	MatchType types.String `tfsdk:"match_type" query:"match_type,optional"`
	Result    types.String `tfsdk:"result" query:"result,optional"`
	TeamID    types.String `tfsdk:"team_id" query:"team_id,optional"`
}
