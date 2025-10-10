import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatNativeDateModule } from '@angular/material/core'
import { MatFormFieldModule } from '@angular/material/form-field'
import { MatInputModule } from '@angular/material/input'
import { MaterialModule } from '../../shared/material/material.module'

import { GamesService } from '../../services/games/games.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { Season, Team, Game } from '../../types/api'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { NotificationService } from '../../services/notifications/notifications.service'
import { MatTimepickerModule } from '@angular/material/timepicker'

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
    public isEditMode = false
    public seasonId: string | null = null
    public competitionId: string | null = null
    public gameId: string | null = null
    public teams: Team[] = []
    public rounds: number[] = []

    ngOnInit(): void {
        const params = this.route.snapshot.paramMap
        this.competitionId = params.get('competition-id')
        this.seasonId = params.get('season-id')
        this.gameId = params.get('game-id')

        this.isEditMode = !!this.gameId
        this.initForm()

        if (this.competitionId && this.seasonId) {
            this.loadSeason(this.competitionId, this.seasonId)
        }

        if (this.isEditMode && this.competitionId && this.seasonId && this.gameId) {
            this.loadGame(this.competitionId, this.seasonId, this.gameId)
        }
    }

    private initForm(): void {
        this.gameForm = this.fb.group({
            round: [null, Validators.required],
            date: [null, Validators.required],
            time: [null, Validators.required],
            home_team_id: [null, Validators.required],
            away_team_id: [null, Validators.required],
            home_score: [null],
            away_score: [null],
            status: ['scheduled', Validators.required]
        })

        this.gameForm.get('status')?.valueChanges.subscribe((status) => {
            if (status === 'playing' || status === 'finished') {
                this.gameForm.get('home_score')?.setValidators([Validators.required])
                this.gameForm.get('away_score')?.setValidators([Validators.required])
            } else {
                this.gameForm.get('home_score')?.clearValidators()
                this.gameForm.get('away_score')?.clearValidators()
            }
            this.gameForm.get('home_score')?.updateValueAndValidity()
            this.gameForm.get('away_score')?.updateValueAndValidity()
        })
    }

    submitForm(): void {
        if (this.gameForm.invalid || !this.competitionId || !this.seasonId) {
            console.error('game form is invalid')
            return
        }

        const { date, time, ...rest } = this.gameForm.value
        const gameData: Game = {
            ...rest,
            season_id: this.seasonId,
            date: this.combineDateAndTime(date, time)
        }

        if (!this.isEditMode) {
            this.createGame(this.competitionId, this.seasonId, gameData)
        } else if (this.seasonId && this.gameId) {
            this.updateGame(this.competitionId, this.seasonId, this.gameId, gameData)
        }
    }

    private combineDateAndTime(date: Date | string, time: Date | string): Date {
        const d = new Date(date)
        let hours: number, minutes: number

        if (time instanceof Date) {
            hours = time.getHours()
            minutes = time.getMinutes()
        } else {
            ;[hours, minutes] = time.split(':').map(Number)
        }

        d.setHours(hours, minutes, 0, 0)
        return d
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
            },
            error: (err) => {
                console.error('Error loading season:', err)
                this.notificationService.showError('Load Error', 'Failed to load season')
            }
        })
    }

    private loadGame(competitionId: string, seasonId: string, gameId: string): void {
        this.gamesService.getGame(competitionId, seasonId, gameId).subscribe({
            next: (game) => {
                this.gameForm.patchValue({
                    ...game,
                    date: new Date(game.date),
                    time: new Date(game.date)
                })
            },
            error: (err) => {
                console.error('Error loading game:', err)
                this.notificationService.showError('Load Error', 'Failed to load game')
            }
        })
    }

    private createGame(competitionId: string, seasonId: string, newGame: Game): void {
        this.gamesService.createGame(competitionId, seasonId, newGame).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Game created successfully', 'OK')
                this.router.navigate(['/admin/competitions', competitionId, 'seasons', seasonId, 'games'])
            },
            error: (err) => {
                console.error('Error creating game:', err)
                this.notificationService.showError('Create Error', 'Failed to create game')
            }
        })
    }

    private updateGame(competitionId: string, seasonId: string, gameId: string, updatedGame: Game): void {
        this.gamesService.updateGame(competitionId, seasonId, gameId, updatedGame).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Game updated successfully', 'OK')
                this.router.navigate(['/admin/competitions', competitionId, 'seasons', seasonId, 'games'])
            },
            error: (err) => {
                console.error('Error updating game:', err)
                this.notificationService.showError('Update Error', 'Failed to update game')
            }
        })
    }

    private removeGame(competitionId: string, seasonId: string, id: string): void {
        if (!competitionId || !seasonId || !id) return

        this.gamesService.deleteGame(competitionId, seasonId, id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Game deleted successfully', 'OK')
                this.router.navigate([
                    '/admin/competitions',
                    this.competitionId,
                    'seasons',
                    this.seasonId,
                    'games'
                ])
            },
            error: (err) => {
                console.error('Error deleting game:', err)
                this.notificationService.showError('Delete Error', 'Failed to delete game')
            }
        })
    }
}
