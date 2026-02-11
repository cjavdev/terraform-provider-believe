// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character

import (
	"context"

	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CharactersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
				CustomType:  customfield.NewNestedObjectListType[CharactersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier",
							Computed:    true,
						},
						"background": schema.StringAttribute{
							Description: "Character background and history",
							Computed:    true,
						},
						"emotional_stats": schema.SingleNestedAttribute{
							Description: "Emotional intelligence stats",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[CharactersEmotionalStatsDataSourceModel](ctx),
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
						"name": schema.StringAttribute{
							Description: "Character's full name",
							Computed:    true,
						},
						"personality_traits": schema.ListAttribute{
							Description: "Key personality traits",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
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
						"date_of_birth": schema.StringAttribute{
							Description: "Character's date of birth",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"email": schema.StringAttribute{
							Description: "Character's email address",
							Computed:    true,
						},
						"growth_arcs": schema.ListNestedAttribute{
							Description: "Character development across seasons",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[CharactersGrowthArcsDataSourceModel](ctx),
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
						"height_meters": schema.Float64Attribute{
							Description: "Height in meters",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(1, 2.5),
							},
						},
						"profile_image_url": schema.StringAttribute{
							Description: "URL to character's profile image",
							Computed:    true,
						},
						"salary_gbp": schema.StringAttribute{
							Description: "Annual salary in GBP",
							Computed:    true,
						},
						"signature_quotes": schema.ListAttribute{
							Description: "Memorable quotes from this character",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"team_id": schema.StringAttribute{
							Description: "ID of the team they belong to",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *CharactersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CharactersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
