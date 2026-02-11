resource "believe_team_member" "example_team_member" {
  member = {
    character_id = "jamie-tartt"
    jersey_number = 9
    position = "forward"
    team_id = "afc-richmond"
    years_with_team = 3
    assists = 23
    goals_scored = 47
    is_captain = false
    member_type = "player"
  }
}
