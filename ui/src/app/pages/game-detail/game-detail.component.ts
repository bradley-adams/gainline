import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import {
    AbstractControl,
    FormBuilder,
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    ValidationErrors,
    Validators
} from '@angular/forms'
import { MatNativeDateModule } from '@angular/material/core'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatFormFieldModule } from '@angular/material/form-field'
import { MatInputModule } from '@angular/material/input'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { MaterialModule } from '../../shared/material/material.module'

import { MatTimepickerModule } from '@angular/material/timepicker'
import { combineLatest, map } from 'rxjs'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { GamesService } from '../../services/games/games.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { Game, Season, Team } from '../../types/api'

@Component({
    selector: 'app-game-detail',
    standalone: true,
    imports: [
        CommonModule,
        RouterModule,
        MaterialModule,
        ReactiveFormsModule,
        MatDatepickerModule,
        MatNativeDateModule,
        MatFormFieldModule,
        MatInputModule,
        MatTimepickerModule,
        BreadcrumbComponent
    ],
    templateUrl: './game-detail.component.html',
    styleUrls: ['./game-detail.component.scss']
})
export class GameDetailComponent {
    private readonly fb = inject(FormBuilder)
    private readonly route = inject(ActivatedRoute)
    private readonly router = inject(Router)
    private readonly gamesService = inject(GamesService)
    private readonly seasonsService = inject(SeasonsService)
    private readonly notificationService = inject(NotificationService)

    public gameForm!: FormGroup
    public dateControl = new FormControl<Date | null>(null, Validators.required)
    public timeControl = new FormControl<Date | null>(null, Validators.required)
    public isEditMode = false
    public seasonId: string | null = null
    public competitionId: string | null = null
    public gameId: string | null = null
    public teams: Team[] = []
    public rounds: number[] = []

    public seasonStart?: Date
    public seasonEnd?: Date

    ngOnInit(): void {
        const params = this.route.snapshot.paramMap
        this.competitionId = params.get('competition-id')
        this.seasonId = params.get('season-id')
        this.gameId = params.get('game-id')

        this.isEditMode = !!this.gameId
        this.initForm()
        this.initDateTimeSync()

        if (this.competitionId && this.seasonId) {
            this.loadSeason(this.competitionId, this.seasonId)
        }

        if (this.isEditMode && this.competitionId && this.seasonId && this.gameId) {
            this.loadGame(this.competitionId, this.seasonId, this.gameId)
        }
    }

    private initForm(): void {
        const baseScoreValidators = [Validators.pattern(/^[0-9]+$/), Validators.min(0)]

        this.gameForm = this.fb.group(
            {
                round: [null, Validators.required],
                datetime: [null, Validators.required],
                home_team_id: [null, Validators.required],
                away_team_id: [null, Validators.required],
                home_score: [null, baseScoreValidators],
                away_score: [null, baseScoreValidators],
                status: ['scheduled', Validators.required]
            },
            { validators: [this.dateWithinSeasonValidator, this.teamsMustDifferValidator] }
        )

        this.gameForm.get('status')?.valueChanges.subscribe((status) => {
            const homeScoreCtrl = this.gameForm.get('home_score')
            const awayScoreCtrl = this.gameForm.get('away_score')

            if (status === 'playing' || status === 'finished') {
                // Scores ARE required
                homeScoreCtrl?.setValidators([...baseScoreValidators, Validators.required])
                awayScoreCtrl?.setValidators([...baseScoreValidators, Validators.required])
            } else if (status === 'scheduled') {
                // Scores must NOT exist
                homeScoreCtrl?.setValidators(baseScoreValidators)
                awayScoreCtrl?.setValidators(baseScoreValidators)
                homeScoreCtrl?.setValue(null)
                awayScoreCtrl?.setValue(null)
            } else if (status === 'cancelled') {
                // Scores must NOT exist (same as scheduled)
                homeScoreCtrl?.setValidators(baseScoreValidators)
                awayScoreCtrl?.setValidators(baseScoreValidators)
                homeScoreCtrl?.setValue(null)
                awayScoreCtrl?.setValue(null)
            }

            homeScoreCtrl?.updateValueAndValidity()
            awayScoreCtrl?.updateValueAndValidity()
        })
    }

    private initDateTimeSync(): void {
        combineLatest([this.dateControl.valueChanges, this.timeControl.valueChanges])
            .pipe(
                map(([date, time]) => {
                    if (date && time) {
                        const combined = new Date(date)
                        combined.setHours(time.getHours(), time.getMinutes(), 0)
                        return combined
                    }
                    return null
                })
            )
            .subscribe((combinedDateTime) => {
                this.gameForm.get('datetime')?.setValue(combinedDateTime)
            })
    }

    private dateWithinSeasonValidator = (control: AbstractControl): ValidationErrors | null => {
        const datetime = control.get('datetime')?.value
        const { seasonStart, seasonEnd } = this

        if (!datetime || !seasonStart || !seasonEnd) return null
        return datetime < seasonStart || datetime > seasonEnd ? { outOfSeason: true } : null
    }

    private teamsMustDifferValidator = (group: AbstractControl): ValidationErrors | null => {
        const homeTeamId = group.get('home_team_id')?.value
        const awayTeamId = group.get('away_team_id')?.value

        if (homeTeamId && awayTeamId && homeTeamId === awayTeamId) {
            return { teamsMustDiffer: true }
        }

        return null
    }

    submitForm(): void {
        if (!this.competitionId || !this.seasonId) {
            console.error('game form is invalid')
            this.notificationService.showWarnAndLog(
                'Form Error',
                'Game form is invalid or competition/season not selected'
            )
            return
        }

        if (this.gameForm.invalid) {
            this.gameForm.markAllAsTouched()
            return
        }

        const { datetime, home_score, away_score, ...rest } = this.gameForm.value

        const parseScore = (score: any) => (score != null && score !== '' ? parseInt(score, 10) : null)

        const gameData: Game = {
            ...rest,
            season_id: this.seasonId,
            date: datetime,
            home_score: parseScore(home_score),
            away_score: parseScore(away_score)
        }

        this.isEditMode
            ? this.updateGame(this.competitionId, this.seasonId, this.gameId!, gameData)
            : this.createGame(this.competitionId, this.seasonId, gameData)
    }

    confirmDelete(): void {
        const confirmationMessage = `Are you sure you want to delete this game?`

        this.notificationService
            .showConfirm('Confirm Delete', confirmationMessage)
            .afterClosed()
            .subscribe((confirmed) => {
                if (confirmed && this.competitionId && this.seasonId && this.gameId) {
                    this.removeGame(this.competitionId, this.seasonId, this.gameId)
                }
            })
    }

    private loadSeason(competitionId: string, id: string): void {
        this.seasonsService.getSeason(competitionId, id).subscribe({
            next: (season: Season) => {
                this.teams = season.teams.map((t) =>
                    typeof t === 'string' ? ({ id: t, name: t } as Team) : t
                )
                this.rounds = Array.from({ length: season.rounds }, (_, i) => i + 1)

                this.seasonStart = new Date(season.start_date)
                this.seasonEnd = new Date(season.end_date)

                this.gameForm.updateValueAndValidity()
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load season', err)
            }
        })
    }

    private loadGame(competitionId: string, seasonId: string, gameId: string): void {
        this.gamesService.getGame(competitionId, seasonId, gameId).subscribe({
            next: (game) => {
                const gameDate = new Date(game.date)
                this.dateControl.setValue(gameDate)
                this.timeControl.setValue(gameDate)
                this.gameForm.patchValue({
                    ...game,
                    datetime: gameDate
                })
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load game', err)
            }
        })
    }

    private createGame(competitionId: string, seasonId: string, newGame: Game): void {
        this.gamesService.createGame(competitionId, seasonId, newGame).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Game created successfully')
                this.router.navigate(['/admin/competitions', competitionId, 'seasons', seasonId, 'games'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Create Error', 'Failed to create game', err)
            }
        })
    }

    private updateGame(competitionId: string, seasonId: string, gameId: string, updatedGame: Game): void {
        this.gamesService.updateGame(competitionId, seasonId, gameId, updatedGame).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Game updated successfully')
                this.router.navigate(['/admin/competitions', competitionId, 'seasons', seasonId, 'games'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Update Error', 'Failed to update game', err)
            }
        })
    }

    private removeGame(competitionId: string, seasonId: string, id: string): void {
        if (!competitionId || !seasonId || !id) return

        this.gamesService.deleteGame(competitionId, seasonId, id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Game deleted successfully')
                this.router.navigate([
                    '/admin/competitions',
                    this.competitionId,
                    'seasons',
                    this.seasonId,
                    'games'
                ])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Delete Error', 'Failed to delete game', err)
            }
        })
    }
}
