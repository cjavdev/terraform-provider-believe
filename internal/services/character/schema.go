// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
	"github.com/stainless-sdks/believe-terraform/internal/customvalidator"
)

var _ resource.ResourceWithConfigValidators = (*CharacterResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"background": schema.StringAttribute{
				Description: "Character background and history",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Character's full name",
				Required:    true,
			},
			"role": schema.StringAttribute{
				Description: "Character's role\nAvailable values: \"coach\", \"player\", \"owner\", \"manager\", \"staff\", \"journalist\", \"family\", \"friend\", \"fan\", \"other\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"coach",
						"player",
						"owner",
						"manager",
						"staff",
						"journalist",
						"family",
						"friend",
						"fan",
						"other",
					),
				},
			},
			"personality_traits": schema.ListAttribute{
				Description: "Key personality traits",
				Required:    true,
				ElementType: types.StringType,
			},
			"emotional_stats": schema.SingleNestedAttribute{
				Description: "Emotional intelligence stats",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"curiosity": schema.Int64Attribute{
						Description: "Level of curiosity over judgment (0-100)",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"empathy": schema.Int64Attribute{
						Description: "Capacity for empathy (0-100)",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"optimism": schema.Int64Attribute{
						Description: "Level of optimism (0-100)",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"resilience": schema.Int64Attribute{
						Description: "Bounce-back ability (0-100)",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"vulnerability": schema.Int64Attribute{
						Description: "Willingness to be vulnerable (0-100)",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
				},
			},
			"date_of_birth": schema.StringAttribute{
				Description: "Character's date of birth",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"email": schema.StringAttribute{
				Description: "Character's email address",
				Optional:    true,
			},
			"height_meters": schema.Float64Attribute{
				Description: "Height in meters",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(1, 2.5),
				},
			},
			"profile_image_url": schema.StringAttribute{
				Description: "URL to character's profile image",
				Optional:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "ID of the team they belong to",
				Optional:    true,
			},
			"signature_quotes": schema.ListAttribute{
				Description: "Memorable quotes from this character",
				Optional:    true,
				ElementType: types.StringType,
			},
			"growth_arcs": schema.ListNestedAttribute{
				Description: "Character development across seasons",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"breakthrough": schema.StringAttribute{
							Description: "Key breakthrough moment",
							Required:    true,
						},
						"challenge": schema.StringAttribute{
							Description: "Main challenge faced",
							Required:    true,
						},
						"ending_point": schema.StringAttribute{
							Description: "Where the character ends up",
							Required:    true,
						},
						"season": schema.Int64Attribute{
							Description: "Season number",
							Required:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 3),
							},
						},
						"starting_point": schema.StringAttribute{
							Description: "Where the character starts emotionally",
							Required:    true,
						},
					},
				},
			},
			"salary_gbp": schema.DynamicAttribute{
				Description: "Annual salary in GBP",
				Optional:    true,
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.Float64Type{}, basetypes.StringType{}),
				},
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
		},
	}
}

func (r *CharacterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CharacterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
