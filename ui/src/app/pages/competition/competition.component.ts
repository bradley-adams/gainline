import { Component, inject } from '@angular/core';
import { CompetitionsService } from '../../services/competitions/competitions.service';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MaterialModule } from '../../shared/material/material.module';
import { Competition, CompetitionUpdate } from '../../types/api';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-competition',
  standalone: true,
  imports: [CommonModule, RouterModule, MaterialModule, ReactiveFormsModule],
  templateUrl: './competition.component.html',
  styleUrl: './competition.component.scss'
})
export class CompetitionComponent {
  private readonly formBuilder: FormBuilder = inject(FormBuilder)
  private readonly route: ActivatedRoute = inject(ActivatedRoute)
  private readonly router: Router = inject(Router)
  private readonly competitionsService = inject(CompetitionsService)
  
  public competitionForm: FormGroup;

  private competitionId: string | null;
  public isEditMode: boolean;

  constructor() {
    this.competitionId = this.route.snapshot.paramMap.get('competition-id')
    this.isEditMode = !!this.competitionId

    this.competitionForm = this.formBuilder.group({
      name: ['', Validators.required],
    });

    if (this.isEditMode && this.competitionId) {
      this.loadCompetition(this.competitionId)
    }
  }

  submitForm(): void {
      if (this.competitionForm.invalid) {
          console.error(`competition form is invalid`)
          return
      }

      const competitionData: CompetitionUpdate = this.competitionForm.value
      if (!this.isEditMode) {
          this.createCompetition(competitionData)
      } else if (this.competitionId) {
          this.updateCompetition(this.competitionId, competitionData)
      }
  }

  loadCompetition(id: string): void {
    this.competitionsService.getCompetition(id).subscribe({
      next: competition => {
        this.competitionForm.patchValue({
          name: competition.name
        })
      },
      error: err => console.error('Error loading competition:', err)
    })
  }

  createCompetition(newCompetition: CompetitionUpdate): void {
    this.competitionsService.createCompetition(newCompetition).subscribe({
      next: competition => {
        this.router.navigate(['/competitions', competition.id])
      },
      error: err => console.error('Error creating competition:', err)
    })
  }

  updateCompetition(id: string, updatedCompetition: CompetitionUpdate): void {
    this.competitionsService.updateCompetition(id, updatedCompetition).subscribe({
      next: competition => {
        this.router.navigate(['/competitions', this.competitionId])
      },
      error: err => console.error('Error updating competition:', err)
    })
  }

}
