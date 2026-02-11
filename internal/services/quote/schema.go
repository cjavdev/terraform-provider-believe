// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*QuoteResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"character_id": schema.StringAttribute{
				Description: "ID of the character who said it",
				Required:    true,
			},
			"context": schema.StringAttribute{
				Description: "Context in which the quote was said",
				Required:    true,
			},
			"moment_type": schema.StringAttribute{
				Description: "Type of moment when the quote was said\nAvailable values: \"halftime_speech\", \"press_conference\", \"locker_room\", \"training\", \"biscuits_with_boss\", \"pub\", \"one_on_one\", \"celebration\", \"crisis\", \"casual\", \"confrontation\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"halftime_speech",
						"press_conference",
						"locker_room",
						"training",
						"biscuits_with_boss",
						"pub",
						"one_on_one",
						"celebration",
						"crisis",
						"casual",
						"confrontation",
					),
				},
			},
			"text": schema.StringAttribute{
				Description: "The quote text",
				Required:    true,
			},
			"theme": schema.StringAttribute{
				Description: "Primary theme of the quote\nAvailable values: \"belief\", \"teamwork\", \"curiosity\", \"kindness\", \"resilience\", \"vulnerability\", \"growth\", \"humor\", \"wisdom\", \"leadership\", \"love\", \"forgiveness\", \"philosophy\", \"romance\", \"cultural-pride\", \"cultural-differences\", \"antagonism\", \"celebration\", \"identity\", \"isolation\", \"power\", \"sacrifice\", \"standards\", \"confidence\", \"conflict\", \"honesty\", \"integrity\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"belief",
						"teamwork",
						"curiosity",
						"kindness",
						"resilience",
						"vulnerability",
						"growth",
						"humor",
						"wisdom",
						"leadership",
						"love",
						"forgiveness",
						"philosophy",
						"romance",
						"cultural-pride",
						"cultural-differences",
						"antagonism",
						"celebration",
						"identity",
						"isolation",
						"power",
						"sacrifice",
						"standards",
						"confidence",
						"conflict",
						"honesty",
						"integrity",
					),
				},
			},
			"episode_id": schema.StringAttribute{
				Description: "Episode where the quote appears",
				Optional:    true,
			},
			"popularity_score": schema.Float64Attribute{
				Description: "Popularity/virality score (0-100)",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 100),
				},
			},
			"times_shared": schema.Int64Attribute{
				Description: "Number of times shared on social media",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"secondary_themes": schema.ListAttribute{
				Description: "Additional themes",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"belief",
							"teamwork",
							"curiosity",
							"kindness",
							"resilience",
							"vulnerability",
							"growth",
							"humor",
							"wisdom",
							"leadership",
							"love",
							"forgiveness",
							"philosophy",
							"romance",
							"cultural-pride",
							"cultural-differences",
							"antagonism",
							"celebration",
							"identity",
							"isolation",
							"power",
							"sacrifice",
							"standards",
							"confidence",
							"conflict",
							"honesty",
							"integrity",
						),
					),
				},
				ElementType: types.StringType,
			},
			"is_funny": schema.BoolAttribute{
				Description: "Whether this quote is humorous",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_inspirational": schema.BoolAttribute{
				Description: "Whether this quote is inspirational",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
		},
	}
}

func (r *QuoteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *QuoteResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
