// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-go/packages/param"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type CharacterDataSourceModel struct {
	ID                types.String                                                     `tfsdk:"id" path:"character_id,computed"`
	CharacterID       types.String                                                     `tfsdk:"character_id" path:"character_id,optional"`
	Background        types.String                                                     `tfsdk:"background" json:"background,computed"`
	DateOfBirth       timetypes.RFC3339                                                `tfsdk:"date_of_birth" json:"date_of_birth,computed" format:"date"`
	Email             types.String                                                     `tfsdk:"email" json:"email,computed"`
	HeightMeters      types.Float64                                                    `tfsdk:"height_meters" json:"height_meters,computed"`
	Name              types.String                                                     `tfsdk:"name" json:"name,computed"`
	ProfileImageURL   types.String                                                     `tfsdk:"profile_image_url" json:"profile_image_url,computed"`
	Role              types.String                                                     `tfsdk:"role" json:"role,computed"`
	SalaryGbp         types.String                                                     `tfsdk:"salary_gbp" json:"salary_gbp,computed"`
	TeamID            types.String                                                     `tfsdk:"team_id" json:"team_id,computed"`
	PersonalityTraits customfield.List[types.String]                                   `tfsdk:"personality_traits" json:"personality_traits,computed"`
	SignatureQuotes   customfield.List[types.String]                                   `tfsdk:"signature_quotes" json:"signature_quotes,computed"`
	EmotionalStats    customfield.NestedObject[CharacterEmotionalStatsDataSourceModel] `tfsdk:"emotional_stats" json:"emotional_stats,computed"`
	GrowthArcs        customfield.NestedObjectList[CharacterGrowthArcsDataSourceModel] `tfsdk:"growth_arcs" json:"growth_arcs,computed"`
	FindOneBy         *CharacterFindOneByDataSourceModel                               `tfsdk:"find_one_by"`
}

func (m *CharacterDataSourceModel) toListParams(_ context.Context) (params believe.CharacterListParams, diags diag.Diagnostics) {
	params = believe.CharacterListParams{}

	if !m.FindOneBy.MinOptimism.IsNull() {
		params.MinOptimism = param.NewOpt(m.FindOneBy.MinOptimism.ValueInt64())
	}
	if !m.FindOneBy.Role.IsNull() {
		params.Role = believe.CharacterRole(m.FindOneBy.Role.ValueString())
	}
	if !m.FindOneBy.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.FindOneBy.TeamID.ValueString())
	}

	return
}

type CharacterEmotionalStatsDataSourceModel struct {
	Curiosity     types.Int64 `tfsdk:"curiosity" json:"curiosity,computed"`
	Empathy       types.Int64 `tfsdk:"empathy" json:"empathy,computed"`
	Optimism      types.Int64 `tfsdk:"optimism" json:"optimism,computed"`
	Resilience    types.Int64 `tfsdk:"resilience" json:"resilience,computed"`
	Vulnerability types.Int64 `tfsdk:"vulnerability" json:"vulnerability,computed"`
}

type CharacterGrowthArcsDataSourceModel struct {
	Breakthrough  types.String `tfsdk:"breakthrough" json:"breakthrough,computed"`
	Challenge     types.String `tfsdk:"challenge" json:"challenge,computed"`
	EndingPoint   types.String `tfsdk:"ending_point" json:"ending_point,computed"`
	Season        types.Int64  `tfsdk:"season" json:"season,computed"`
	StartingPoint types.String `tfsdk:"starting_point" json:"starting_point,computed"`
}

type CharacterFindOneByDataSourceModel struct {
	MinOptimism types.Int64  `tfsdk:"min_optimism" query:"min_optimism,optional"`
	Role        types.String `tfsdk:"role" query:"role,optional"`
	TeamID      types.String `tfsdk:"team_id" query:"team_id,optional"`
}
