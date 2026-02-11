// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pep_talk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*PepTalkDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"stream": schema.BoolAttribute{
				Description: "If true, returns SSE stream instead of full response",
				Computed:    true,
				Optional:    true,
			},
			"text": schema.StringAttribute{
				Description: "The full pep talk text",
				Computed:    true,
			},
			"chunks": schema.ListNestedAttribute{
				Description: "Individual chunks of the pep talk",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PepTalkChunksDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"chunk_id": schema.Int64Attribute{
							Description: "Chunk sequence number",
							Computed:    true,
						},
						"is_final": schema.BoolAttribute{
							Description: "Is this the final chunk",
							Computed:    true,
						},
						"text": schema.StringAttribute{
							Description: "The text of this chunk",
							Computed:    true,
						},
						"emotional_beat": schema.StringAttribute{
							Description: "The emotional purpose of this chunk (e.g., greeting, acknowledgment, wisdom, affirmation, encouragement)",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *PepTalkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PepTalkDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
