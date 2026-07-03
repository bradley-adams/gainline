import { CommonModule } from '@angular/common'
import { Component, OnDestroy, OnInit, inject } from '@angular/core'
import { ActivatedRoute } from '@angular/router'
import { Subscription } from 'rxjs'
import { GamesService } from '../../services/games/games.service'
import { LiveGameService } from '../../services/live-game/live-game.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { MaterialModule } from '../../shared/material/material.module'
import { Game, GameState, GameStatus, Team } from '../../types/api'

@Component({
    selector: 'app-schedule-game',
    standalone: true,
    imports: [CommonModule, MaterialModule],
    templateUrl: './schedule-game.component.html',
    styleUrls: ['./schedule-game.component.scss']
})
export class ScheduleGameComponent implements OnInit, OnDestroy {
    private readonly route = inject(ActivatedRoute)
    private readonly gamesService = inject(GamesService)
    private readonly seasonsService = inject(SeasonsService)
    private readonly liveGameService = inject(LiveGameService)
    private readonly notificationService = inject(NotificationService)

    public game: Game | null = null
    public liveState: GameState | null = null
    public competitionId: string | null = null
    public seasonId: string | null = null
    public gameId: string | null = null

    public homeTeamName: string | null = null
    public awayTeamName: string | null = null

    private teams: Team[] = []
    private liveSubscription: Subscription | null = null

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')
        this.seasonId = this.route.snapshot.paramMap.get('season-id')
        this.gameId = this.route.snapshot.paramMap.get('game-id')

        if (this.competitionId && this.seasonId && this.gameId) {
            this.loadSeason(this.competitionId, this.seasonId)
            this.loadGame(this.competitionId, this.seasonId, this.gameId)
        }
    }

    ngOnDestroy(): void {
        this.liveSubscription?.unsubscribe()
    }

    public get homeScore(): number {
        return this.liveState?.home_score ?? this.game?.home_score ?? 0
    }

    public get awayScore(): number {
        return this.liveState?.away_score ?? this.game?.away_score ?? 0
    }

    public get status(): string {
        return this.liveState?.status ?? this.game?.status ?? ''
    }

    private loadSeason(competitionId: string, seasonId: string): void {
        this.seasonsService.getSeason(competitionId, seasonId).subscribe({
            next: (season) => {
                this.teams = season.teams.map((t) =>
                    typeof t === 'string' ? ({ id: t, name: t } as Team) : t
                )
                this.resolveTeamNames()
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load season', err)
            }
        })
    }

    private loadGame(competitionId: string, seasonId: string, gameId: string): void {
        this.gamesService.getGame(competitionId, seasonId, gameId).subscribe({
            next: (game) => {
                this.game = game
                this.resolveTeamNames()
                if (game.status === GameStatus.PLAYING) {
                    this.watchLiveState(competitionId, seasonId, gameId)
                }
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load game', err)
            }
        })
    }

    private resolveTeamNames(): void {
        if (!this.game || !this.teams.length) return
        this.homeTeamName =
            this.teams.find((t) => t.id === this.game!.home_team_id)?.name ?? this.game.home_team_id
        this.awayTeamName =
            this.teams.find((t) => t.id === this.game!.away_team_id)?.name ?? this.game.away_team_id
    }

    private watchLiveState(competitionId: string, seasonId: string, gameId: string): void {
        this.liveSubscription = this.liveGameService.watchGame(competitionId, seasonId, gameId).subscribe({
            next: (state) => {
                this.liveState = state
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Live Error', 'Lost live connection', err)
            }
        })
    }
}
