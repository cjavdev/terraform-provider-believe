// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

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

var _ datasource.DataSourceWithConfigValidators = (*TeamsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
				CustomType:  customfield.NewNestedObjectListType[TeamsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier",
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
						"stadium": schema.StringAttribute{
							Description: "Home stadium name",
							Computed:    true,
						},
						"values": schema.SingleNestedAttribute{
							Description: "Team's core values",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[TeamsValuesDataSourceModel](ctx),
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
						"is_active": schema.BoolAttribute{
							Description: "Whether the team is currently active",
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
						"rival_teams": schema.ListAttribute{
							Description: "List of rival team IDs",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"secondary_color": schema.StringAttribute{
							Description: "Secondary team color (hex)",
							Computed:    true,
						},
						"stadium_location": schema.SingleNestedAttribute{
							Description: "Geographic coordinates for a location.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[TeamsStadiumLocationDataSourceModel](ctx),
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
					},
				},
			},
		},
	}
}

func (d *TeamsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *TeamsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
