// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_member

import (
	"context"

	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*TeamMemberResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for this team membership",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"member": schema.SingleNestedAttribute{
				Description: "A football player on the team.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"character_id": schema.StringAttribute{
						Description: "ID of the character (references /characters/{id})",
						Required:    true,
					},
					"jersey_number": schema.Int64Attribute{
						Description: "Jersey/shirt number",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 99),
						},
					},
					"position": schema.StringAttribute{
						Description: "Playing position on the field\nAvailable values: \"goalkeeper\", \"defender\", \"midfielder\", \"forward\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"goalkeeper",
								"defender",
								"midfielder",
								"forward",
							),
						},
					},
					"team_id": schema.StringAttribute{
						Description: "ID of the team they belong to",
						Required:    true,
					},
					"years_with_team": schema.Int64Attribute{
						Description: "Number of years with the current team",
						Required:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 50),
						},
					},
					"assists": schema.Int64Attribute{
						Description: "Total assists for the team",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
						Default: int64default.StaticInt64(0),
					},
					"goals_scored": schema.Int64Attribute{
						Description: "Total goals scored for the team",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
						Default: int64default.StaticInt64(0),
					},
					"is_captain": schema.BoolAttribute{
						Description: "Whether this player is team captain",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"member_type": schema.StringAttribute{
						Description: "Discriminator field indicating this is a player\nAvailable values: \"player\", \"coach\", \"medical_staff\", \"equipment_manager\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"player",
								"coach",
								"medical_staff",
								"equipment_manager",
							),
						},
					},
					"specialty": schema.StringAttribute{
						Description: "Coaching specialty/role\nAvailable values: \"head_coach\", \"assistant_coach\", \"goalkeeping_coach\", \"fitness_coach\", \"tactical_analyst\", \"team_doctor\", \"physiotherapist\", \"sports_psychologist\", \"nutritionist\", \"massage_therapist\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"head_coach",
								"assistant_coach",
								"goalkeeping_coach",
								"fitness_coach",
								"tactical_analyst",
								"team_doctor",
								"physiotherapist",
								"sports_psychologist",
								"nutritionist",
								"massage_therapist",
							),
						},
					},
					"certifications": schema.ListAttribute{
						Description: "Coaching certifications and licenses",
						Optional:    true,
						ElementType: types.StringType,
					},
					"win_rate": schema.Float64Attribute{
						Description: "Career win rate (0.0 to 1.0)",
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"license_number": schema.StringAttribute{
						Description: "Professional license number",
						Optional:    true,
					},
					"qualifications": schema.ListAttribute{
						Description: "Medical qualifications and degrees",
						Optional:    true,
						ElementType: types.StringType,
					},
					"is_head_kitman": schema.BoolAttribute{
						Description: "Whether this is the head equipment manager",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"responsibilities": schema.ListAttribute{
						Description: "List of responsibilities",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"updates": schema.SingleNestedAttribute{
				Description: "Update model for players.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"assists": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"goals_scored": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"is_captain": schema.BoolAttribute{
						Optional: true,
					},
					"jersey_number": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.Between(1, 99),
						},
					},
					"position": schema.StringAttribute{
						Description: "Football positions for players.\nAvailable values: \"goalkeeper\", \"defender\", \"midfielder\", \"forward\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"goalkeeper",
								"defender",
								"midfielder",
								"forward",
							),
						},
					},
					"team_id": schema.StringAttribute{
						Optional: true,
					},
					"years_with_team": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.Between(0, 50),
						},
					},
					"certifications": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"specialty": schema.StringAttribute{
						Description: "Coaching specialties.\nAvailable values: \"head_coach\", \"assistant_coach\", \"goalkeeping_coach\", \"fitness_coach\", \"tactical_analyst\", \"team_doctor\", \"physiotherapist\", \"sports_psychologist\", \"nutritionist\", \"massage_therapist\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"head_coach",
								"assistant_coach",
								"goalkeeping_coach",
								"fitness_coach",
								"tactical_analyst",
								"team_doctor",
								"physiotherapist",
								"sports_psychologist",
								"nutritionist",
								"massage_therapist",
							),
						},
					},
					"win_rate": schema.Float64Attribute{
						Optional: true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
					},
					"license_number": schema.StringAttribute{
						Optional: true,
					},
					"qualifications": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"is_head_kitman": schema.BoolAttribute{
						Optional: true,
					},
					"responsibilities": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"assists": schema.Int64Attribute{
				Description: "Total assists for the team",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default: int64default.StaticInt64(0),
			},
			"character_id": schema.StringAttribute{
				Description: "ID of the character (references /characters/{id})",
				Computed:    true,
			},
			"goals_scored": schema.Int64Attribute{
				Description: "Total goals scored for the team",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
				Default: int64default.StaticInt64(0),
			},
			"is_captain": schema.BoolAttribute{
				Description: "Whether this player is team captain",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_head_kitman": schema.BoolAttribute{
				Description: "Whether this is the head equipment manager",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"jersey_number": schema.Int64Attribute{
				Description: "Jersey/shirt number",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 99),
				},
			},
			"license_number": schema.StringAttribute{
				Description: "Professional license number",
				Computed:    true,
			},
			"member_type": schema.StringAttribute{
				Description: "Discriminator field indicating this is a player\nAvailable values: \"player\", \"coach\", \"medical_staff\", \"equipment_manager\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"player",
						"coach",
						"medical_staff",
						"equipment_manager",
					),
				},
			},
			"position": schema.StringAttribute{
				Description: "Playing position on the field\nAvailable values: \"goalkeeper\", \"defender\", \"midfielder\", \"forward\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"goalkeeper",
						"defender",
						"midfielder",
						"forward",
					),
				},
			},
			"specialty": schema.StringAttribute{
				Description: "Coaching specialty/role\nAvailable values: \"head_coach\", \"assistant_coach\", \"goalkeeping_coach\", \"fitness_coach\", \"tactical_analyst\", \"team_doctor\", \"physiotherapist\", \"sports_psychologist\", \"nutritionist\", \"massage_therapist\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"head_coach",
						"assistant_coach",
						"goalkeeping_coach",
						"fitness_coach",
						"tactical_analyst",
						"team_doctor",
						"physiotherapist",
						"sports_psychologist",
						"nutritionist",
						"massage_therapist",
					),
				},
			},
			"team_id": schema.StringAttribute{
				Description: "ID of the team they belong to",
				Computed:    true,
			},
			"win_rate": schema.Float64Attribute{
				Description: "Career win rate (0.0 to 1.0)",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
			},
			"years_with_team": schema.Int64Attribute{
				Description: "Number of years with the current team",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 50),
				},
			},
			"certifications": schema.ListAttribute{
				Description: "Coaching certifications and licenses",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"qualifications": schema.ListAttribute{
				Description: "Medical qualifications and degrees",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"responsibilities": schema.ListAttribute{
				Description: "List of responsibilities",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (r *TeamMemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TeamMemberResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
