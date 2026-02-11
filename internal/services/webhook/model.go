// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/apijson"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

type WebhookModel struct {
	WebhookID   types.String                                  `tfsdk:"webhook_id" path:"webhook_id,optional"`
	URL         types.String                                  `tfsdk:"url" json:"url,required"`
	Description types.String                                  `tfsdk:"description" json:"description,optional"`
	EventTypes  *[]types.String                               `tfsdk:"event_types" json:"event_types,optional"`
	CreatedAt   timetypes.RFC3339                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID          types.String                                  `tfsdk:"id" json:"id,computed"`
	Message     types.String                                  `tfsdk:"message" json:"message,computed,no_refresh"`
	Secret      types.String                                  `tfsdk:"secret" json:"secret,computed"`
	TedSays     types.String                                  `tfsdk:"ted_says" json:"ted_says,computed,no_refresh"`
	Webhook     customfield.NestedObject[WebhookWebhookModel] `tfsdk:"webhook" json:"webhook,computed,no_refresh"`
}

func (m WebhookModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WebhookModel) MarshalJSONForUpdate(state WebhookModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WebhookWebhookModel struct {
	ID          types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	EventTypes  customfield.List[types.String] `tfsdk:"event_types" json:"event_types,computed"`
	Secret      types.String                   `tfsdk:"secret" json:"secret,computed"`
	URL         types.String                   `tfsdk:"url" json:"url,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
}
