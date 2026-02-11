// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode

import (
	"context"

	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*EpisodesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"character_focus": schema.StringAttribute{
				Description: "Filter by character focus (character ID)",
				Optional:    true,
			},
			"season": schema.Int64Attribute{
				Description: "Filter by season",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3),
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
				CustomType:  customfield.NewNestedObjectListType[EpisodesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier (format: s##e##)",
							Computed:    true,
						},
						"air_date": schema.StringAttribute{
							Description: "Original air date",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"character_focus": schema.ListAttribute{
							Description: "Characters with significant development",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"director": schema.StringAttribute{
							Description: "Episode director",
							Computed:    true,
						},
						"episode_number": schema.Int64Attribute{
							Description: "Episode number within season",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 12),
							},
						},
						"main_theme": schema.StringAttribute{
							Description: "Central theme of the episode",
							Computed:    true,
						},
						"runtime_minutes": schema.Int64Attribute{
							Description: "Episode runtime in minutes",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(20, 60),
							},
						},
						"season": schema.Int64Attribute{
							Description: "Season number",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 3),
							},
						},
						"synopsis": schema.StringAttribute{
							Description: "Brief plot synopsis",
							Computed:    true,
						},
						"ted_wisdom": schema.StringAttribute{
							Description: "Key piece of Ted wisdom from the episode",
							Computed:    true,
						},
						"title": schema.StringAttribute{
							Description: "Episode title",
							Computed:    true,
						},
						"writer": schema.StringAttribute{
							Description: "Episode writer(s)",
							Computed:    true,
						},
						"biscuits_with_boss_moment": schema.StringAttribute{
							Description: "Notable biscuits with the boss scene",
							Computed:    true,
						},
						"memorable_moments": schema.ListAttribute{
							Description: "Standout moments from the episode",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"us_viewers_millions": schema.Float64Attribute{
							Description: "US viewership in millions",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.AtLeast(0),
							},
						},
						"viewer_rating": schema.Float64Attribute{
							Description: "Viewer rating out of 10",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 10),
							},
						},
					},
				},
			},
		},
	}
}

func (d *EpisodesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *EpisodesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
