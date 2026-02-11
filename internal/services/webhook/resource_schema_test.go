// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/webhook"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestWebhookModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*webhook.WebhookModel)(nil)
	schema := webhook.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
