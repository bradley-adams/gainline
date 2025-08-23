import { Component, inject } from '@angular/core'
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { Competition } from '../../types/api'
import { CommonModule } from '@angular/common'
import { MaterialModule } from '../../shared/material/material.module'
import { NotificationService } from '../../services/notifications/notifications.service'

@Component({
    selector: 'app-competition-detail',
    standalone: true,
    imports: [CommonModule, RouterModule, MaterialModule, ReactiveFormsModule],
    templateUrl: './competition-detail.component.html',
    styleUrl: './competition-detail.component.scss'
})
export class CompetitionDetailComponent {
    private readonly formBuilder = inject(FormBuilder)
    private readonly route = inject(ActivatedRoute)
    private readonly router = inject(Router)
    private readonly competitionsService = inject(CompetitionsService)
    private readonly notificationService = inject(NotificationService)

    public competitionForm!: FormGroup
    private competitionId: string | null = null
    public isEditMode = false

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')
        this.isEditMode = !!this.competitionId

        this.competitionForm = this.formBuilder.group({
            name: ['', Validators.required]
        })

        if (this.isEditMode && this.competitionId) {
            this.loadCompetition(this.competitionId)
        }
    }

    submitForm(): void {
        if (this.competitionForm.invalid) {
            console.error('competition form is invalid')
            this.notificationService.showError('Form Error', 'Please fill out all required fields.')
            return
        }

        const competitionData: Competition = this.competitionForm.value
        if (!this.isEditMode) {
            this.createCompetition(competitionData)
        } else if (this.competitionId) {
            this.updateCompetition(this.competitionId, competitionData)
        }
    }

    confirmDelete(): void {
        const competitionName = this.competitionForm.value.name
        const confirmationMessage = `Are you sure you want to delete competition "${competitionName}"?`

        this.notificationService
            .showConfirm('Confirm Delete', confirmationMessage)
            .afterClosed()
            .subscribe((confirmed) => {
                if (confirmed && this.competitionId) {
                    this.removeCompetition(this.competitionId)
                }
            })
    }

    private loadCompetition(id: string): void {
        this.competitionsService.getCompetition(id).subscribe({
            next: (competition) => {
                this.competitionForm.patchValue({
                    name: competition.name
                })
            },
            error: (err) => {
                console.error('Error loading competition:', err)
                this.notificationService.showError('Load Error', 'Failed to load competition')
            }
        })
    }

    private createCompetition(newCompetition: Competition): void {
        this.competitionsService.createCompetition(newCompetition).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Competition created successfully', 'OK')
                this.router.navigate(['/admin/competitions'])
            },
            error: (err) => {
                console.error('Error creating competition:', err)
                this.notificationService.showError('Create Error', 'Failed to create competition')
            }
        })
    }

    private updateCompetition(id: string, updatedCompetition: Competition): void {
        this.competitionsService.updateCompetition(id, updatedCompetition).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Competition updated successfully', 'OK')
                this.router.navigate(['/admin/competitions'])
            },
            error: (err) => {
                console.error('Error updating competition:', err)
                this.notificationService.showError('Update Error', 'Failed to update competition')
            }
        })
    }

    private removeCompetition(competitionId: string): void {
        if (!competitionId) return

        this.competitionsService.deleteCompetition(competitionId).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Competition deleted successfully', 'OK')
                this.router.navigate(['/admin/competitions'])
            },
            error: (err) => {
                console.error('Error deleting competition:', err)
                this.notificationService.showError('Delete Error', 'Failed to delete competition')
            }
        })
    }
}
