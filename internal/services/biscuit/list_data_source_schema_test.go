// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package biscuit_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/biscuit"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestBiscuitsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*biscuit.BiscuitsDataSourceModel)(nil)
	schema := biscuit.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
