import { CommonModule } from '@angular/common'
import { Component, inject, OnInit } from '@angular/core'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { MatInputModule } from '@angular/material/input'
import { MatFormFieldModule } from '@angular/material/form-field'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatNativeDateModule } from '@angular/material/core'
import { MatTimepickerModule } from '@angular/material/timepicker'

import { MaterialModule } from '../../shared/material/material.module'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { TeamsService } from '../../services/teams/teams.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { Season, Team } from '../../types/api'

@Component({
    selector: 'app-season-detail',
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
    templateUrl: './season-detail.component.html',
    styleUrl: './season-detail.component.scss'
})
export class SeasonDetailComponent implements OnInit {
    private readonly formBuilder = inject(FormBuilder)
    private readonly route = inject(ActivatedRoute)
    private readonly router = inject(Router)
    private readonly seasonsService = inject(SeasonsService)
    private readonly teamsService = inject(TeamsService)
    private readonly notificationService = inject(NotificationService)

    seasonForm!: FormGroup
    competitionId: string | null = null
    private seasonId: string | null = null
    isEditMode = false
    teams: Team[] = []

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')
        this.seasonId = this.route.snapshot.paramMap.get('season-id')
        this.isEditMode = !!this.seasonId

        this.initForm()
        this.loadTeams()

        if (this.isEditMode && this.competitionId && this.seasonId) {
            this.loadSeason(this.competitionId, this.seasonId)
        }
    }

    private initForm(): void {
        this.seasonForm = this.formBuilder.group({
            start_date: [null, Validators.required],
            start_time: [null, Validators.required],
            end_date: [null, Validators.required],
            end_time: [null, Validators.required],
            rounds: ['', [Validators.required, Validators.min(1), Validators.max(50)]],
            teams: [[]]
        })
    }

    submitForm(): void {
        if (this.seasonForm.invalid || !this.competitionId) {
            console.error('Season form is invalid')
            return
        }

        const { start_date, start_time, end_date, end_time, rounds, teams } = this.seasonForm.value

        const seasonData: Partial<Season> = {
            start_date: this.combineDateAndTime(start_date, start_time),
            end_date: this.combineDateAndTime(end_date, end_time),
            rounds,
            teams
        }

        if (!this.isEditMode) {
            this.createSeason(this.competitionId, seasonData)
        } else if (this.competitionId && this.seasonId) {
            this.updateSeason(this.competitionId, this.seasonId, seasonData)
        }
    }

    private combineDateAndTime(date: Date | string, time: Date | string): Date {
        const d = new Date(date) // local Date from form
        let hours: number, minutes: number

        if (time instanceof Date) {
            hours = time.getHours()
            minutes = time.getMinutes()
        } else {
            ;[hours, minutes] = time.split(':').map(Number)
        }

        // Create a new Date in UTC
        const utcDate = new Date(Date.UTC(d.getFullYear(), d.getMonth(), d.getDate(), hours, minutes, 0, 0))

        return utcDate
    }

    confirmDelete(): void {
        const confirmationMessage = `Are you sure you want to delete this season?`

        this.notificationService
            .showConfirm('Confirm Delete', confirmationMessage)
            .afterClosed()
            .subscribe((confirmed) => {
                if (confirmed && this.competitionId && this.seasonId) {
                    this.removeSeason(this.competitionId, this.seasonId)
                }
            })
    }

    private loadTeams(): void {
        this.teamsService.getTeams().subscribe({
            next: (teams) => (this.teams = teams),
            error: (err) => {
                console.error('Error loading teams:', err)
                this.notificationService.showError('Load Error', 'Failed to load teams')
            }
        })
    }

    private loadSeason(competitionId: string, id: string): void {
        this.seasonsService.getSeason(competitionId, id).subscribe({
            next: (season) => {
                const startUtc = new Date(season.start_date)
                const endUtc = new Date(season.end_date)

                this.seasonForm.patchValue({
                    start_date: startUtc,
                    start_time: startUtc,
                    end_date: endUtc,
                    end_time: endUtc,
                    rounds: season.rounds,
                    teams: (season.teams as Team[]).map((t) => t.id)
                })
            },
            error: (err) => {
                console.error('Error loading season:', err)
                this.notificationService.showError('Load Error', 'Failed to load season')
            }
        })
    }

    private createSeason(competitionId: string, newSeason: Partial<Season>): void {
        this.seasonsService.createSeason(competitionId, newSeason).subscribe({
            next: () => {
                this.router.navigate(['/admin/competitions', this.competitionId, 'seasons'])
            },
            error: (err) => {
                console.error('Error creating season:', err)
                this.notificationService.showError('Create Error', 'Failed to create season')
            }
        })
    }

    private updateSeason(competitionId: string, id: string, updatedSeason: Partial<Season>): void {
        this.seasonsService.updateSeason(competitionId, id, updatedSeason).subscribe({
            next: () => {
                this.router.navigate(['/admin/competitions', this.competitionId, 'seasons'])
            },
            error: (err) => {
                console.error('Error updating season:', err)
                this.notificationService.showError('Update Error', 'Failed to update season')
            }
        })
    }

    private removeSeason(competitionId: string, id: string): void {
        if (!competitionId || !id) return

        this.seasonsService.deleteSeason(competitionId, id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Season deleted successfully', 'OK')
                this.router.navigate(['/admin/competitions', this.competitionId, 'seasons'])
            },
            error: (err) => {
                console.error('Error deleting season:', err)
                this.notificationService.showError('Delete Error', 'Failed to delete season')
            }
        })
    }
}
