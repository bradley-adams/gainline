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
        MatInputModule
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

    public gameForm!: FormGroup
    public isEditMode = false
    public seasonId: string | null = null
    public competitionId: string | null = null
    public gameId: string | null = null
    public teams: Team[] = []
    public rounds: number[] = []

    ngOnInit(): void {
        this.route.paramMap.subscribe((params) => {
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
        })
    }

    private initForm(): void {
        this.gameForm = this.fb.group({
            round: [null, Validators.required],
            date: [null, Validators.required],
            home_team_id: [null, Validators.required],
            away_team_id: [null, Validators.required],
            home_score: [0],
            away_score: [0],
            status: ['scheduled', Validators.required]
        })
    }

    submitForm(): void {
        if (this.gameForm.invalid || !this.competitionId || !this.seasonId) {
            console.error('game form is invalid')
            return
        }

        const gameData: Game = this.gameForm.value
        gameData.season_id = this.seasonId

        if (!this.isEditMode) {
            this.createGame(this.competitionId, this.seasonId, gameData)
        } else if (this.seasonId && this.gameId) {
            this.updateGame(this.competitionId, this.seasonId, this.gameId, gameData)
        }
    }

    private loadSeason(competitionId: string, id: string): void {
        this.seasonsService.getSeason(competitionId, id).subscribe({
            next: (season: Season) => {
                this.teams = season.teams.map((t) =>
                    typeof t === 'string' ? ({ id: t, name: t } as Team) : t
                )
                this.rounds = Array.from({ length: season.rounds }, (_, i) => i + 1)
            },
            error: (err) => console.error('Error loading season:', err)
        })
    }

    private loadGame(competitionId: string, seasonId: string, gameId: string): void {
        this.gamesService.getGame(competitionId, seasonId, gameId).subscribe({
            next: (game) => {
                this.gameForm.patchValue(game)
            },
            error: (err) => console.error('Error loading game:', err)
        })
    }

    private createGame(competitionId: string, seasonId: string, newGame: Game): void {
        this.gamesService.createGame(competitionId, seasonId, newGame).subscribe({
            next: () => {
                this.router.navigate(['/admin/competitions', competitionId, 'seasons', seasonId, 'games'])
            },
            error: (err) => console.error('Error creating game:', err)
        })
    }

    private updateGame(competitionId: string, seasonId: string, gameId: string, updatedGame: Game): void {
        this.gamesService.updateGame(competitionId, seasonId, gameId, updatedGame).subscribe({
            next: () => {
                this.router.navigate(['/admin/competitions', competitionId, 'seasons', seasonId, 'games'])
            },
            error: (err) => console.error('Error updating game:', err)
        })
    }
}
