// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

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

var _ datasource.DataSourceWithConfigValidators = (*TeamDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"team_id": schema.StringAttribute{
				Optional: true,
			},
			"annual_budget_gbp": schema.StringAttribute{
				Description: "Annual budget in GBP",
				Computed:    true,
			},
			"average_attendance": schema.Float64Attribute{
				Description: "Average match attendance",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"contact_email": schema.StringAttribute{
				Description: "Team contact email",
				Computed:    true,
			},
			"culture_score": schema.Int64Attribute{
				Description: "Team culture/morale score (0-100)",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			"founded_year": schema.Int64Attribute{
				Description: "Year the club was founded",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1800, 2030),
				},
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the team is currently active",
				Computed:    true,
			},
			"league": schema.StringAttribute{
				Description: "Current league\nAvailable values: \"Premier League\", \"Championship\", \"League One\", \"League Two\", \"La Liga\", \"Serie A\", \"Bundesliga\", \"Ligue 1\".",
				Computed:    true,
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
				Computed:    true,
			},
			"nickname": schema.StringAttribute{
				Description: "Team nickname",
				Computed:    true,
			},
			"primary_color": schema.StringAttribute{
				Description: "Primary team color (hex)",
				Computed:    true,
			},
			"secondary_color": schema.StringAttribute{
				Description: "Secondary team color (hex)",
				Computed:    true,
			},
			"stadium": schema.StringAttribute{
				Description: "Home stadium name",
				Computed:    true,
			},
			"website": schema.StringAttribute{
				Description: "Official team website",
				Computed:    true,
			},
			"win_percentage": schema.Float64Attribute{
				Description: "Season win percentage",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 100),
				},
			},
			"rival_teams": schema.ListAttribute{
				Description: "List of rival team IDs",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"stadium_location": schema.SingleNestedAttribute{
				Description: "Geographic coordinates for a location.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TeamStadiumLocationDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"latitude": schema.Float64Attribute{
						Description: "Latitude in degrees",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(-90, 90),
						},
					},
					"longitude": schema.Float64Attribute{
						Description: "Longitude in degrees",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(-180, 180),
						},
					},
				},
			},
			"values": schema.SingleNestedAttribute{
				Description: "Team's core values",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TeamValuesDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"primary_value": schema.StringAttribute{
						Description: "The team's primary guiding value",
						Computed:    true,
					},
					"secondary_values": schema.ListAttribute{
						Description: "Supporting values",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"team_motto": schema.StringAttribute{
						Description: "Team's motivational motto",
						Computed:    true,
					},
				},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"league": schema.StringAttribute{
						Description: "Filter by league\nAvailable values: \"Premier League\", \"Championship\", \"League One\", \"League Two\", \"La Liga\", \"Serie A\", \"Bundesliga\", \"Ligue 1\".",
						Optional:    true,
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
					"min_culture_score": schema.Int64Attribute{
						Description: "Minimum culture score",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(0, 100),
						},
					},
				},
			},
		},
	}
}

func (d *TeamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TeamDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("team_id"), path.MatchRoot("find_one_by")),
	}
}
