// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/episode"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestEpisodesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*episode.EpisodesDataSourceModel)(nil)
	schema := episode.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
