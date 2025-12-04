CREATE TABLE competitions (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX unique_competition_name_ci 
ON competitions (LOWER(name)) 
WHERE deleted_at IS NULL;

CREATE TABLE seasons (
    id UUID PRIMARY KEY,
    competition_id UUID NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    rounds INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_seasons_competitions FOREIGN KEY (competition_id) REFERENCES competitions(id) ON DELETE CASCADE
);

CREATE TABLE teams (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    abbreviation TEXT NOT NULL,
    location TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE season_teams (
    id UUID PRIMARY KEY,
    season_id UUID NOT NULL,
    team_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_season_teams_season FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE CASCADE,
    CONSTRAINT fk_season_teams_team FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
    UNIQUE (season_id, team_id)
);

CREATE TYPE stage_type AS ENUM ('scheduled', 'playing', 'finished');

CREATE TABLE stages (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  season_id UUID NOT NULL REFERENCES seasons(id),
  name TEXT NOT NULL,
  stage_type TEXT NOT NULL,
  order_index INT NOT NULL
);

CREATE TYPE game_status AS ENUM ('scheduled', 'playing', 'finished');

CREATE TABLE games (
    id UUID PRIMARY KEY,
    season_id UUID NOT NULL,
    stage_id UUID NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    home_team_id UUID NOT NULL,
    away_team_id UUID NOT NULL,
    home_score INTEGER,
    away_score INTEGER,
    status game_status NOT NULL DEFAULT 'scheduled',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_games_seasons FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE CASCADE,
    CONSTRAINT fk_games_stage FOREIGN KEY (stage_id) REFERENCES stages(id) ON DELETE CASCADE,
    CONSTRAINT fk_games_home_team FOREIGN KEY (home_team_id) REFERENCES teams(id) ON DELETE CASCADE,
    CONSTRAINT fk_games_away_team FOREIGN KEY (away_team_id) REFERENCES teams(id) ON DELETE CASCADE
);