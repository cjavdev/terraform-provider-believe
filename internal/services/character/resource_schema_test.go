// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/character"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestCharacterModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*character.CharacterModel)(nil)
	schema := character.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
