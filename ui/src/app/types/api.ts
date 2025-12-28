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
    start_date: Date
    stages: Stage[]
    teams: (string | Team)[]
    updated_at: Date
}

export enum StageType {
    StageTypeRegular = 'regular',
    StageTypeFinals = 'finals'
}

export interface Stage {
    id: string
    name: string
    stageType: StageType
    orderIndex: number
    created_at: Date
    updated_at: Date
    deleted_at?: Date
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

export enum GameStatus {
    SCHEDULED = 'scheduled',
    PLAYING = 'playing',
    FINISHED = 'finished'
}

export interface Game {
    id: string
    season_id: string
    stage_id: string
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
