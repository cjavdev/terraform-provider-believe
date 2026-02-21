// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package coaching_principle_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/coaching_principle"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestCoachingPrincipleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*coaching_principle.CoachingPrincipleDataSourceModel)(nil)
	schema := coaching_principle.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
