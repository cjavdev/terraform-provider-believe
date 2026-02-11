// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/team"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestTeamModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*team.TeamModel)(nil)
	schema := team.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
