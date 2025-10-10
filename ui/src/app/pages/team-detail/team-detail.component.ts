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
            name: ['', Validators.required],
            abbreviation: ['', Validators.required],
            location: ['', Validators.required]
        })

        if (this.isEditMode && this.teamId) {
            this.loadTeam(this.teamId)
        }
    }

    submitForm(): void {
        if (this.teamForm.invalid) {
            console.error('team form is invalid')
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
                console.error('Error loading team:', err)
                this.notificationService.showError('Load Error', 'Failed to load team')
            }
        })
    }

    private createTeam(newTeam: Team): void {
        this.teamsService.createTeam(newTeam).subscribe({
            next: () => {
                this.router.navigate(['/admin/teams'])
            },
            error: (err) => {
                console.error('Error creating team:', err)
                this.notificationService.showError('Create Error', 'Failed to create team')
            }
        })
    }

    private updateTeam(id: string, updatedTeam: Team): void {
        this.teamsService.updateTeam(id, updatedTeam).subscribe({
            next: () => {
                this.router.navigate(['/admin/teams'])
            },
            error: (err) => {
                console.error('Error updating team:', err)
                this.notificationService.showError('Update Error', 'Failed to update team')
            }
        })
    }

    private deleteTeam(id: string): void {
        this.teamsService.deleteTeam(id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Team deleted successfully', 'OK')
                this.router.navigate(['/admin/teams'])
            },
            error: (err) => {
                console.error('Error deleting team:', err)
                this.notificationService.showError('Delete Error', 'Failed to delete team')
            }
        })
    }
}
