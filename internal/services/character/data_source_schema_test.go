// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/character"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestCharacterDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*character.CharacterDataSourceModel)(nil)
	schema := character.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
