// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/team"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestTeamsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*team.TeamsDataSourceModel)(nil)
	schema := team.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
