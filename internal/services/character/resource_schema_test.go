// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package character_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/character"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestCharacterModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*character.CharacterModel)(nil)
	schema := character.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
