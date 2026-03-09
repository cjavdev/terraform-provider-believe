// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_ticket_sale

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ClientTicketSaleResource)(nil)

func (r *ClientTicketSaleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
