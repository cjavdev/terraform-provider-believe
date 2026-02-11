// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*QuoteResource)(nil)

func (r *QuoteResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
