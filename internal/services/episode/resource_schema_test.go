// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/believe-terraform/internal/services/episode"
	"github.com/stainless-sdks/believe-terraform/internal/test_helpers"
)

func TestEpisodeModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*episode.EpisodeModel)(nil)
	schema := episode.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
