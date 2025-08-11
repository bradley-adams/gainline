import { Component, OnInit, inject } from '@angular/core'
import { CommonModule } from '@angular/common'
import { RouterModule } from '@angular/router'
import { MatTableDataSource } from '@angular/material/table'
import { MaterialModule } from '../../shared/material/material.module'
import { Game, Competition, Season } from '../../types/api'
import { GamesService } from '../../services/games/games.service'
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { CompetitionsService } from '../../services/competitions/competitions.service'

@Component({
    selector: 'app-schedule',
    standalone: true,
    imports: [CommonModule, RouterModule, MaterialModule, ReactiveFormsModule],
    templateUrl: './schedule.component.html',
    styleUrls: ['./schedule.component.scss']
})
export class ScheduleComponent implements OnInit {
    private readonly gamesService = inject(GamesService)
    private readonly seasonsService = inject(SeasonsService)
    private readonly competitionsService = inject(CompetitionsService)
    private readonly formBuilder = inject(FormBuilder)

    public dataSource = new MatTableDataSource<Game>([])
    public games: Game[] = []
    public seasons: Season[] = []
    public competitions: Competition[] = []
    public rounds: number[] = []

    scheduleForm!: FormGroup

    ngOnInit(): void {
        this.scheduleForm = this.formBuilder.group({
            competition: [''],
            season: [''],
            round: ['']
        })

        this.initFormListeners()
        this.loadCompetitions()
    }

    private initFormListeners(): void {
        this.scheduleForm
            .get('competition')!
            .valueChanges.subscribe(this.onCompetitionChange.bind(this))
        this.scheduleForm.get('season')!.valueChanges.subscribe(this.onSeasonChange.bind(this))
        this.scheduleForm.get('round')!.valueChanges.subscribe(this.onRoundChange.bind(this))
    }

    private onCompetitionChange(compId: string): void {
        this.resetSeasonsAndRounds()
        if (compId) {
            this.loadSeasons(compId)
        }
    }

    private onSeasonChange(seasonId: string): void {
        const compId = this.scheduleForm.get('competition')!.value
        const season = this.seasons.find((s) => s.id === seasonId)

        if (!season) {
            this.rounds = []
            this.scheduleForm.patchValue({ round: '' }, { emitEvent: false })
            return
        }

        if (season.rounds > 0) {
            this.rounds = Array.from({ length: season.rounds }, (_, i) => i + 1)
            this.scheduleForm.patchValue({ round: 1 }, { emitEvent: true })
        }

        if (compId) {
            this.loadGames(compId, seasonId, 1)
        }
    }

    private onRoundChange(round: number): void {
        const compId = this.scheduleForm.get('competition')!.value
        const seasonId = this.scheduleForm.get('season')!.value
        if (compId && seasonId && round) {
            this.loadGames(compId, seasonId, round)
        }
    }

    private resetSeasonsAndRounds(): void {
        this.seasons = []
        this.rounds = []
        this.games = []
        this.dataSource.data = []
        this.scheduleForm.patchValue({ season: '', round: '' }, { emitEvent: false })
    }

    private loadCompetitions(): void {
        this.competitionsService.getCompetitions().subscribe({
            next: (competitions) => {
                this.competitions = competitions
                if (competitions.length > 0) {
                    this.scheduleForm.patchValue(
                        { competition: competitions[0].id },
                        { emitEvent: true }
                    )
                }
            },
            error: (err) => console.error('Error loading competitions:', err)
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
            error: (err) => console.error('Error loading seasons:', err)
        })
    }

    private loadGames(competitionId: string, seasonId: string, round: number): void {
        this.gamesService.getGames(competitionId, seasonId).subscribe({
            next: (games) => {
                this.games = games.filter((g) => g.round === round)
                this.dataSource.data = this.games
            },
            error: (err) => console.error('Error loading games:', err)
        })
    }
}
