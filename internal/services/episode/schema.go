// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*EpisodeResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier (format: s##e##)",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"air_date": schema.StringAttribute{
				Description: "Original air date",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"director": schema.StringAttribute{
				Description: "Episode director",
				Required:    true,
			},
			"episode_number": schema.Int64Attribute{
				Description: "Episode number within season",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 12),
				},
			},
			"main_theme": schema.StringAttribute{
				Description: "Central theme of the episode",
				Required:    true,
			},
			"runtime_minutes": schema.Int64Attribute{
				Description: "Episode runtime in minutes",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(20, 60),
				},
			},
			"season": schema.Int64Attribute{
				Description: "Season number",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3),
				},
			},
			"synopsis": schema.StringAttribute{
				Description: "Brief plot synopsis",
				Required:    true,
			},
			"ted_wisdom": schema.StringAttribute{
				Description: "Key piece of Ted wisdom from the episode",
				Required:    true,
			},
			"title": schema.StringAttribute{
				Description: "Episode title",
				Required:    true,
			},
			"writer": schema.StringAttribute{
				Description: "Episode writer(s)",
				Required:    true,
			},
			"character_focus": schema.ListAttribute{
				Description: "Characters with significant development",
				Required:    true,
				ElementType: types.StringType,
			},
			"biscuits_with_boss_moment": schema.StringAttribute{
				Description: "Notable biscuits with the boss scene",
				Optional:    true,
			},
			"us_viewers_millions": schema.Float64Attribute{
				Description: "US viewership in millions",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"viewer_rating": schema.Float64Attribute{
				Description: "Viewer rating out of 10",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 10),
				},
			},
			"memorable_moments": schema.ListAttribute{
				Description: "Standout moments from the episode",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *EpisodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EpisodeResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
