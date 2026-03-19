resource "believe_ticket_sale" "example_ticket_sale" {
  buyer_name = "Mae Green"
  currency = "GBP"
  discount = "9.00"
  match_id = "match-001"
  purchase_method = "online"
  quantity = 2
  subtotal = "90.00"
  tax = "16.20"
  total = "97.20"
  unit_price = "45.00"
  buyer_email = "mae.green@example.com"
  coupon_code = "BELIEVE10"
}
