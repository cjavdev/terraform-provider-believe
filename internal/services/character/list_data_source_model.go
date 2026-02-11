// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CharactersDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CharactersItemsDataSourceModel] `json:"data,computed"`
}

type CharactersDataSourceModel struct {
	MinOptimism types.Int64                                                  `tfsdk:"min_optimism" query:"min_optimism,optional"`
	Role        types.String                                                 `tfsdk:"role" query:"role,optional"`
	TeamID      types.String                                                 `tfsdk:"team_id" query:"team_id,optional"`
	MaxItems    types.Int64                                                  `tfsdk:"max_items"`
	Items       customfield.NestedObjectList[CharactersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CharactersDataSourceModel) toListParams(_ context.Context) (params believe.CharacterListParams, diags diag.Diagnostics) {
	params = believe.CharacterListParams{}

	if !m.MinOptimism.IsNull() {
		params.MinOptimism = param.NewOpt(m.MinOptimism.ValueInt64())
	}
	if !m.Role.IsNull() {
		params.Role = believe.CharacterRole(m.Role.ValueString())
	}
	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type CharactersItemsDataSourceModel struct {
	ID                types.String                                                      `tfsdk:"id" json:"id,computed"`
	Background        types.String                                                      `tfsdk:"background" json:"background,computed"`
	EmotionalStats    customfield.NestedObject[CharactersEmotionalStatsDataSourceModel] `tfsdk:"emotional_stats" json:"emotional_stats,computed"`
	Name              types.String                                                      `tfsdk:"name" json:"name,computed"`
	PersonalityTraits customfield.List[types.String]                                    `tfsdk:"personality_traits" json:"personality_traits,computed"`
	Role              types.String                                                      `tfsdk:"role" json:"role,computed"`
	DateOfBirth       timetypes.RFC3339                                                 `tfsdk:"date_of_birth" json:"date_of_birth,computed" format:"date"`
	Email             types.String                                                      `tfsdk:"email" json:"email,computed"`
	GrowthArcs        customfield.NestedObjectList[CharactersGrowthArcsDataSourceModel] `tfsdk:"growth_arcs" json:"growth_arcs,computed"`
	HeightMeters      types.Float64                                                     `tfsdk:"height_meters" json:"height_meters,computed"`
	ProfileImageURL   types.String                                                      `tfsdk:"profile_image_url" json:"profile_image_url,computed"`
	SalaryGbp         types.String                                                      `tfsdk:"salary_gbp" json:"salary_gbp,computed"`
	SignatureQuotes   customfield.List[types.String]                                    `tfsdk:"signature_quotes" json:"signature_quotes,computed"`
	TeamID            types.String                                                      `tfsdk:"team_id" json:"team_id,computed"`
}

type CharactersEmotionalStatsDataSourceModel struct {
	Curiosity     types.Int64 `tfsdk:"curiosity" json:"curiosity,computed"`
	Empathy       types.Int64 `tfsdk:"empathy" json:"empathy,computed"`
	Optimism      types.Int64 `tfsdk:"optimism" json:"optimism,computed"`
	Resilience    types.Int64 `tfsdk:"resilience" json:"resilience,computed"`
	Vulnerability types.Int64 `tfsdk:"vulnerability" json:"vulnerability,computed"`
}

type CharactersGrowthArcsDataSourceModel struct {
	Breakthrough  types.String `tfsdk:"breakthrough" json:"breakthrough,computed"`
	Challenge     types.String `tfsdk:"challenge" json:"challenge,computed"`
	EndingPoint   types.String `tfsdk:"ending_point" json:"ending_point,computed"`
	Season        types.Int64  `tfsdk:"season" json:"season,computed"`
	StartingPoint types.String `tfsdk:"starting_point" json:"starting_point,computed"`
}
