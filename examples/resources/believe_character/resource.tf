resource "believe_character" "example_character" {
  background = "Legendary midfielder for Chelsea and AFC Richmond, now assistant coach. Known for his gruff exterior hiding a heart of gold."
  emotional_stats = {
    curiosity = 40
    empathy = 85
    optimism = 45
    resilience = 95
    vulnerability = 60
  }
  name = "Roy Kent"
  personality_traits = ["intense", "loyal", "secretly caring", "profane"]
  role = "coach"
  date_of_birth = "1977-03-15"
  email = "roy.kent@afcrichmond.com"
  growth_arcs = [{
    breakthrough = "Finding purpose beyond playing"
    challenge = "Accepting his body\'s limitations"
    ending_point = "Retired but lost"
    season = 1
    starting_point = "Aging footballer facing retirement"
  }]
  height_meters = 1.78
  profile_image_url = "https://afcrichmond.com/images/roy-kent.jpg"
  salary_gbp = "175000.00"
  signature_quotes = ["He\'s here, he\'s there, he\'s every-f***ing-where, Roy Kent!", "Whistle!"]
  team_id = "afc-richmond"
}
