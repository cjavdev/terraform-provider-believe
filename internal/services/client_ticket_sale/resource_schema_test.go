// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_ticket_sale_test

import (
	"context"
	"testing"

	"github.com/cjavdev/terraform-provider-believe/internal/services/client_ticket_sale"
	"github.com/cjavdev/terraform-provider-believe/internal/test_helpers"
)

func TestClientTicketSaleModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*client_ticket_sale.ClientTicketSaleModel)(nil)
	schema := client_ticket_sale.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
