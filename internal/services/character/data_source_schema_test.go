// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/character"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestCharacterDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*character.CharacterDataSourceModel)(nil)
	schema := character.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
