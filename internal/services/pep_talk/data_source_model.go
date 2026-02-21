// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pep_talk

import (
	"context"

	"github.com/cjavdev/believe-go"
	"github.com/cjavdev/believe-go/packages/param"
	"github.com/cjavdev/terraform-provider-believe/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PepTalkDataSourceModel struct {
	Stream types.Bool                                                 `tfsdk:"stream" query:"stream,computed_optional"`
	Text   types.String                                               `tfsdk:"text" json:"text,computed"`
	Chunks customfield.NestedObjectList[PepTalkChunksDataSourceModel] `tfsdk:"chunks" json:"chunks,computed"`
}

func (m *PepTalkDataSourceModel) toReadParams(_ context.Context) (params believe.PepTalkGetParams, diags diag.Diagnostics) {
	params = believe.PepTalkGetParams{}

	if !m.Stream.IsNull() {
		params.Stream = param.NewOpt(m.Stream.ValueBool())
	}

	return
}

type PepTalkChunksDataSourceModel struct {
	ChunkID       types.Int64  `tfsdk:"chunk_id" json:"chunk_id,computed"`
	IsFinal       types.Bool   `tfsdk:"is_final" json:"is_final,computed"`
	Text          types.String `tfsdk:"text" json:"text,computed"`
	EmotionalBeat types.String `tfsdk:"emotional_beat" json:"emotional_beat,computed"`
}
