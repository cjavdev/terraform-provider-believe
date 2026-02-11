// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_member_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/team_member"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestTeamMemberDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*team_member.TeamMemberDataSourceModel)(nil)
	schema := team_member.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
