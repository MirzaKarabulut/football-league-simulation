# Premier League Simulation

This project simulates a football league involving four teams (Man City, Real Madrid, Barcelona, Bayern Munich) using GoLang and PostgreSQL. The simulation generates random match results and updates the league table over 5 weeks. Users can view the league table in the command line.

## Installation

### Requirements

- Go (version 1.20 or higher)
- PostgreSQL
- Git

### Steps

1. **Clone the Project Repository**
   ```bash
   git clone https://github.com/MirzaKarabulut/football-league-simulation.git
   
   cd football-league-simulation


# Install Dependencies

- go mod tidy

# Run the Application

- go run main.go


# Initialize Teams: (/init)

- To initialize the teams in the league, make a POST request to the /init endpoint. This will add the teams to the database.

# Simulate Matches: (/simulate)

- To simulate the matches and update the league table, make a GET request to the /simulate endpoint. The league table and match results will be displayed in the command line. 


# Predict champion: (/predict-champion-before-last-week)

- After the fourth week, send a GET request to /predict-champion-before-last-week to see predicted champions and their win probabilities.


Before the last week (Week 4), the predicted champions and their probabilities in the command line look like this:

- **Team: Man City**
  - Win Probability: 23.33%
- **Team: Real Madrid**
  - Win Probability: 33.33%
- **Team: Barcelona**
  - Win Probability: 13.33%
- **Team: Bayern Munich**
  - Win Probability: 30.00%


# API Documentation- Swagger: (/docs/index.html)

- GET /docs/index.html : Access the Swagger UI for API documentation.

# SQL Scheme and Queries

- Added SQL scheme and queries.
