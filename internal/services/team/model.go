// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/apijson"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type TeamModel struct {
	ID                types.String                       `tfsdk:"id" json:"id,computed"`
	CultureScore      types.Int64                        `tfsdk:"culture_score" json:"culture_score,required"`
	FoundedYear       types.Int64                        `tfsdk:"founded_year" json:"founded_year,required"`
	League            types.String                       `tfsdk:"league" json:"league,required"`
	Name              types.String                       `tfsdk:"name" json:"name,required"`
	Stadium           types.String                       `tfsdk:"stadium" json:"stadium,required"`
	Values            *TeamValuesModel                   `tfsdk:"values" json:"values,required"`
	AverageAttendance types.Float64                      `tfsdk:"average_attendance" json:"average_attendance,optional"`
	ContactEmail      types.String                       `tfsdk:"contact_email" json:"contact_email,optional"`
	Nickname          types.String                       `tfsdk:"nickname" json:"nickname,optional"`
	PrimaryColor      types.String                       `tfsdk:"primary_color" json:"primary_color,optional"`
	SecondaryColor    types.String                       `tfsdk:"secondary_color" json:"secondary_color,optional"`
	Website           types.String                       `tfsdk:"website" json:"website,optional"`
	WinPercentage     types.Float64                      `tfsdk:"win_percentage" json:"win_percentage,optional"`
	RivalTeams        *[]types.String                    `tfsdk:"rival_teams" json:"rival_teams,optional"`
	StadiumLocation   *TeamStadiumLocationModel          `tfsdk:"stadium_location" json:"stadium_location,optional"`
	AnnualBudgetGbp   customfield.NormalizedDynamicValue `tfsdk:"annual_budget_gbp" json:"annual_budget_gbp,optional,no_refresh"`
	IsActive          types.Bool                         `tfsdk:"is_active" json:"is_active,computed_optional"`
}

func (m TeamModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TeamModel) MarshalJSONForUpdate(state TeamModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type TeamValuesModel struct {
	PrimaryValue    types.String    `tfsdk:"primary_value" json:"primary_value,required"`
	SecondaryValues *[]types.String `tfsdk:"secondary_values" json:"secondary_values,required"`
	TeamMotto       types.String    `tfsdk:"team_motto" json:"team_motto,required"`
}

type TeamStadiumLocationModel struct {
	Latitude  types.Float64 `tfsdk:"latitude" json:"latitude,required"`
	Longitude types.Float64 `tfsdk:"longitude" json:"longitude,required"`
}
