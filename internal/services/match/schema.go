// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package match

import (
	"context"

	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/cjavdev/terraform-provider-believe/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.ResourceWithConfigValidators = (*MatchResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"away_team_id": schema.StringAttribute{
				Description: "Away team ID",
				Required:    true,
			},
			"date": schema.StringAttribute{
				Description: "Match date and time",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"home_team_id": schema.StringAttribute{
				Description: "Home team ID",
				Required:    true,
			},
			"match_type": schema.StringAttribute{
				Description: "Type of match\nAvailable values: \"league\", \"cup\", \"friendly\", \"playoff\", \"final\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"league",
						"cup",
						"friendly",
						"playoff",
						"final",
					),
				},
			},
			"attendance": schema.Int64Attribute{
				Description: "Match attendance",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"episode_id": schema.StringAttribute{
				Description: "Episode ID where this match is featured",
				Optional:    true,
			},
			"lesson_learned": schema.StringAttribute{
				Description: "The life lesson learned from this match",
				Optional:    true,
			},
			"possession_percentage": schema.Float64Attribute{
				Description: "Home team possession percentage",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 100),
				},
			},
			"result": schema.StringAttribute{
				Description: "Match result from home team perspective\nAvailable values: \"win\", \"loss\", \"draw\", \"pending\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"win",
						"loss",
						"draw",
						"pending",
					),
				},
			},
			"ted_halftime_speech": schema.StringAttribute{
				Description: "Ted's inspirational halftime speech",
				Optional:    true,
			},
			"weather_temp_celsius": schema.Float64Attribute{
				Description: "Temperature at kickoff in Celsius",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(-30, 50),
				},
			},
			"turning_points": schema.ListNestedAttribute{
				Description: "Key moments that changed the match",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Description: "What happened",
							Required:    true,
						},
						"emotional_impact": schema.StringAttribute{
							Description: "How this affected the team emotionally",
							Required:    true,
						},
						"minute": schema.Int64Attribute{
							Description: "Minute of the match",
							Required:    true,
							Validators: []validator.Int64{
								int64validator.Between(0, 120),
							},
						},
						"character_involved": schema.StringAttribute{
							Description: "Character ID who was central to this moment",
							Optional:    true,
						},
					},
				},
			},
			"ticket_revenue_gbp": schema.DynamicAttribute{
				Description: "Total ticket revenue in GBP",
				Optional:    true,
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.Float64Type{}, basetypes.StringType{}),
				},
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
			"away_score": schema.Int64Attribute{
				Description: "Away team score",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default: int64default.StaticInt64(0),
			},
			"home_score": schema.Int64Attribute{
				Description: "Home team score",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default: int64default.StaticInt64(0),
			},
		},
	}
}

func (r *MatchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MatchResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
