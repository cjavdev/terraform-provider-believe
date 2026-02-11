// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_member

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamMembersDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[TeamMembersItemsDataSourceModel] `json:"data,computed"`
}

type TeamMembersDataSourceModel struct {
	MemberType types.String                                                  `tfsdk:"member_type" query:"member_type,optional"`
	TeamID     types.String                                                  `tfsdk:"team_id" query:"team_id,optional"`
	MaxItems   types.Int64                                                   `tfsdk:"max_items"`
	Items      customfield.NestedObjectList[TeamMembersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *TeamMembersDataSourceModel) toListParams(_ context.Context) (params believe.TeamMemberListParams, diags diag.Diagnostics) {
	params = believe.TeamMemberListParams{}

	if !m.MemberType.IsNull() {
		params.MemberType = believe.TeamMemberListParamsMemberType(m.MemberType.ValueString())
	}
	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type TeamMembersItemsDataSourceModel struct {
	ID               types.String                   `tfsdk:"id" json:"id,computed"`
	CharacterID      types.String                   `tfsdk:"character_id" json:"character_id,computed"`
	JerseyNumber     types.Int64                    `tfsdk:"jersey_number" json:"jersey_number,computed"`
	Position         types.String                   `tfsdk:"position" json:"position,computed"`
	TeamID           types.String                   `tfsdk:"team_id" json:"team_id,computed"`
	YearsWithTeam    types.Int64                    `tfsdk:"years_with_team" json:"years_with_team,computed"`
	Assists          types.Int64                    `tfsdk:"assists" json:"assists,computed"`
	GoalsScored      types.Int64                    `tfsdk:"goals_scored" json:"goals_scored,computed"`
	IsCaptain        types.Bool                     `tfsdk:"is_captain" json:"is_captain,computed"`
	MemberType       types.String                   `tfsdk:"member_type" json:"member_type,computed"`
	Specialty        types.String                   `tfsdk:"specialty" json:"specialty,computed"`
	Certifications   customfield.List[types.String] `tfsdk:"certifications" json:"certifications,computed"`
	WinRate          types.Float64                  `tfsdk:"win_rate" json:"win_rate,computed"`
	LicenseNumber    types.String                   `tfsdk:"license_number" json:"license_number,computed"`
	Qualifications   customfield.List[types.String] `tfsdk:"qualifications" json:"qualifications,computed"`
	IsHeadKitman     types.Bool                     `tfsdk:"is_head_kitman" json:"is_head_kitman,computed"`
	Responsibilities customfield.List[types.String] `tfsdk:"responsibilities" json:"responsibilities,computed"`
}
