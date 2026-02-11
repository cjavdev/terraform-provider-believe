// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[TeamsItemsDataSourceModel] `json:"data,computed"`
}

type TeamsDataSourceModel struct {
	League          types.String                                            `tfsdk:"league" query:"league,optional"`
	MinCultureScore types.Int64                                             `tfsdk:"min_culture_score" query:"min_culture_score,optional"`
	MaxItems        types.Int64                                             `tfsdk:"max_items"`
	Items           customfield.NestedObjectList[TeamsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *TeamsDataSourceModel) toListParams(_ context.Context) (params believe.TeamListParams, diags diag.Diagnostics) {
	params = believe.TeamListParams{}

	if !m.League.IsNull() {
		params.League = believe.League(m.League.ValueString())
	}
	if !m.MinCultureScore.IsNull() {
		params.MinCultureScore = param.NewOpt(m.MinCultureScore.ValueInt64())
	}

	return
}

type TeamsItemsDataSourceModel struct {
	ID                types.String                                                  `tfsdk:"id" json:"id,computed"`
	CultureScore      types.Int64                                                   `tfsdk:"culture_score" json:"culture_score,computed"`
	FoundedYear       types.Int64                                                   `tfsdk:"founded_year" json:"founded_year,computed"`
	League            types.String                                                  `tfsdk:"league" json:"league,computed"`
	Name              types.String                                                  `tfsdk:"name" json:"name,computed"`
	Stadium           types.String                                                  `tfsdk:"stadium" json:"stadium,computed"`
	Values            customfield.NestedObject[TeamsValuesDataSourceModel]          `tfsdk:"values" json:"values,computed"`
	AnnualBudgetGbp   types.String                                                  `tfsdk:"annual_budget_gbp" json:"annual_budget_gbp,computed"`
	AverageAttendance types.Float64                                                 `tfsdk:"average_attendance" json:"average_attendance,computed"`
	ContactEmail      types.String                                                  `tfsdk:"contact_email" json:"contact_email,computed"`
	IsActive          types.Bool                                                    `tfsdk:"is_active" json:"is_active,computed"`
	Nickname          types.String                                                  `tfsdk:"nickname" json:"nickname,computed"`
	PrimaryColor      types.String                                                  `tfsdk:"primary_color" json:"primary_color,computed"`
	RivalTeams        customfield.List[types.String]                                `tfsdk:"rival_teams" json:"rival_teams,computed"`
	SecondaryColor    types.String                                                  `tfsdk:"secondary_color" json:"secondary_color,computed"`
	StadiumLocation   customfield.NestedObject[TeamsStadiumLocationDataSourceModel] `tfsdk:"stadium_location" json:"stadium_location,computed"`
	Website           types.String                                                  `tfsdk:"website" json:"website,computed"`
	WinPercentage     types.Float64                                                 `tfsdk:"win_percentage" json:"win_percentage,computed"`
}

type TeamsValuesDataSourceModel struct {
	PrimaryValue    types.String                   `tfsdk:"primary_value" json:"primary_value,computed"`
	SecondaryValues customfield.List[types.String] `tfsdk:"secondary_values" json:"secondary_values,computed"`
	TeamMotto       types.String                   `tfsdk:"team_motto" json:"team_motto,computed"`
}

type TeamsStadiumLocationDataSourceModel struct {
	Latitude  types.Float64 `tfsdk:"latitude" json:"latitude,computed"`
	Longitude types.Float64 `tfsdk:"longitude" json:"longitude,computed"`
}
