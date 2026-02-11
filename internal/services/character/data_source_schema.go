// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CharacterDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"character_id": schema.StringAttribute{
				Optional: true,
			},
			"background": schema.StringAttribute{
				Description: "Character background and history",
				Computed:    true,
			},
			"date_of_birth": schema.StringAttribute{
				Description: "Character's date of birth",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"email": schema.StringAttribute{
				Description: "Character's email address",
				Computed:    true,
			},
			"height_meters": schema.Float64Attribute{
				Description: "Height in meters",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(1, 2.5),
				},
			},
			"name": schema.StringAttribute{
				Description: "Character's full name",
				Computed:    true,
			},
			"profile_image_url": schema.StringAttribute{
				Description: "URL to character's profile image",
				Computed:    true,
			},
			"role": schema.StringAttribute{
				Description: "Character's role\nAvailable values: \"coach\", \"player\", \"owner\", \"manager\", \"staff\", \"journalist\", \"family\", \"friend\", \"fan\", \"other\".",
				Computed:    true,
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
			"salary_gbp": schema.StringAttribute{
				Description: "Annual salary in GBP",
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "ID of the team they belong to",
				Computed:    true,
			},
			"personality_traits": schema.ListAttribute{
				Description: "Key personality traits",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"signature_quotes": schema.ListAttribute{
				Description: "Memorable quotes from this character",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"emotional_stats": schema.SingleNestedAttribute{
				Description: "Emotional intelligence stats",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CharacterEmotionalStatsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"curiosity": schema.Int64Attribute{
						Description: "Level of curiosity over judgment (0-100)",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"empathy": schema.Int64Attribute{
						Description: "Capacity for empathy (0-100)",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"optimism": schema.Int64Attribute{
						Description: "Level of optimism (0-100)",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"resilience": schema.Int64Attribute{
						Description: "Bounce-back ability (0-100)",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"vulnerability": schema.Int64Attribute{
						Description: "Willingness to be vulnerable (0-100)",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
				},
			},
			"growth_arcs": schema.ListNestedAttribute{
				Description: "Character development across seasons",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CharacterGrowthArcsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"breakthrough": schema.StringAttribute{
							Description: "Key breakthrough moment",
							Computed:    true,
						},
						"challenge": schema.StringAttribute{
							Description: "Main challenge faced",
							Computed:    true,
						},
						"ending_point": schema.StringAttribute{
							Description: "Where the character ends up",
							Computed:    true,
						},
						"season": schema.Int64Attribute{
							Description: "Season number",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 3),
							},
						},
						"starting_point": schema.StringAttribute{
							Description: "Where the character starts emotionally",
							Computed:    true,
						},
					},
				},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"min_optimism": schema.Int64Attribute{
						Description: "Minimum optimism score",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
					"role": schema.StringAttribute{
						Description: "Filter by role\nAvailable values: \"coach\", \"player\", \"owner\", \"manager\", \"staff\", \"journalist\", \"family\", \"friend\", \"fan\", \"other\".",
						Optional:    true,
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
					"team_id": schema.StringAttribute{
						Description: "Filter by team ID",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *CharacterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CharacterDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("character_id"), path.MatchRoot("find_one_by")),
	}
}
