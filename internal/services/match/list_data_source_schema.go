// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package match

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*MatchesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"match_type": schema.StringAttribute{
				Description: "Filter by match type\nAvailable values: \"league\", \"cup\", \"friendly\", \"playoff\", \"final\".",
				Optional:    true,
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
			"result": schema.StringAttribute{
				Description: "Filter by result\nAvailable values: \"win\", \"loss\", \"draw\", \"pending\".",
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
			"team_id": schema.StringAttribute{
				Description: "Filter by team (home or away)",
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
				CustomType:  customfield.NewNestedObjectListType[MatchesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier",
							Computed:    true,
						},
						"away_team_id": schema.StringAttribute{
							Description: "Away team ID",
							Computed:    true,
						},
						"date": schema.StringAttribute{
							Description: "Match date and time",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"home_team_id": schema.StringAttribute{
							Description: "Home team ID",
							Computed:    true,
						},
						"match_type": schema.StringAttribute{
							Description: "Type of match\nAvailable values: \"league\", \"cup\", \"friendly\", \"playoff\", \"final\".",
							Computed:    true,
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
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"away_score": schema.Int64Attribute{
							Description: "Away team score",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"episode_id": schema.StringAttribute{
							Description: "Episode ID where this match is featured",
							Computed:    true,
						},
						"home_score": schema.Int64Attribute{
							Description: "Home team score",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"lesson_learned": schema.StringAttribute{
							Description: "The life lesson learned from this match",
							Computed:    true,
						},
						"possession_percentage": schema.Float64Attribute{
							Description: "Home team possession percentage",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 100),
							},
						},
						"result": schema.StringAttribute{
							Description: "Match result from home team perspective\nAvailable values: \"win\", \"loss\", \"draw\", \"pending\".",
							Computed:    true,
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
							Computed:    true,
						},
						"ticket_revenue_gbp": schema.StringAttribute{
							Description: "Total ticket revenue in GBP",
							Computed:    true,
						},
						"turning_points": schema.ListNestedAttribute{
							Description: "Key moments that changed the match",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[MatchesTurningPointsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										Description: "What happened",
										Computed:    true,
									},
									"emotional_impact": schema.StringAttribute{
										Description: "How this affected the team emotionally",
										Computed:    true,
									},
									"minute": schema.Int64Attribute{
										Description: "Minute of the match",
										Computed:    true,
										Validators: []validator.Int64{
											int64validator.Between(0, 120),
										},
									},
									"character_involved": schema.StringAttribute{
										Description: "Character ID who was central to this moment",
										Computed:    true,
									},
								},
							},
						},
						"weather_temp_celsius": schema.Float64Attribute{
							Description: "Temperature at kickoff in Celsius",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(-30, 50),
							},
						},
					},
				},
			},
		},
	}
}

func (d *MatchesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MatchesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
