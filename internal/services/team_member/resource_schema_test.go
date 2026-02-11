// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_member_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/team_member"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestTeamMemberModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*team_member.TeamMemberModel)(nil)
	schema := team_member.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
