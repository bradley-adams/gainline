import { ComponentFixture, TestBed } from '@angular/core/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { ActivatedRoute, provideRouter } from '@angular/router'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { of, throwError } from 'rxjs'

import { CompetitionDetailComponent } from './competition-detail.component'
import { CompetitionListComponent } from '../competition-list/competition-list.component'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { Competition } from '../../types/api'

describe('CompetitionDetailComponent', () => {
    let component: CompetitionDetailComponent
    let fixture: ComponentFixture<CompetitionDetailComponent>

    let competitionsService: jasmine.SpyObj<CompetitionsService>
    let notificationService: jasmine.SpyObj<NotificationService>

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
            'updateCompetition',
            'deleteCompetition'
        ])
        competitionsService.getCompetition.and.returnValue(of(mockCompetition1))
        competitionsService.createCompetition.and.returnValue(of(mockCompetition1))
        competitionsService.updateCompetition.and.returnValue(of(mockCompetition2))
        competitionsService.deleteCompetition.and.returnValue(of(undefined))

        notificationService = jasmine.createSpyObj('NotificationService', [
            'showConfirm',
            'showErrorAndLog',
            'showWarnAndLog',
            'showSnackbar'
        ])

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
                { provide: CompetitionsService, useValue: competitionsService },
                { provide: NotificationService, useValue: notificationService }
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

        it('should mark name as invalid if too short', () => {
            const control = component.competitionForm.get('name')
            control?.setValue('Te') // 2 characters
            expect(control?.hasError('minlength')).toBeTrue()
            expect(control?.valid).toBeFalse()
        })

        it('should mark name as invalid if too long', () => {
            const control = component.competitionForm.get('name')
            control?.setValue('A'.repeat(101)) // 101 characters
            expect(control?.hasError('maxlength')).toBeTrue()
            expect(control?.valid).toBeFalse()
        })

        it('should mark name as invalid if it contains invalid characters', () => {
            const control = component.competitionForm.get('name')
            control?.setValue('@@!!') // invalid chars
            expect(control?.hasError('pattern')).toBeTrue()
            expect(control?.valid).toBeFalse()
        })

        it('should not submit if name is too short', () => {
            component.competitionForm.setValue({ name: 'Te' })
            component.submitForm()
            expect(competitionsService.createCompetition).not.toHaveBeenCalled()
            expect(notificationService.showWarnAndLog).not.toHaveBeenCalled() // because frontend validation shows errors
        })

        it('should create a competition when form is valid', () => {
            component.competitionForm.setValue({ name: 'New Comp' })
            component.submitForm()
            expect(competitionsService.createCompetition).toHaveBeenCalledWith({ name: 'New Comp' })
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Competition created successfully')
        })

        it('should show an error when createCompetition fails', () => {
            const mockError = new Error('Failed')
            competitionsService.createCompetition.and.returnValue(throwError(() => mockError))

            component.competitionForm.setValue({ name: 'Bad Comp' })
            component.submitForm()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Create Error',
                'Failed to create competition',
                mockError
            )
        })
    })

    describe('Edit mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute('123') })
            fixture = TestBed.createComponent(CompetitionDetailComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should load the existing competition in edit mode', () => {
            expect(competitionsService.getCompetition).toHaveBeenCalledWith('123')
            expect(component.competitionForm.value.name).toBe('Competition 1')
        })

        it('should show an error when loading the competition fails', () => {
            const mockError = new Error('Failed')
            competitionsService.getCompetition.and.returnValue(throwError(() => mockError))

            component['loadCompetition']('123')
            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Load Error',
                'Failed to load competition',
                mockError
            )
        })

        it('should update the competition when form is valid', () => {
            component.competitionForm.setValue({ name: 'Updated Comp' })
            component.submitForm()
            expect(competitionsService.updateCompetition).toHaveBeenCalledWith('123', {
                name: 'Updated Comp'
            })
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Competition updated successfully')
        })

        it('should show an error when updating the competition fails', () => {
            const mockError = new Error('Failed')
            competitionsService.updateCompetition.and.returnValue(throwError(() => mockError))

            component.competitionForm.setValue({ name: 'Bad Update' })
            component.submitForm()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Update Error',
                'Failed to update competition',
                mockError
            )
        })

        it('should delete the competition when confirmed', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(competitionsService.deleteCompetition).toHaveBeenCalledWith('123')
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Competition deleted successfully')
        })

        it('should not delete the competition when cancelled', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(false) } as any)
            component.confirmDelete()
            expect(competitionsService.deleteCompetition).not.toHaveBeenCalled()
            expect(notificationService.showSnackbar).not.toHaveBeenCalled()
        })

        it('should show an error when deleting the competition fails', () => {
            const mockError = new Error('Failed')
            competitionsService.deleteCompetition.and.returnValue(throwError(() => mockError))

            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Delete Error',
                'Failed to delete competition',
                mockError
            )
        })
    })
})
