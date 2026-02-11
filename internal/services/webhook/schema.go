// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*WebhookResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"webhook_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"url": schema.StringAttribute{
				Description:   "The URL to send webhook events to",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description:   "Optional description for this webhook",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"event_types": schema.ListAttribute{
				Description: "List of event types to subscribe to. If not provided, subscribes to all events.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("match.completed", "team_member.transferred"),
					),
				},
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"created_at": schema.StringAttribute{
				Description: "When the webhook was registered",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Unique webhook identifier",
				Computed:    true,
			},
			"message": schema.StringAttribute{
				Description: "Status message",
				Computed:    true,
				Default:     stringdefault.StaticString("Webhook registered successfully"),
			},
			"secret": schema.StringAttribute{
				Description: "The secret key for verifying webhook signatures (base64 encoded)",
				Computed:    true,
			},
			"ted_says": schema.StringAttribute{
				Description: "Ted's reaction",
				Computed:    true,
				Default:     stringdefault.StaticString("You know what? Staying connected is what it's all about. Welcome to the team!"),
			},
			"webhook": schema.SingleNestedAttribute{
				Description: "The registered webhook details",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[WebhookWebhookModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique webhook identifier",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Description: "When the webhook was registered",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"event_types": schema.ListAttribute{
						Description: "List of event types this webhook is subscribed to",
						Computed:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive("match.completed", "team_member.transferred"),
							),
						},
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"secret": schema.StringAttribute{
						Description: "The secret key for verifying webhook signatures (base64 encoded)",
						Computed:    true,
					},
					"url": schema.StringAttribute{
						Description: "The URL to send webhook events to",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description for this webhook",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *WebhookResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WebhookResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
