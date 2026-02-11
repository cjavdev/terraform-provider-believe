// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/quote"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestQuoteModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*quote.QuoteModel)(nil)
	schema := quote.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
