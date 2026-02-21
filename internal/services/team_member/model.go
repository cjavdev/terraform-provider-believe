// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_member

import (
	"github.com/cjavdev/terraform-provider-believe/internal/apijson"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamMemberModel struct {
	ID               types.String                   `tfsdk:"id" json:"id,computed"`
	Member           *TeamMemberMemberModel         `tfsdk:"member" json:"member,required,no_refresh"`
	Updates          *TeamMemberUpdatesModel        `tfsdk:"updates" json:"updates,optional,no_refresh"`
	Assists          types.Int64                    `tfsdk:"assists" json:"assists,computed"`
	CharacterID      types.String                   `tfsdk:"character_id" json:"character_id,computed"`
	GoalsScored      types.Int64                    `tfsdk:"goals_scored" json:"goals_scored,computed"`
	IsCaptain        types.Bool                     `tfsdk:"is_captain" json:"is_captain,computed"`
	IsHeadKitman     types.Bool                     `tfsdk:"is_head_kitman" json:"is_head_kitman,computed"`
	JerseyNumber     types.Int64                    `tfsdk:"jersey_number" json:"jersey_number,computed"`
	LicenseNumber    types.String                   `tfsdk:"license_number" json:"license_number,computed"`
	MemberType       types.String                   `tfsdk:"member_type" json:"member_type,computed"`
	Position         types.String                   `tfsdk:"position" json:"position,computed"`
	Specialty        types.String                   `tfsdk:"specialty" json:"specialty,computed"`
	TeamID           types.String                   `tfsdk:"team_id" json:"team_id,computed"`
	WinRate          types.Float64                  `tfsdk:"win_rate" json:"win_rate,computed"`
	YearsWithTeam    types.Int64                    `tfsdk:"years_with_team" json:"years_with_team,computed"`
	Certifications   customfield.List[types.String] `tfsdk:"certifications" json:"certifications,computed"`
	Qualifications   customfield.List[types.String] `tfsdk:"qualifications" json:"qualifications,computed"`
	Responsibilities customfield.List[types.String] `tfsdk:"responsibilities" json:"responsibilities,computed"`
}

func (m TeamMemberModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Member)
}

func (m TeamMemberModel) MarshalJSONForUpdate(state TeamMemberModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m.Member, state.Member)
}

type TeamMemberMemberModel struct {
	CharacterID      types.String    `tfsdk:"character_id" json:"character_id,required"`
	JerseyNumber     types.Int64     `tfsdk:"jersey_number" json:"jersey_number,optional"`
	Position         types.String    `tfsdk:"position" json:"position,optional"`
	TeamID           types.String    `tfsdk:"team_id" json:"team_id,required"`
	YearsWithTeam    types.Int64     `tfsdk:"years_with_team" json:"years_with_team,required"`
	Assists          types.Int64     `tfsdk:"assists" json:"assists,computed_optional"`
	GoalsScored      types.Int64     `tfsdk:"goals_scored" json:"goals_scored,computed_optional"`
	IsCaptain        types.Bool      `tfsdk:"is_captain" json:"is_captain,computed_optional"`
	MemberType       types.String    `tfsdk:"member_type" json:"member_type,computed_optional"`
	Specialty        types.String    `tfsdk:"specialty" json:"specialty,optional"`
	Certifications   *[]types.String `tfsdk:"certifications" json:"certifications,optional"`
	WinRate          types.Float64   `tfsdk:"win_rate" json:"win_rate,optional"`
	LicenseNumber    types.String    `tfsdk:"license_number" json:"license_number,optional"`
	Qualifications   *[]types.String `tfsdk:"qualifications" json:"qualifications,optional"`
	IsHeadKitman     types.Bool      `tfsdk:"is_head_kitman" json:"is_head_kitman,computed_optional"`
	Responsibilities *[]types.String `tfsdk:"responsibilities" json:"responsibilities,optional"`
}

type TeamMemberUpdatesModel struct {
	Assists          types.Int64     `tfsdk:"assists" json:"assists,optional"`
	GoalsScored      types.Int64     `tfsdk:"goals_scored" json:"goals_scored,optional"`
	IsCaptain        types.Bool      `tfsdk:"is_captain" json:"is_captain,optional"`
	JerseyNumber     types.Int64     `tfsdk:"jersey_number" json:"jersey_number,optional"`
	Position         types.String    `tfsdk:"position" json:"position,optional"`
	TeamID           types.String    `tfsdk:"team_id" json:"team_id,optional"`
	YearsWithTeam    types.Int64     `tfsdk:"years_with_team" json:"years_with_team,optional"`
	Certifications   *[]types.String `tfsdk:"certifications" json:"certifications,optional"`
	Specialty        types.String    `tfsdk:"specialty" json:"specialty,optional"`
	WinRate          types.Float64   `tfsdk:"win_rate" json:"win_rate,optional"`
	LicenseNumber    types.String    `tfsdk:"license_number" json:"license_number,optional"`
	Qualifications   *[]types.String `tfsdk:"qualifications" json:"qualifications,optional"`
	IsHeadKitman     types.Bool      `tfsdk:"is_head_kitman" json:"is_head_kitman,optional"`
	Responsibilities *[]types.String `tfsdk:"responsibilities" json:"responsibilities,optional"`
}
