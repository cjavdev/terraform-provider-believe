// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package match_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/match"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestMatchModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*match.MatchModel)(nil)
	schema := match.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
