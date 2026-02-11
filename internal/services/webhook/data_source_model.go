// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type WebhookDataSourceModel struct {
	WebhookID   types.String                   `tfsdk:"webhook_id" path:"webhook_id,required"`
	CreatedAt   timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	ID          types.String                   `tfsdk:"id" json:"id,computed"`
	Secret      types.String                   `tfsdk:"secret" json:"secret,computed"`
	URL         types.String                   `tfsdk:"url" json:"url,computed"`
	EventTypes  customfield.List[types.String] `tfsdk:"event_types" json:"event_types,computed"`
}
