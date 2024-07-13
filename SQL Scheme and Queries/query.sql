-- Added teams to the teams table

INSERT INTO teams (name, power, pts, w, d, l, gd) VALUES ('Man City', 90, 0, 0, 0, 0, 0);
INSERT INTO teams (name, power, pts, w, d, l, gd) VALUES ('Real Madrid', 85, 0, 0, 0, 0, 0);
INSERT INTO teams (name, power, pts, w, d, l, gd) VALUES ('Barcelona', 80, 0, 0, 0, 0, 0);
INSERT INTO teams (name, power, pts, w, d, l, gd) VALUES ('Bayern Munich', 88, 0, 0, 0, 0, 0);

-- Updated the teams table

UPDATE teams SET pts = 3, w = 1, d = 0, l = 0, gd = 2 WHERE id = 1;
UPDATE teams SET pts = 0, w = 0, d = 0, l = 1, gd = -2 WHERE id = 2;

-- Select the teams table

SELECT id, name, power FROM teams;
SELECT id, name, pts, w, d, l, gd FROM teams;

-- Select the match_results table

SELECT name, pts, w, d, l, gd FROM teams ORDER BY pts DESC, gd DESC;
