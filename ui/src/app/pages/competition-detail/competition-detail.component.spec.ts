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
import { NotificationService } from '../../services/notifications/notifications.service'

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
            'showError',
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

        it('should not submit if form is invalid', () => {
            component.submitForm()
            expect(notificationService.showError).toHaveBeenCalledWith(
                'Form Error',
                'Please fill out all required fields.'
            )
        })

        it('should call createCompetition when in create mode and form is valid', () => {
            component.competitionForm.setValue({ name: 'New Comp' })
            component.submitForm()
            expect(competitionsService.createCompetition).toHaveBeenCalledWith({ name: 'New Comp' })
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Competition created successfully')
        })

        it('should show error if createCompetition fails', () => {
            competitionsService.createCompetition.and.returnValue(throwError(() => new Error('Failed')))
            component.competitionForm.setValue({ name: 'Bad Comp' })
            component.submitForm()
            expect(notificationService.showError).toHaveBeenCalledWith(
                'Create Error',
                'Failed to create competition'
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

        it('should load competition when in edit mode', () => {
            expect(competitionsService.getCompetition).toHaveBeenCalledWith('123')
            expect(component.competitionForm.value.name).toBe('Competition 1')
        })

        it('should show error if loadCompetition fails', () => {
            competitionsService.getCompetition.and.returnValue(throwError(() => new Error('Failed')))
            component['loadCompetition']('123')
            expect(notificationService.showError).toHaveBeenCalledWith(
                'Load Error',
                'Failed to load competition'
            )
        })

        it('should call updateCompetition when in edit mode and form is valid', () => {
            component.competitionForm.setValue({ name: 'Updated Comp' })
            component.submitForm()
            expect(competitionsService.updateCompetition).toHaveBeenCalledWith('123', {
                name: 'Updated Comp'
            })
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Competition updated successfully')
        })

        it('should show error if updateCompetition fails', () => {
            competitionsService.updateCompetition.and.returnValue(throwError(() => new Error('Failed')))
            component.competitionForm.setValue({ name: 'Bad Update' })
            component.submitForm()
            expect(notificationService.showError).toHaveBeenCalledWith(
                'Update Error',
                'Failed to update competition'
            )
        })

        it('should call deleteCompetition when confirmed', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(competitionsService.deleteCompetition).toHaveBeenCalledWith('123')
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Competition deleted successfully')
        })

        it('should not call deleteCompetition when cancelled', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(false) } as any)
            component.confirmDelete()
            expect(competitionsService.deleteCompetition).not.toHaveBeenCalled()
            expect(notificationService.showSnackbar).not.toHaveBeenCalled()
        })

        it('should show error if deleteCompetition fails', () => {
            competitionsService.deleteCompetition.and.returnValue(throwError(() => new Error('Failed')))
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(notificationService.showError).toHaveBeenCalledWith(
                'Delete Error',
                'Failed to delete competition'
            )
        })
    })
})
