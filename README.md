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
   git clone https://github.com/yourusername/football-league-simulation.git
   
   cd football-league-simulation


# Install Dependencies

- go mod tidy


# Create the Database and Table

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    power INT NOT NULL,
    pts INT NOT NULL,
    w INT NOT NULL,
    d INT NOT NULL,
    l INT NOT NULL,
    gd INT NOT NULL
);

# Run the Application

- go run main.go


# Initialize Teams

- To initialize the teams in the league, make a POST request to the /init endpoint. This will add the teams to the database.

# Simulate Matches

- To simulate the matches and update the league table, make a POST request to the /simulate endpoint. The league table will be displayed in the command line.

# View League Table

- After simulating the matches, the league table will be printed in the command line. It will display the teams, points (PTS), wins (W), draws (D), losses (L), and goal difference (GD).

League table in the commend line look like this:

| Team           | PTS | W  | D  | L  | GD  |
|----------------|-----|----|----|----|-----|
| Real Madrid    | 30  | 10 | 0  | 5  | 196 |
| Bayern Munich  | 27  | 9  | 0  | 6  | -26 |
| Man City       | 21  | 7  | 0  | 8  | 77  |
| Barcelona      | 12  | 4  | 0  | 11 | -247|


# Get Match Results for Week 5

- **URL**: `get-match-results/:week`
- **Method**: GET
- **Description**: Retrieves match results for weeks.

# Predict champion:

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
