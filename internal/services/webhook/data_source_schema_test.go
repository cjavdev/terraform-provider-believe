// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/webhook"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestWebhookDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*webhook.WebhookDataSourceModel)(nil)
	schema := webhook.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
