import { Component, inject } from '@angular/core'
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { Competition, Team } from '../../types/api'
import { CommonModule } from '@angular/common'
import { MaterialModule } from '../../shared/material/material.module'
import { TeamsService } from '../../services/teams/teams.service'
import { NotificationService } from '../../services/notifications/notifications.service'

@Component({
    selector: 'app-team-detail',
    standalone: true,
    imports: [CommonModule, RouterModule, MaterialModule, ReactiveFormsModule],
    templateUrl: './team-detail.component.html',
    styleUrls: ['./team-detail.component.scss']
})
export class TeamDetailComponent {
    private readonly formBuilder = inject(FormBuilder)
    private readonly route = inject(ActivatedRoute)
    private readonly router = inject(Router)
    private readonly teamsService = inject(TeamsService)
    private readonly notificationService = inject(NotificationService)

    public teamForm!: FormGroup
    private teamId: string | null = null
    public isEditMode = false

    ngOnInit(): void {
        this.teamId = this.route.snapshot.paramMap.get('team-id')
        this.isEditMode = !!this.teamId

        this.teamForm = this.formBuilder.group({
            name: [
                '',
                [
                    Validators.required,
                    Validators.minLength(3),
                    Validators.maxLength(100),
                    Validators.pattern(/^[A-Za-z0-9 .,'-]+$/)
                ]
            ],
            abbreviation: [
                '',
                [
                    Validators.required,
                    Validators.minLength(2),
                    Validators.maxLength(4),
                    Validators.pattern(/^[A-Za-z]+$/)
                ]
            ],
            location: ['', [Validators.minLength(2), Validators.maxLength(100)]]
        })

        if (this.isEditMode && this.teamId) {
            this.loadTeam(this.teamId)
        }
    }

    submitForm(): void {
        if (this.teamForm.invalid) {
            this.teamForm.markAllAsTouched()
            return
        }

        const teamData: Team = this.teamForm.value
        if (!this.isEditMode) {
            this.createTeam(teamData)
        } else if (this.teamId) {
            this.updateTeam(this.teamId, teamData)
        }
    }

    confirmDelete(): void {
        const confirmationMessage = `Are you sure you want to delete this team?`

        this.notificationService
            .showConfirm('Confirm Delete', confirmationMessage)
            .afterClosed()
            .subscribe((confirmed) => {
                if (confirmed && this.teamId) {
                    this.deleteTeam(this.teamId)
                }
            })
    }

    private loadTeam(id: string): void {
        this.teamsService.getTeam(id).subscribe({
            next: (team) => {
                this.teamForm.patchValue({
                    name: team.name,
                    abbreviation: team.abbreviation,
                    location: team.location
                })
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load team', err)
            }
        })
    }

    private createTeam(newTeam: Team): void {
        this.teamsService.createTeam(newTeam).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Team created successfully')
                this.router.navigate(['/admin/teams'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Create Error', 'Failed to create team', err)
            }
        })
    }

    private updateTeam(id: string, updatedTeam: Team): void {
        this.teamsService.updateTeam(id, updatedTeam).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Team updated successfully')
                this.router.navigate(['/admin/teams'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Update Error', 'Failed to update team', err)
            }
        })
    }

    private deleteTeam(id: string): void {
        this.teamsService.deleteTeam(id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Team deleted successfully')
                this.router.navigate(['/admin/teams'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Delete Error', 'Failed to delete team', err)
            }
        })
    }
}
