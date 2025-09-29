import { Component, inject, OnInit } from '@angular/core'
import { GamesService } from '../../services/games/games.service'
import { MatTableDataSource } from '@angular/material/table'
import { Game, Season, Team } from '../../types/api'
import { CommonModule } from '@angular/common'
import { MaterialModule } from '../../shared/material/material.module'
import { ActivatedRoute, RouterModule } from '@angular/router'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { NotificationService } from '../../services/notifications/notifications.service'

@Component({
    selector: 'app-game-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule, BreadcrumbComponent],
    templateUrl: './game-list.component.html',
    styleUrl: './game-list.component.scss'
})
export class GameListComponent implements OnInit {
    private readonly route = inject(ActivatedRoute)
    private readonly gamesService = inject(GamesService)
    private readonly seasonsService = inject(SeasonsService)
    private readonly notificationService = inject(NotificationService)

    dataSource = new MatTableDataSource<Game>()
    public competitionId: string | null = null
    public seasonId: string | null = null
    public teams: Team[] = []

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')
        this.seasonId = this.route.snapshot.paramMap.get('season-id')

        if (this.competitionId && this.seasonId) {
            this.loadSeasonAndGames(this.competitionId, this.seasonId)
        }
    }

    private loadSeasonAndGames(competitionId: string, seasonId: string): void {
        this.seasonsService.getSeason(competitionId, seasonId).subscribe({
            next: (season: Season) => {
                // Store only the teams for this season
                this.teams = season.teams.map((t) =>
                    typeof t === 'string' ? ({ id: t, name: t } as Team) : t
                )
                this.loadGames(competitionId, seasonId)
            },
            error: (err) => {
                console.error('Error loading season:', err)
                this.notificationService.showError('Load Error', 'Failed to load season')
            }
        })
    }

    private loadGames(competitionId: string, seasonId: string): void {
        this.gamesService.getGames(competitionId, seasonId).subscribe({
            next: (games) => {
                // Precompute team names for each game
                const gamesWithNames = games.map((game) => ({
                    ...game,
                    home_team_name:
                        this.teams.find((t) => t.id === game.home_team_id)?.name || game.home_team_id,
                    away_team_name:
                        this.teams.find((t) => t.id === game.away_team_id)?.name || game.away_team_id
                }))
                this.dataSource.data = gamesWithNames
            },
            error: (err) => console.error('Error loading games:', err)
        })
    }
}
