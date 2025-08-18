import { Component, inject } from '@angular/core'
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { Competition } from '../../types/api'
import { CommonModule } from '@angular/common'
import { MaterialModule } from '../../shared/material/material.module'

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
            return
        }

        const competitionData: Competition = this.competitionForm.value
        if (!this.isEditMode) {
            this.createCompetition(competitionData)
        } else if (this.competitionId) {
            this.updateCompetition(this.competitionId, competitionData)
        }
    }

    private loadCompetition(id: string): void {
        this.competitionsService.getCompetition(id).subscribe({
            next: (competition) => {
                this.competitionForm.patchValue({
                    name: competition.name
                })
            },
            error: (err) => console.error('Error loading competition:', err)
        })
    }

    private createCompetition(newCompetition: Competition): void {
        this.competitionsService.createCompetition(newCompetition).subscribe({
            next: () => {
                this.router.navigate(['/admin/competitions'])
            },
            error: (err) => console.error('Error creating competition:', err)
        })
    }

    private updateCompetition(id: string, updatedCompetition: Competition): void {
        this.competitionsService.updateCompetition(id, updatedCompetition).subscribe({
            next: () => {
                this.router.navigate(['/admin/competitions'])
            },
            error: (err) => console.error('Error updating competition:', err)
        })
    }
}
