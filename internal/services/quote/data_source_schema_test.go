// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/quote"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestQuoteDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*quote.QuoteDataSourceModel)(nil)
	schema := quote.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
