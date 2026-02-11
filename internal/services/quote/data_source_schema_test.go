// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package quote_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/quote"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestQuoteDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*quote.QuoteDataSourceModel)(nil)
	schema := quote.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
