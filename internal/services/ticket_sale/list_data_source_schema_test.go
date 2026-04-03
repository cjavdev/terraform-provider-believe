// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ticket_sale_test

import (
  "context"
  "testing"

  "github.com/cjavdev/terraform-provider-believe/internal/services/ticket_sale"
  "github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestTicketSalesDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*ticket_sale.TicketSalesDataSourceModel)(nil)
  schema := ticket_sale.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
