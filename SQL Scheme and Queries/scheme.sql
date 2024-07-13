-- Create teams table

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

-- Create match_results table

CREATE TABLE match_results (
    id SERIAL PRIMARY KEY,
    week INT NOT NULL,
    team_a VARCHAR(255) NOT NULL,
    team_b VARCHAR(255) NOT NULL,
    score_a INT NOT NULL,
    score_b INT NOT NULL
);
