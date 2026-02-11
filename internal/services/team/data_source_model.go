// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-go/packages/param"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type TeamDataSourceModel struct {
	ID                types.String                                                 `tfsdk:"id" path:"team_id,computed"`
	TeamID            types.String                                                 `tfsdk:"team_id" path:"team_id,optional"`
	AnnualBudgetGbp   types.String                                                 `tfsdk:"annual_budget_gbp" json:"annual_budget_gbp,computed"`
	AverageAttendance types.Float64                                                `tfsdk:"average_attendance" json:"average_attendance,computed"`
	ContactEmail      types.String                                                 `tfsdk:"contact_email" json:"contact_email,computed"`
	CultureScore      types.Int64                                                  `tfsdk:"culture_score" json:"culture_score,computed"`
	FoundedYear       types.Int64                                                  `tfsdk:"founded_year" json:"founded_year,computed"`
	IsActive          types.Bool                                                   `tfsdk:"is_active" json:"is_active,computed"`
	League            types.String                                                 `tfsdk:"league" json:"league,computed"`
	Name              types.String                                                 `tfsdk:"name" json:"name,computed"`
	Nickname          types.String                                                 `tfsdk:"nickname" json:"nickname,computed"`
	PrimaryColor      types.String                                                 `tfsdk:"primary_color" json:"primary_color,computed"`
	SecondaryColor    types.String                                                 `tfsdk:"secondary_color" json:"secondary_color,computed"`
	Stadium           types.String                                                 `tfsdk:"stadium" json:"stadium,computed"`
	Website           types.String                                                 `tfsdk:"website" json:"website,computed"`
	WinPercentage     types.Float64                                                `tfsdk:"win_percentage" json:"win_percentage,computed"`
	RivalTeams        customfield.List[types.String]                               `tfsdk:"rival_teams" json:"rival_teams,computed"`
	StadiumLocation   customfield.NestedObject[TeamStadiumLocationDataSourceModel] `tfsdk:"stadium_location" json:"stadium_location,computed"`
	Values            customfield.NestedObject[TeamValuesDataSourceModel]          `tfsdk:"values" json:"values,computed"`
	FindOneBy         *TeamFindOneByDataSourceModel                                `tfsdk:"find_one_by"`
}

func (m *TeamDataSourceModel) toListParams(_ context.Context) (params believe.TeamListParams, diags diag.Diagnostics) {
	params = believe.TeamListParams{}

	if !m.FindOneBy.League.IsNull() {
		params.League = believe.League(m.FindOneBy.League.ValueString())
	}
	if !m.FindOneBy.MinCultureScore.IsNull() {
		params.MinCultureScore = param.NewOpt(m.FindOneBy.MinCultureScore.ValueInt64())
	}

	return
}

type TeamStadiumLocationDataSourceModel struct {
	Latitude  types.Float64 `tfsdk:"latitude" json:"latitude,computed"`
	Longitude types.Float64 `tfsdk:"longitude" json:"longitude,computed"`
}

type TeamValuesDataSourceModel struct {
	PrimaryValue    types.String                   `tfsdk:"primary_value" json:"primary_value,computed"`
	SecondaryValues customfield.List[types.String] `tfsdk:"secondary_values" json:"secondary_values,computed"`
	TeamMotto       types.String                   `tfsdk:"team_motto" json:"team_motto,computed"`
}

type TeamFindOneByDataSourceModel struct {
	League          types.String `tfsdk:"league" query:"league,optional"`
	MinCultureScore types.Int64  `tfsdk:"min_culture_score" query:"min_culture_score,optional"`
}
