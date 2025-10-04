import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { MaterialModule } from '../../shared/material/material.module'
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { TeamsService } from '../../services/teams/teams.service'
import { Season, Team } from '../../types/api'
import { MatInputModule } from '@angular/material/input'
import { MatFormFieldModule } from '@angular/material/form-field'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatNativeDateModule } from '@angular/material/core'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { NotificationService } from '../../services/notifications/notifications.service'

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
        BreadcrumbComponent
    ],
    templateUrl: './season-detail.component.html',
    styleUrl: './season-detail.component.scss'
})
export class SeasonDetailComponent {
    private readonly formBuilder = inject(FormBuilder)
    private readonly route = inject(ActivatedRoute)
    private readonly router = inject(Router)
    private readonly seasonsService = inject(SeasonsService)
    private readonly teamsService = inject(TeamsService)
    private readonly notificationService = inject(NotificationService)

    public seasonForm!: FormGroup
    public competitionId: string | null = null
    private seasonId: string | null = null
    public isEditMode = false

    public teams: Team[] = []

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')
        this.seasonId = this.route.snapshot.paramMap.get('season-id')
        this.isEditMode = !!this.seasonId

        this.seasonForm = this.formBuilder.group({
            start_date: ['', Validators.required],
            end_date: ['', Validators.required],
            rounds: ['', [Validators.required, Validators.min(1), Validators.max(50)]],
            teams: [[]]
        })

        this.loadTeams()

        if (this.isEditMode && this.competitionId && this.seasonId) {
            this.loadSeason(this.competitionId, this.seasonId)
        }
    }

    submitForm(): void {
        if (this.seasonForm.invalid || !this.competitionId) {
            console.error('season form is invalid')
            return
        }

        const seasonData: Season = this.seasonForm.value
        if (!this.isEditMode) {
            this.createSeason(this.competitionId, seasonData)
        } else if (this.competitionId && this.seasonId) {
            this.updateSeason(this.competitionId, this.seasonId, seasonData)
        }
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
            },
            error: (err) => {
                console.error('Error loading teams:', err)
                this.notificationService.showError('Load Error', 'Failed to load teams')
            }
        })
    }

    private loadSeason(competitionId: string, id: string): void {
        this.seasonsService.getSeason(competitionId, id).subscribe({
            next: (season) => {
                this.seasonForm.patchValue({
                    start_date: season.start_date,
                    end_date: season.end_date,
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

    private createSeason(competitionId: string, newSeason: Season): void {
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

    private updateSeason(competitionId: string, id: string, updatedSeason: Season): void {
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
