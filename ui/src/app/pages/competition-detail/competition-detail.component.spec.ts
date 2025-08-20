import { ComponentFixture, TestBed } from '@angular/core/testing'
import { ActivatedRoute, provideRouter } from '@angular/router'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { provideHttpClient } from '@angular/common/http'
import { CompetitionDetailComponent } from './competition-detail.component'
import { CompetitionListComponent } from '../competition-list/competition-list.component'
import { Competition } from '../../types/api'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { of, throwError } from 'rxjs'

describe('CompetitionDetailComponent', () => {
    let component: CompetitionDetailComponent
    let fixture: ComponentFixture<CompetitionDetailComponent>

    let competitionsService: jasmine.SpyObj<CompetitionsService>

    const mockCompetition1: Competition = {
        id: 'comp1',
        name: 'Competition 1',
        created_at: new Date('2023-01-01T00:00:00Z'),
        updated_at: new Date('2023-01-02T00:00:00Z')
    }

    const mockCompetition2: Competition = {
        id: 'comp2',
        name: 'Competition 2',
        created_at: new Date('2023-02-01T00:00:00Z'),
        updated_at: new Date('2023-02-02T00:00:00Z')
    }

    function mockRoute(id: string | null) {
        return {
            snapshot: {
                paramMap: {
                    get: (key: string) => (key === 'competition-id' ? id : null)
                }
            }
        }
    }

    beforeEach(async () => {
        competitionsService = jasmine.createSpyObj('CompetitionsService', [
            'getCompetition',
            'createCompetition',
            'updateCompetition'
        ])
        competitionsService.getCompetition.and.returnValue(of(mockCompetition1))
        competitionsService.createCompetition.and.returnValue(of(mockCompetition1))
        competitionsService.updateCompetition.and.returnValue(of(mockCompetition2))

        await TestBed.configureTestingModule({
            imports: [CompetitionDetailComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions',
                        component: CompetitionListComponent
                    }
                ]),
                { provide: CompetitionsService, useValue: competitionsService }
            ]
        }).compileComponents()
    })

    describe('Create mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute(null) })
            fixture = TestBed.createComponent(CompetitionDetailComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should create', () => {
            expect(component).toBeTruthy()
        })

        it('should mark name as required', () => {
            const control = component.competitionForm.get('name')
            control?.setValue('')
            expect(control?.valid).toBeFalse()

            control?.setValue('World Cup')
            expect(control?.valid).toBeTrue()
        })

        it('should not submit if form is invalid', () => {
            spyOn(console, 'error')
            component.submitForm()
            expect(console.error).toHaveBeenCalledWith('competition form is invalid')
        })

        it('should call createCompetition when in create mode and form is valid', () => {
            component.isEditMode = false
            component.competitionForm.setValue({ name: 'New Comp' })

            component.submitForm()

            expect(competitionsService.createCompetition).toHaveBeenCalledWith({ name: 'New Comp' })
        })

        it('should log error if createCompetition fails', () => {
            const error = new Error('Failed to create')
            spyOn(console, 'error')
            competitionsService.createCompetition.and.returnValue(throwError(() => error))

            component.isEditMode = false
            component.competitionForm.setValue({ name: 'Bad Comp' })
            component.submitForm()

            expect(console.error).toHaveBeenCalledWith('Error creating competition:', error)
        })
    })

    describe('Edit mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute('123') })
            fixture = TestBed.createComponent(CompetitionDetailComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should load competition when in edit mode', () => {
            expect(competitionsService.getCompetition).toHaveBeenCalledWith('123')
            expect(component.competitionForm.value.name).toBe('Competition 1')
        })

        it('should log error if loadCompetition fails', () => {
            const error = new Error('Failed to load')
            spyOn(console, 'error')
            competitionsService.getCompetition.and.returnValue(throwError(() => error))

            component['loadCompetition']('123') // trigger loadComp again

            expect(console.error).toHaveBeenCalledWith('Error loading competition:', error)
        })

        it('should call updateCompetition when in edit mode and form is valid', () => {
            component.competitionForm.setValue({ name: 'Updated Comp' })

            component.submitForm()

            expect(competitionsService.updateCompetition).toHaveBeenCalledWith('123', {
                name: 'Updated Comp'
            })
        })

        it('should log error if updateCompetition fails', () => {
            const error = new Error('Failed to update')
            spyOn(console, 'error')
            competitionsService.updateCompetition.and.returnValue(throwError(() => error))

            component.competitionForm.setValue({ name: 'Bad Update' })
            component.submitForm()

            expect(console.error).toHaveBeenCalledWith('Error updating competition:', error)
        })
    })
})
