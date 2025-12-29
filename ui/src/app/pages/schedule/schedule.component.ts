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
import { Competition, Game, Season, Stage } from '../../types/api'

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

    public dataSource = new MatTableDataSource<Game>([])
    public games: Game[] = []
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

    private initFormListeners(): void {
        this.scheduleForm.get('competition')!.valueChanges.subscribe(this.onCompetitionChange.bind(this))
        this.scheduleForm.get('season')!.valueChanges.subscribe(this.onSeasonChange.bind(this))
        this.scheduleForm.get('stage')!.valueChanges.subscribe(() => {
            this.filterGamesByRound()
        })
    }

    private onCompetitionChange(compId: string): void {
        this.resetSeasons()
        if (compId) {
            this.loadSeasons(compId)
        }
    }

    private onSeasonChange(seasonId: string): void {
        const compId = this.scheduleForm.get('competition')!.value
        const season = this.seasons.find((s) => s.id === seasonId)

        if (!season) return

        this.stages = season.stages || []

        if (this.stages.length > 0) {
            this.scheduleForm.patchValue({ stage: this.stages[0].id }, { emitEvent: false })
        }

        if (compId) {
            this.loadGames(compId, seasonId)
        }
    }

    private resetSeasons(): void {
        this.seasons = []
        this.games = []
        this.dataSource.data = []
        this.scheduleForm.patchValue({ season: '' }, { emitEvent: false })
    }

    private loadCompetitions(): void {
        this.competitionsService.getCompetitions().subscribe({
            next: (competitions) => {
                this.competitions = competitions
                if (competitions.length > 0) {
                    this.scheduleForm.patchValue({ competition: competitions[0].id }, { emitEvent: true })
                }
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load competitions', err)
            }
        })
    }

    private loadSeasons(competitionId: string): void {
        this.seasonsService.getSeasons(competitionId).subscribe({
            next: (seasons) => {
                this.seasons = seasons
                if (seasons.length > 0) {
                    this.scheduleForm.patchValue({ season: seasons[0].id }, { emitEvent: true })
                }
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load seasons', err)
            }
        })
    }

    private loadGames(competitionId: string, seasonId: string): void {
        this.gamesService.getGames(competitionId, seasonId).subscribe({
            next: (games) => {
                this.games = games

                const selectedStageId = this.scheduleForm.get('stage')!.value
                if (!selectedStageId && this.stages.length > 0) {
                    this.scheduleForm.patchValue({ stage: this.stages[0].id }, { emitEvent: false })
                }

                this.filterGamesByRound()
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load games', err)
            }
        })
    }

    private filterGamesByRound(): void {
        const selectedRoundId = this.scheduleForm.get('stage')!.value
        if (selectedRoundId) {
            this.dataSource.data = this.games.filter((game) => game.stage_id === selectedRoundId)
        } else {
            this.dataSource.data = this.games
        }
    }
}
