// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

import (
	"context"

	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/cjavdev/terraform-provider-believe/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.ResourceWithConfigValidators = (*TeamResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"culture_score": schema.Int64Attribute{
				Description: "Team culture/morale score (0-100)",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"founded_year": schema.Int64Attribute{
				Description: "Year the club was founded",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(1800, 2030),
				},
			},
			"league": schema.StringAttribute{
				Description: "Current league\nAvailable values: \"Premier League\", \"Championship\", \"League One\", \"League Two\", \"La Liga\", \"Serie A\", \"Bundesliga\", \"Ligue 1\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"Premier League",
						"Championship",
						"League One",
						"League Two",
						"La Liga",
						"Serie A",
						"Bundesliga",
						"Ligue 1",
					),
				},
			},
			"name": schema.StringAttribute{
				Description: "Team name",
				Required:    true,
			},
			"stadium": schema.StringAttribute{
				Description: "Home stadium name",
				Required:    true,
			},
			"values": schema.SingleNestedAttribute{
				Description: "Team's core values",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"primary_value": schema.StringAttribute{
						Description: "The team's primary guiding value",
						Required:    true,
					},
					"secondary_values": schema.ListAttribute{
						Description: "Supporting values",
						Required:    true,
						ElementType: types.StringType,
					},
					"team_motto": schema.StringAttribute{
						Description: "Team's motivational motto",
						Required:    true,
					},
				},
			},
			"average_attendance": schema.Float64Attribute{
				Description: "Average match attendance",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"contact_email": schema.StringAttribute{
				Description: "Team contact email",
				Optional:    true,
			},
			"nickname": schema.StringAttribute{
				Description: "Team nickname",
				Optional:    true,
			},
			"primary_color": schema.StringAttribute{
				Description: "Primary team color (hex)",
				Optional:    true,
			},
			"secondary_color": schema.StringAttribute{
				Description: "Secondary team color (hex)",
				Optional:    true,
			},
			"website": schema.StringAttribute{
				Description: "Official team website",
				Optional:    true,
			},
			"win_percentage": schema.Float64Attribute{
				Description: "Season win percentage",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 100),
				},
			},
			"rival_teams": schema.ListAttribute{
				Description: "List of rival team IDs",
				Optional:    true,
				ElementType: types.StringType,
			},
			"stadium_location": schema.SingleNestedAttribute{
				Description: "Geographic coordinates for a location.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"latitude": schema.Float64Attribute{
						Description: "Latitude in degrees",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(-90, 90),
						},
					},
					"longitude": schema.Float64Attribute{
						Description: "Longitude in degrees",
						Required:    true,
						Validators: []validator.Float64{
							float64validator.Between(-180, 180),
						},
					},
				},
			},
			"annual_budget_gbp": schema.DynamicAttribute{
				Description: "Annual budget in GBP",
				Optional:    true,
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.Float64Type{}, basetypes.StringType{}),
				},
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the team is currently active",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
		},
	}
}

func (r *TeamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TeamResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
