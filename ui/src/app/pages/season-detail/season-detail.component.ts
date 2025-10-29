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
        this.seasonForm = this.formBuilder.group(
            {
                start_date: [null, Validators.required],
                start_time: [null, Validators.required],
                end_date: [null, Validators.required],
                end_time: [null, Validators.required],
                rounds: [
                    '',
                    [
                        Validators.required,
                        Validators.min(1),
                        Validators.max(52),
                        Validators.pattern(/^(?:[1-9]|[1-4][0-9]|50)$/)
                    ]
                ],
                teams: [[], [Validators.required, this.minMaxArrayValidator(2, this.teams.length)]]
            },
            {
                validators: this.endDateAfterStartDateValidator()
            }
        )
    }

    submitForm(): void {
        if (!this.competitionId) {
            this.notificationService.showWarnAndLog('Form Error', 'No competition selected')
            return
        }

        if (this.seasonForm.invalid) {
            this.seasonForm.markAllAsTouched()
            return
        }

        const { start_date, start_time, end_date, end_time, rounds, teams } = this.seasonForm.value

        const seasonData: Partial<Season> = {
            start_date: this.combineDateAndTime(start_date, start_time),
            end_date: this.combineDateAndTime(end_date, end_time),
            rounds: parseInt(rounds, 10),
            teams
        }

        if (!this.isEditMode) {
            this.createSeason(this.competitionId, seasonData)
        } else if (this.competitionId && this.seasonId) {
            this.updateSeason(this.competitionId, this.seasonId, seasonData)
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
            next: (teams) => {
                this.teams = teams

                // Update teams control validator to reflect the real max
                const teamsControl = this.seasonForm.get('teams')
                if (teamsControl) {
                    teamsControl.setValidators([
                        Validators.required,
                        this.minMaxArrayValidator(2, this.teams.length)
                    ])
                    teamsControl.updateValueAndValidity()
                }
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load teams', err)
            }
        })
    }

    private loadSeason(competitionId: string, id: string): void {
        this.seasonsService.getSeason(competitionId, id).subscribe({
            next: (season) => {
                this.seasonForm.patchValue({
                    start_date: new Date(season.start_date),
                    start_time: new Date(season.start_date),
                    end_date: new Date(season.end_date),
                    end_time: new Date(season.end_date),
                    rounds: season.rounds.toString(),
                    teams: (season.teams as Team[]).map((t) => t.id)
                })
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load season', err)
            }
        })
    }

    private createSeason(competitionId: string, newSeason: Partial<Season>): void {
        this.seasonsService.createSeason(competitionId, newSeason).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Season created successfully')
                this.router.navigate(['/admin/competitions', this.competitionId, 'seasons'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Create Error', 'Failed to create season', err)
            }
        })
    }

    private updateSeason(competitionId: string, id: string, updatedSeason: Partial<Season>): void {
        this.seasonsService.updateSeason(competitionId, id, updatedSeason).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Competition updated successfully')
                this.router.navigate(['/admin/competitions', this.competitionId, 'seasons'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Update Error', 'Failed to update season', err)
            }
        })
    }

    private removeSeason(competitionId: string, id: string): void {
        if (!competitionId || !id) return

        this.seasonsService.deleteSeason(competitionId, id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Season deleted successfully')
                this.router.navigate(['/admin/competitions', this.competitionId, 'seasons'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Delete Error', 'Failed to delete season', err)
            }
        })
    }

    private endDateAfterStartDateValidator() {
        return (group: FormGroup) => {
            const startDate = group.get('start_date')?.value
            const startTime = group.get('start_time')?.value
            const endDate = group.get('end_date')?.value
            const endTime = group.get('end_time')?.value

            if (!startDate || !startTime || !endDate || !endTime) return null

            const start = this.combineDateAndTime(startDate, startTime)
            const end = this.combineDateAndTime(endDate, endTime)

            if (end <= start) {
                group.get('end_date')?.setErrors({ endBeforeStart: true })
                group.get('end_time')?.setErrors({ endBeforeStart: true })
            } else {
                ;['end_date', 'end_time'].forEach((field) => {
                    const errors = group.get(field)?.errors
                    if (errors) {
                        delete errors['endBeforeStart']
                        if (Object.keys(errors).length === 0) group.get(field)?.setErrors(null)
                    }
                })
            }

            return null
        }
    }

    private minMaxArrayValidator(min: number, max: number) {
        return (control: any) => {
            const value = control.value
            if (!Array.isArray(value)) return { invalidArray: true }
            if (value.length < min) return { minArray: true }
            if (value.length > max) return { maxArray: true }
            return null
        }
    }
}
