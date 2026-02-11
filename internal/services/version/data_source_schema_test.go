// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package version_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/version"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestVersionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*version.VersionDataSourceModel)(nil)
	schema := version.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
