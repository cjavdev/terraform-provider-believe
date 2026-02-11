// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package biscuit_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/biscuit"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestBiscuitDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*biscuit.BiscuitDataSourceModel)(nil)
	schema := biscuit.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
