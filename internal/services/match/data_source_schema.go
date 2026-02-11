// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package match

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
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*MatchDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"match_id": schema.StringAttribute{
				Optional: true,
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
			"away_team_id": schema.StringAttribute{
				Description: "Away team ID",
				Computed:    true,
			},
			"date": schema.StringAttribute{
				Description: "Match date and time",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
			"home_team_id": schema.StringAttribute{
				Description: "Home team ID",
				Computed:    true,
			},
			"lesson_learned": schema.StringAttribute{
				Description: "The life lesson learned from this match",
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
			"weather_temp_celsius": schema.Float64Attribute{
				Description: "Temperature at kickoff in Celsius",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(-30, 50),
				},
			},
			"turning_points": schema.ListNestedAttribute{
				Description: "Key moments that changed the match",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MatchTurningPointsDataSourceModel](ctx),
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
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
				},
			},
		},
	}
}

func (d *MatchDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MatchDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("match_id"), path.MatchRoot("find_one_by")),
	}
}
