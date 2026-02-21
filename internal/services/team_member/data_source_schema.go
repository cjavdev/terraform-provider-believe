// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_member

import (
	"context"

	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*TeamMemberDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"member_id": schema.StringAttribute{
				Optional: true,
			},
			"assists": schema.Int64Attribute{
				Description: "Total assists for the team",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
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
			},
			"is_captain": schema.BoolAttribute{
				Description: "Whether this player is team captain",
				Computed:    true,
			},
			"is_head_kitman": schema.BoolAttribute{
				Description: "Whether this is the head equipment manager",
				Computed:    true,
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"member_type": schema.StringAttribute{
						Description: "Filter by member type\nAvailable values: \"player\", \"coach\", \"medical_staff\", \"equipment_manager\".",
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
					"team_id": schema.StringAttribute{
						Description: "Filter by team ID",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *TeamMemberDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TeamMemberDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("member_id"), path.MatchRoot("find_one_by")),
	}
}
