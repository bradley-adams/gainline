import { CommonModule } from '@angular/common'
import { Component, inject, OnInit } from '@angular/core'
import { MatPaginator, PageEvent } from '@angular/material/paginator'
import { MatTableDataSource } from '@angular/material/table'
import { ActivatedRoute, RouterModule } from '@angular/router'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { GamesService } from '../../services/games/games.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { MaterialModule } from '../../shared/material/material.module'
import { Game, Season, Team } from '../../types/api'

type GameWithNames = Game & {
    home_team_name: string
    away_team_name: string
}

@Component({
    selector: 'app-game-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule, BreadcrumbComponent, MatPaginator],
    templateUrl: './game-list.component.html',
    styleUrl: './game-list.component.scss'
})
export class GameListComponent implements OnInit {
    private readonly route = inject(ActivatedRoute)
    private readonly gamesService = inject(GamesService)
    private readonly seasonsService = inject(SeasonsService)
    private readonly notificationService = inject(NotificationService)

    dataSource = new MatTableDataSource<GameWithNames>([])

    public competitionId: string | null = null
    public seasonId: string | null = null
    public teams: Team[] = []

    public total = 0
    public page = 1
    public pageSize = 10

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')
        this.seasonId = this.route.snapshot.paramMap.get('season-id')

        if (this.competitionId && this.seasonId) {
            this.loadSeasonAndGames(this.competitionId, this.seasonId)
        }
    }

    public onPageChange(event: PageEvent): void {
        this.page = event.pageIndex + 1
        this.pageSize = event.pageSize

        if (this.competitionId && this.seasonId) {
            this.loadGames(this.competitionId, this.seasonId)
        }
    }

    private loadSeasonAndGames(competitionId: string, seasonId: string): void {
        this.seasonsService.getSeason(competitionId, seasonId).subscribe({
            next: (season: Season) => {
                this.teams = season.teams.map((t) =>
                    typeof t === 'string' ? ({ id: t, name: t } as Team) : t
                )
                this.loadGames(competitionId, seasonId)
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load season', err)
            }
        })
    }

    private loadGames(competitionId: string, seasonId: string): void {
        this.gamesService.getGamesPaginated(competitionId, seasonId, this.page, this.pageSize).subscribe({
            next: (response) => {
                this.total = response.pagination.total

                const gamesWithNames = response.data.map((game) => ({
                    ...game,
                    home_team_name:
                        this.teams.find((t) => t.id === game.home_team_id)?.name || game.home_team_id,
                    away_team_name:
                        this.teams.find((t) => t.id === game.away_team_id)?.name || game.away_team_id
                }))
                this.dataSource.data = gamesWithNames
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load games', err)
            }
        })
    }

    confirmDelete(game: Game): void {
        const confirmationMessage = `Are you sure you want to delete this game?`

        this.notificationService
            .showConfirm('Confirm Delete', confirmationMessage)
            .afterClosed()
            .subscribe((confirmed) => {
                if (confirmed && this.competitionId && this.seasonId) {
                    this.removeGame(this.competitionId, this.seasonId, game.id)
                }
            })
    }

    private removeGame(competitionId: string, seasonId: string, id: string): void {
        this.gamesService.deleteGame(competitionId, seasonId, id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Game deleted successfully')
                this.loadGames(competitionId, seasonId)
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Delete Error', 'Failed to delete game', err)
            }
        })
    }
}
