// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pep_talk_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/pep_talk"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestPepTalkDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pep_talk.PepTalkDataSourceModel)(nil)
	schema := pep_talk.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
