// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package match_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/match"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestMatchesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*match.MatchesDataSourceModel)(nil)
	schema := match.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
