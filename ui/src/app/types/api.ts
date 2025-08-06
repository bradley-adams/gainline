export enum GameStatus {
    SCHEDULED = 'scheduled',
    PLAYING = 'playing',
    FINISHED = 'finished'
}

export interface Game {
    id: string
    season_id: string
    round: number
    date: Date
    home_team_id: string
    away_team_id: string
    home_score: number
    away_score: number
    status: GameStatus
    created_at: Date
    updated_at: Date
    deleted_at?: Date
}

export interface Competition {
    id: string
    name: string
    created_at: Date
    updated_at: Date
    deleted_at?: Date
}

export interface Season {
    competition_id: string
    created_at: Date
    deleted_at?: Date
    end_date: Date
    id: string
    rounds: number
    start_date: Date
    teams: Team[]
    updated_at: Date
}

export interface Team {
    abbreviation: string
    created_at: Date
    deleted_at?: Date
    id: string
    location: string
    name: string
    updated_at: Date
}