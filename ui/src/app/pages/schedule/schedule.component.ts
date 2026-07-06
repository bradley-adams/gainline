import { CommonModule } from '@angular/common'
import { Component, OnInit, inject } from '@angular/core'
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms'
import { MatTableDataSource } from '@angular/material/table'
import { RouterModule } from '@angular/router'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { GamesService } from '../../services/games/games.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { MaterialModule } from '../../shared/material/material.module'
import { Competition, Game, Season, Stage, Team } from '../../types/api'

interface GameWithTeamNames extends Game {
    home_team_name: string
    away_team_name: string
}

@Component({
    selector: 'app-schedule',
    standalone: true,
    imports: [CommonModule, RouterModule, MaterialModule, ReactiveFormsModule],
    templateUrl: './schedule.component.html',
    styleUrls: ['./schedule.component.scss']
})
export class ScheduleComponent implements OnInit {
    private readonly competitionsService = inject(CompetitionsService)
    private readonly seasonsService = inject(SeasonsService)
    private readonly gamesService = inject(GamesService)
    private readonly formBuilder = inject(FormBuilder)
    private readonly notificationService = inject(NotificationService)

    public dataSource = new MatTableDataSource<GameWithTeamNames>([])
    public games: GameWithTeamNames[] = []
    public stages: Stage[] = []
    public seasons: Season[] = []
    public competitions: Competition[] = []

    scheduleForm!: FormGroup

    ngOnInit(): void {
        this.scheduleForm = this.formBuilder.group({
            competition: [''],
            season: [''],
            stage: ['']
        })

        this.initFormListeners()
        this.loadCompetitions()
    }

    private get currentSeasonTeams(): Team[] {
        const seasonId = this.scheduleForm.get('season')!.value
        const season = this.seasons.find((s) => s.id === seasonId)
        return (season?.teams ?? []).map((t) => (typeof t === 'string' ? ({ id: t, name: t } as Team) : t))
    }

    private initFormListeners(): void {
        this.scheduleForm.get('competition')!.valueChanges.subscribe(this.onCompetitionChange.bind(this))
        this.scheduleForm.get('season')!.valueChanges.subscribe(this.onSeasonChange.bind(this))
        this.scheduleForm.get('stage')!.valueChanges.subscribe(this.onStageChange.bind(this))
    }

    private onCompetitionChange(compId: string): void {
        this.resetSeasons()
        if (compId) {
            this.loadSeasons(compId)
        }
    }

    private onSeasonChange(seasonId: string): void {
        const season = this.seasons.find((s) => s.id === seasonId)

        if (!season) return

        this.stages = [...(season.stages ?? [])].sort((a, b) => a.order_index - b.order_index)
        this.games = []
        this.dataSource.data = []

        if (this.stages.length > 0) {
            this.scheduleForm.patchValue({ stage: this.stages[0].id }, { emitEvent: true })
        }
    }

    private onStageChange(stageId: string): void {
        const compId = this.scheduleForm.get('competition')!.value
        const seasonId = this.scheduleForm.get('season')!.value

        if (compId && seasonId && stageId) {
            this.loadGames(compId, seasonId, stageId)
        }
    }

    private resetSeasons(): void {
        this.seasons = []
        this.games = []
        this.stages = []
        this.dataSource.data = []
        this.scheduleForm.patchValue({ season: '', stage: '' }, { emitEvent: false })
    }

    private loadCompetitions(): void {
        const allCompetitions: Competition[] = []

        const fetchPage = (page = 1) => {
            this.competitionsService.getCompetitions(page).subscribe({
                next: (response) => {
                    allCompetitions.push(...response.data)
                    if (page < response.pagination.total_pages) {
                        fetchPage(page + 1)
                    } else {
                        this.competitions = allCompetitions
                        if (allCompetitions.length > 0) {
                            this.scheduleForm.patchValue(
                                { competition: allCompetitions[0].id },
                                { emitEvent: true }
                            )
                        }
                    }
                },
                error: (err) => {
                    this.notificationService.showErrorAndLog('Load Error', 'Failed to load competitions', err)
                }
            })
        }

        fetchPage()
    }

    private loadSeasons(competitionId: string): void {
        const allSeasons: Season[] = []

        const fetchPage = (page = 1) => {
            this.seasonsService.getSeasons(competitionId, page).subscribe({
                next: (response) => {
                    allSeasons.push(...response.data)
                    if (page < response.pagination.total_pages) {
                        fetchPage(page + 1)
                    } else {
                        this.seasons = allSeasons
                        if (allSeasons.length > 0) {
                            this.scheduleForm.patchValue({ season: allSeasons[0].id }, { emitEvent: true })
                        }
                    }
                },
                error: (err) => {
                    this.notificationService.showErrorAndLog('Load Error', 'Failed to load seasons', err)
                }
            })
        }

        fetchPage()
    }

    private loadGames(competitionId: string, seasonId: string, stageId: string): void {
        this.gamesService.getGames(competitionId, seasonId, stageId).subscribe({
            next: (games) => {
                const teams = this.currentSeasonTeams
                const mapped: GameWithTeamNames[] = games.map((g) => ({
                    ...g,
                    home_team_name: teams.find((t) => t.id === g.home_team_id)?.name ?? g.home_team_id,
                    away_team_name: teams.find((t) => t.id === g.away_team_id)?.name ?? g.away_team_id
                }))
                this.games = mapped
                this.dataSource.data = mapped
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load games', err)
            }
        })
    }
}
