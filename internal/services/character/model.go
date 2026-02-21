// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character

import (
	"github.com/cjavdev/terraform-provider-believe/internal/apijson"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CharacterModel struct {
	ID                types.String                       `tfsdk:"id" json:"id,computed"`
	Background        types.String                       `tfsdk:"background" json:"background,required"`
	Name              types.String                       `tfsdk:"name" json:"name,required"`
	Role              types.String                       `tfsdk:"role" json:"role,required"`
	PersonalityTraits *[]types.String                    `tfsdk:"personality_traits" json:"personality_traits,required"`
	EmotionalStats    *CharacterEmotionalStatsModel      `tfsdk:"emotional_stats" json:"emotional_stats,required"`
	DateOfBirth       timetypes.RFC3339                  `tfsdk:"date_of_birth" json:"date_of_birth,optional" format:"date"`
	Email             types.String                       `tfsdk:"email" json:"email,optional"`
	HeightMeters      types.Float64                      `tfsdk:"height_meters" json:"height_meters,optional"`
	ProfileImageURL   types.String                       `tfsdk:"profile_image_url" json:"profile_image_url,optional"`
	TeamID            types.String                       `tfsdk:"team_id" json:"team_id,optional"`
	SignatureQuotes   *[]types.String                    `tfsdk:"signature_quotes" json:"signature_quotes,optional"`
	GrowthArcs        *[]*CharacterGrowthArcsModel       `tfsdk:"growth_arcs" json:"growth_arcs,optional"`
	SalaryGbp         customfield.NormalizedDynamicValue `tfsdk:"salary_gbp" json:"salary_gbp,optional,no_refresh"`
}

func (m CharacterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CharacterModel) MarshalJSONForUpdate(state CharacterModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type CharacterEmotionalStatsModel struct {
	Curiosity     types.Int64 `tfsdk:"curiosity" json:"curiosity,required"`
	Empathy       types.Int64 `tfsdk:"empathy" json:"empathy,required"`
	Optimism      types.Int64 `tfsdk:"optimism" json:"optimism,required"`
	Resilience    types.Int64 `tfsdk:"resilience" json:"resilience,required"`
	Vulnerability types.Int64 `tfsdk:"vulnerability" json:"vulnerability,required"`
}

type CharacterGrowthArcsModel struct {
	Breakthrough  types.String `tfsdk:"breakthrough" json:"breakthrough,required"`
	Challenge     types.String `tfsdk:"challenge" json:"challenge,required"`
	EndingPoint   types.String `tfsdk:"ending_point" json:"ending_point,required"`
	Season        types.Int64  `tfsdk:"season" json:"season,required"`
	StartingPoint types.String `tfsdk:"starting_point" json:"starting_point,required"`
}
