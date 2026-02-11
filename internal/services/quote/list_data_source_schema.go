// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*QuotesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"character_id": schema.StringAttribute{
				Description: "Filter by character",
				Optional:    true,
			},
			"funny": schema.BoolAttribute{
				Description: "Filter funny quotes",
				Optional:    true,
			},
			"inspirational": schema.BoolAttribute{
				Description: "Filter inspirational quotes",
				Optional:    true,
			},
			"moment_type": schema.StringAttribute{
				Description: "Filter by moment type\nAvailable values: \"halftime_speech\", \"press_conference\", \"locker_room\", \"training\", \"biscuits_with_boss\", \"pub\", \"one_on_one\", \"celebration\", \"crisis\", \"casual\", \"confrontation\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"halftime_speech",
						"press_conference",
						"locker_room",
						"training",
						"biscuits_with_boss",
						"pub",
						"one_on_one",
						"celebration",
						"crisis",
						"casual",
						"confrontation",
					),
				},
			},
			"theme": schema.StringAttribute{
				Description: "Filter by theme\nAvailable values: \"belief\", \"teamwork\", \"curiosity\", \"kindness\", \"resilience\", \"vulnerability\", \"growth\", \"humor\", \"wisdom\", \"leadership\", \"love\", \"forgiveness\", \"philosophy\", \"romance\", \"cultural-pride\", \"cultural-differences\", \"antagonism\", \"celebration\", \"identity\", \"isolation\", \"power\", \"sacrifice\", \"standards\", \"confidence\", \"conflict\", \"honesty\", \"integrity\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"belief",
						"teamwork",
						"curiosity",
						"kindness",
						"resilience",
						"vulnerability",
						"growth",
						"humor",
						"wisdom",
						"leadership",
						"love",
						"forgiveness",
						"philosophy",
						"romance",
						"cultural-pride",
						"cultural-differences",
						"antagonism",
						"celebration",
						"identity",
						"isolation",
						"power",
						"sacrifice",
						"standards",
						"confidence",
						"conflict",
						"honesty",
						"integrity",
					),
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
				CustomType:  customfield.NewNestedObjectListType[QuotesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier",
							Computed:    true,
						},
						"character_id": schema.StringAttribute{
							Description: "ID of the character who said it",
							Computed:    true,
						},
						"context": schema.StringAttribute{
							Description: "Context in which the quote was said",
							Computed:    true,
						},
						"moment_type": schema.StringAttribute{
							Description: "Type of moment when the quote was said\nAvailable values: \"halftime_speech\", \"press_conference\", \"locker_room\", \"training\", \"biscuits_with_boss\", \"pub\", \"one_on_one\", \"celebration\", \"crisis\", \"casual\", \"confrontation\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"halftime_speech",
									"press_conference",
									"locker_room",
									"training",
									"biscuits_with_boss",
									"pub",
									"one_on_one",
									"celebration",
									"crisis",
									"casual",
									"confrontation",
								),
							},
						},
						"text": schema.StringAttribute{
							Description: "The quote text",
							Computed:    true,
						},
						"theme": schema.StringAttribute{
							Description: "Primary theme of the quote\nAvailable values: \"belief\", \"teamwork\", \"curiosity\", \"kindness\", \"resilience\", \"vulnerability\", \"growth\", \"humor\", \"wisdom\", \"leadership\", \"love\", \"forgiveness\", \"philosophy\", \"romance\", \"cultural-pride\", \"cultural-differences\", \"antagonism\", \"celebration\", \"identity\", \"isolation\", \"power\", \"sacrifice\", \"standards\", \"confidence\", \"conflict\", \"honesty\", \"integrity\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"belief",
									"teamwork",
									"curiosity",
									"kindness",
									"resilience",
									"vulnerability",
									"growth",
									"humor",
									"wisdom",
									"leadership",
									"love",
									"forgiveness",
									"philosophy",
									"romance",
									"cultural-pride",
									"cultural-differences",
									"antagonism",
									"celebration",
									"identity",
									"isolation",
									"power",
									"sacrifice",
									"standards",
									"confidence",
									"conflict",
									"honesty",
									"integrity",
								),
							},
						},
						"episode_id": schema.StringAttribute{
							Description: "Episode where the quote appears",
							Computed:    true,
						},
						"is_funny": schema.BoolAttribute{
							Description: "Whether this quote is humorous",
							Computed:    true,
						},
						"is_inspirational": schema.BoolAttribute{
							Description: "Whether this quote is inspirational",
							Computed:    true,
						},
						"popularity_score": schema.Float64Attribute{
							Description: "Popularity/virality score (0-100)",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 100),
							},
						},
						"secondary_themes": schema.ListAttribute{
							Description: "Additional themes",
							Computed:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"belief",
										"teamwork",
										"curiosity",
										"kindness",
										"resilience",
										"vulnerability",
										"growth",
										"humor",
										"wisdom",
										"leadership",
										"love",
										"forgiveness",
										"philosophy",
										"romance",
										"cultural-pride",
										"cultural-differences",
										"antagonism",
										"celebration",
										"identity",
										"isolation",
										"power",
										"sacrifice",
										"standards",
										"confidence",
										"conflict",
										"honesty",
										"integrity",
									),
								),
							},
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"times_shared": schema.Int64Attribute{
							Description: "Number of times shared on social media",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},
		},
	}
}

func (d *QuotesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *QuotesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
