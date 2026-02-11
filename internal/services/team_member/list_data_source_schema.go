// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_member

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*TeamMembersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[TeamMembersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for this team membership",
							Computed:    true,
						},
						"character_id": schema.StringAttribute{
							Description: "ID of the character (references /characters/{id})",
							Computed:    true,
						},
						"jersey_number": schema.Int64Attribute{
							Description: "Jersey/shirt number",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
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
						"team_id": schema.StringAttribute{
							Description: "ID of the team they belong to",
							Computed:    true,
						},
						"years_with_team": schema.Int64Attribute{
							Description: "Number of years with the current team",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(0, 50),
							},
						},
						"assists": schema.Int64Attribute{
							Description: "Total assists for the team",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
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
						"certifications": schema.ListAttribute{
							Description: "Coaching certifications and licenses",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"win_rate": schema.Float64Attribute{
							Description: "Career win rate (0.0 to 1.0)",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
						"license_number": schema.StringAttribute{
							Description: "Professional license number",
							Computed:    true,
						},
						"qualifications": schema.ListAttribute{
							Description: "Medical qualifications and degrees",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"is_head_kitman": schema.BoolAttribute{
							Description: "Whether this is the head equipment manager",
							Computed:    true,
						},
						"responsibilities": schema.ListAttribute{
							Description: "List of responsibilities",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *TeamMembersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *TeamMembersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
