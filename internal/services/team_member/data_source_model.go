// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_member

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type TeamMemberDataSourceModel struct {
	ID               types.String                        `tfsdk:"id" path:"member_id,computed"`
	MemberID         types.String                        `tfsdk:"member_id" path:"member_id,optional"`
	Assists          types.Int64                         `tfsdk:"assists" json:"assists,computed"`
	CharacterID      types.String                        `tfsdk:"character_id" json:"character_id,computed"`
	GoalsScored      types.Int64                         `tfsdk:"goals_scored" json:"goals_scored,computed"`
	IsCaptain        types.Bool                          `tfsdk:"is_captain" json:"is_captain,computed"`
	IsHeadKitman     types.Bool                          `tfsdk:"is_head_kitman" json:"is_head_kitman,computed"`
	JerseyNumber     types.Int64                         `tfsdk:"jersey_number" json:"jersey_number,computed"`
	LicenseNumber    types.String                        `tfsdk:"license_number" json:"license_number,computed"`
	MemberType       types.String                        `tfsdk:"member_type" json:"member_type,computed"`
	Position         types.String                        `tfsdk:"position" json:"position,computed"`
	Specialty        types.String                        `tfsdk:"specialty" json:"specialty,computed"`
	TeamID           types.String                        `tfsdk:"team_id" json:"team_id,computed"`
	WinRate          types.Float64                       `tfsdk:"win_rate" json:"win_rate,computed"`
	YearsWithTeam    types.Int64                         `tfsdk:"years_with_team" json:"years_with_team,computed"`
	Certifications   customfield.List[types.String]      `tfsdk:"certifications" json:"certifications,computed"`
	Qualifications   customfield.List[types.String]      `tfsdk:"qualifications" json:"qualifications,computed"`
	Responsibilities customfield.List[types.String]      `tfsdk:"responsibilities" json:"responsibilities,computed"`
	FindOneBy        *TeamMemberFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *TeamMemberDataSourceModel) toListParams(_ context.Context) (params believe.TeamMemberListParams, diags diag.Diagnostics) {
	params = believe.TeamMemberListParams{}

	if !m.FindOneBy.MemberType.IsNull() {
		params.MemberType = believe.TeamMemberListParamsMemberType(m.FindOneBy.MemberType.ValueString())
	}
	if !m.FindOneBy.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.FindOneBy.TeamID.ValueString())
	}

	return
}

type TeamMemberFindOneByDataSourceModel struct {
	MemberType types.String `tfsdk:"member_type" query:"member_type,optional"`
	TeamID     types.String `tfsdk:"team_id" query:"team_id,optional"`
}
