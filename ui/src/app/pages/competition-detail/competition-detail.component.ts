import { Component, inject } from '@angular/core'
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { Competition } from '../../types/api'
import { CommonModule } from '@angular/common'
import { MaterialModule } from '../../shared/material/material.module'
import { NotificationService } from '../../services/notifications/notifications.service'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'

@Component({
    selector: 'app-competition-detail',
    standalone: true,
    imports: [CommonModule, RouterModule, MaterialModule, ReactiveFormsModule, BreadcrumbComponent],
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
            this.notificationService.showErrorAndLog(
                'Form Error',
                'Required fields cannot be left blank',
                new Error('competition form is invalid')
            )
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
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load competition', err)
            }
        })
    }

    private createCompetition(newCompetition: Competition): void {
        this.competitionsService.createCompetition(newCompetition).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Competition created successfully')
                this.router.navigate(['/admin/competitions'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Create Error', 'Failed to create competition', err)
            }
        })
    }

    private updateCompetition(id: string, updatedCompetition: Competition): void {
        this.competitionsService.updateCompetition(id, updatedCompetition).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Competition updated successfully')
                this.router.navigate(['/admin/competitions'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Update Error', 'Failed to update competition', err)
            }
        })
    }

    private removeCompetition(competitionId: string): void {
        if (!competitionId) return

        this.competitionsService.deleteCompetition(competitionId).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Competition deleted successfully')
                this.router.navigate(['/admin/competitions'])
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Delete Error', 'Failed to delete competition', err)
            }
        })
    }
}
