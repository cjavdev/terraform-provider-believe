resource "believe_webhook" "example_webhook" {
  url = "https://example.com/webhooks"
  description = "Production webhook for match notifications"
  event_types = ["match.completed", "team_member.transferred"]
}
