resource "believe_team" "example_team" {
  culture_score = 70
  founded_year = 1895
  league = "Premier League"
  name = "West Ham United"
  stadium = "London Stadium"
  values = {
    primary_value = "Pride"
    secondary_values = ["History", "Community", "Passion"]
    team_motto = "Forever Blowing Bubbles"
  }
  annual_budget_gbp = "150000000.00"
  average_attendance = 59988
  contact_email = "info@westhamunited.co.uk"
  is_active = true
  nickname = "The Hammers"
  primary_color = "#7A263A"
  rival_teams = ["afc-richmond", "tottenham"]
  secondary_color = "#1BB1E7"
  stadium_location = {
    latitude = 51.5387
    longitude = -0.0166
  }
  website = "https://www.whufc.com"
  win_percentage = 52.3
}
