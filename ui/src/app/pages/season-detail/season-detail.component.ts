import { COMMA, ENTER } from '@angular/cdk/keycodes'
import { CommonModule } from '@angular/common'
import { Component, ElementRef, inject, OnInit, ViewChild } from '@angular/core'
import {
    AbstractControl,
    FormArray,
    FormBuilder,
    FormControl,
    FormGroup,
    ReactiveFormsModule,
    ValidationErrors,
    Validators
} from '@angular/forms'
import { MatAutocompleteModule } from '@angular/material/autocomplete'
import { MatChipInputEvent, MatChipsModule } from '@angular/material/chips'
import { MatNativeDateModule } from '@angular/material/core'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatFormFieldModule } from '@angular/material/form-field'
import { MatIconModule } from '@angular/material/icon'
import { MatInputModule } from '@angular/material/input'
import { MatTimepickerModule } from '@angular/material/timepicker'
import { ActivatedRoute, Router, RouterModule } from '@angular/router'
import { combineLatest, map } from 'rxjs'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { TeamsService } from '../../services/teams/teams.service'
import { MaterialModule } from '../../shared/material/material.module'
import { Season, StageType, Team } from '../../types/api'

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
        MatChipsModule,
        MatIconModule,
        MatAutocompleteModule,
        BreadcrumbComponent
    ],
    templateUrl: './season-detail.component.html',
    styleUrls: ['./season-detail.component.scss']
})
export class SeasonDetailComponent implements OnInit {
    private readonly formBuilder = inject(FormBuilder)
    private readonly route = inject(ActivatedRoute)
    private readonly router = inject(Router)
    private readonly seasonsService = inject(SeasonsService)
    private readonly teamsService = inject(TeamsService)
    private readonly notificationService = inject(NotificationService)

    public StageType = StageType

    public stageTypeOptions = [
        { value: StageType.StageTypeRegular, label: 'Regular' },
        { value: StageType.StageTypeFinals, label: 'Finals' }
    ]

    public seasonForm!: FormGroup
    public startDateControl = new FormControl<Date | null>(null, Validators.required)
    public startTimeControl = new FormControl<Date | null>(null, Validators.required)
    public endDateControl = new FormControl<Date | null>(null, Validators.required)
    public endTimeControl = new FormControl<Date | null>(null, Validators.required)
    public isEditMode = false
    public competitionId: string | null = null
    public seasonId: string | null = null
    public teams: Team[] = []
    public filteredTeams: Team[] = []
    public separatorKeysCodes: number[] = [ENTER, COMMA]
    public selectedTeamIds: string[] = []

    @ViewChild('teamInput') teamInput!: ElementRef<HTMLInputElement>

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')
        this.seasonId = this.route.snapshot.paramMap.get('season-id')
        this.isEditMode = !!this.seasonId

        this.initForm()
        this.initDateTimeSync()
        this.loadTeams()

        if (!this.isEditMode) {
            this.addStage()
        }

        if (this.isEditMode && this.competitionId && this.seasonId) {
            this.loadSeason(this.competitionId, this.seasonId)
        }
    }

    private initForm(): void {
        this.seasonForm = this.formBuilder.group(
            {
                start_datetime: [null, Validators.required],
                end_datetime: [null, Validators.required],
                teams: [[], [Validators.required, this.minMaxArrayValidator(2, 20)]],

                stages: this.formBuilder.array([], Validators.required)
            },
            {
                validators: this.endAfterStartValidator
            }
        )
    }

    private initDateTimeSync(): void {
        combineLatest([this.startDateControl.valueChanges, this.startTimeControl.valueChanges])
            .pipe(map(([date, time]) => this.combineDateTime(date, time)))
            .subscribe((start) => this.seasonForm.get('start_datetime')?.setValue(start))

        combineLatest([this.endDateControl.valueChanges, this.endTimeControl.valueChanges])
            .pipe(map(([date, time]) => this.combineDateTime(date, time)))
            .subscribe((end) => this.seasonForm.get('end_datetime')?.setValue(end))
    }

    private combineDateTime(date: Date | null, time: Date | null): Date | null {
        if (!date || !time) return null
        const combined = new Date(date)
        combined.setHours(time.getHours(), time.getMinutes(), 0)
        return combined
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

        const { start_datetime, end_datetime, teams, stages } = this.seasonForm.value

        const seasonData: Partial<Season> = {
            start_date: start_datetime,
            end_date: end_datetime,
            teams,
            stages: stages.map((stage: any, index: number) => ({
                ...stage,
                orderIndex: index + 1
            }))
        }

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
                this.filteredTeams = [...teams]
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load teams', err)
            }
        })
    }

    private loadSeason(competitionId: string, id: string): void {
        this.seasonsService.getSeason(competitionId, id).subscribe({
            next: (season) => {
                const startDate = new Date(season.start_date)
                const endDate = new Date(season.end_date)

                this.startDateControl.setValue(startDate)
                this.startTimeControl.setValue(startDate)
                this.endDateControl.setValue(endDate)
                this.endTimeControl.setValue(endDate)

                this.seasonForm.patchValue({
                    start_datetime: startDate,
                    end_datetime: endDate,
                    teams: (season.teams as Team[]).map((t) => t.id)
                })

                this.stages.clear()
                season.stages?.forEach((stage) => {
                    this.stages.push(
                        this.formBuilder.group({
                            name: [stage.name, Validators.required],
                            stageType: [stage.stage_type, Validators.required]
                        })
                    )
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

    private endAfterStartValidator = (group: AbstractControl): ValidationErrors | null => {
        const start = group.get('start_datetime')?.value
        const end = group.get('end_datetime')?.value
        if (!start || !end) return null
        return new Date(end) <= new Date(start) ? { endBeforeStart: true } : null
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

    getTeamName(id: string): string {
        const team = this.teams.find((t) => t.id === id)
        return team ? team.name : id
    }

    addTeamFromInput(event: MatChipInputEvent) {
        const value = event.value?.trim()
        if (value) {
            const match = this.teams.find((t) => t.name.toLowerCase() === value.toLowerCase())
            if (match) this.addTeam(match.id)
        }

        this.teamInput.nativeElement.value = ''
    }

    selectTeam(teamId: string) {
        this.addTeam(teamId)
        this.teamInput.nativeElement.value = ''
    }

    private addTeam(teamId: string) {
        if (!this.selectedTeamIds.includes(teamId)) {
            this.selectedTeamIds.push(teamId)
            ;(this.seasonForm.get('teams') as FormControl).setValue([...this.selectedTeamIds])
        }
    }

    removeTeam(teamId: string) {
        this.selectedTeamIds = this.selectedTeamIds.filter((id) => id !== teamId)
        ;(this.seasonForm.get('teams') as FormControl).setValue([...this.selectedTeamIds])
    }

    get stages(): FormArray {
        return this.seasonForm.get('stages') as FormArray
    }

    addStage(): void {
        this.stages.push(
            this.formBuilder.group({
                name: ['', Validators.required],
                stageType: [StageType.StageTypeRegular, Validators.required]
            })
        )
    }

    removeStage(index: number): void {
        this.stages.removeAt(index)
    }

    moveStageUp(i: number) {
        if (i === 0) return
        const stage = this.stages.at(i)
        this.stages.removeAt(i)
        this.stages.insert(i - 1, stage)
    }

    moveStageDown(i: number) {
        if (i === this.stages.length - 1) return
        const stage = this.stages.at(i)
        this.stages.removeAt(i)
        this.stages.insert(i + 1, stage)
    }
}
