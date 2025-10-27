import { ComponentFixture, TestBed } from '@angular/core/testing'
import { ActivatedRoute, provideRouter, Router } from '@angular/router'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { of, throwError } from 'rxjs'

import { TeamDetailComponent } from './team-detail.component'
import { TeamListComponent } from '../team-list/team-list.component'
import { TeamsService } from '../../services/teams/teams.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { Team } from '../../types/api'

describe('TeamDetailComponent', () => {
    let component: TeamDetailComponent
    let fixture: ComponentFixture<TeamDetailComponent>
    let router: Router
    let teamsService: jasmine.SpyObj<TeamsService>
    let notificationService: jasmine.SpyObj<NotificationService>

    const mockTeam1: Team = {
        id: 'team1',
        abbreviation: 'T1',
        location: 'City A',
        name: 'Team One',
        created_at: new Date('2024-01-01T00:00:00Z'),
        updated_at: new Date('2024-01-01T00:00:00Z')
    }

    const mockTeam2: Team = {
        id: 'team2',
        abbreviation: 'T2',
        location: 'City B',
        name: 'Team Two',
        created_at: new Date('2024-01-02T00:00:00Z'),
        updated_at: new Date('2024-01-02T00:00:00Z')
    }

    function mockRoute(id: string | null) {
        return {
            snapshot: {
                paramMap: {
                    get: (key: string) => (key === 'team-id' ? id : null)
                }
            }
        }
    }

    beforeEach(async () => {
        teamsService = jasmine.createSpyObj('TeamsService', [
            'getTeam',
            'createTeam',
            'updateTeam',
            'deleteTeam'
        ])
        teamsService.getTeam.and.returnValue(of(mockTeam1))
        teamsService.createTeam.and.returnValue(of(mockTeam1))
        teamsService.updateTeam.and.returnValue(of(mockTeam2))
        teamsService.deleteTeam.and.returnValue(of(undefined))

        notificationService = jasmine.createSpyObj('NotificationService', [
            'showConfirm',
            'showErrorAndLog',
            'showWarnAndLog',
            'showSnackbar'
        ])

        await TestBed.configureTestingModule({
            imports: [TeamDetailComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/teams',
                        component: TeamListComponent
                    }
                ]),
                { provide: TeamsService, useValue: teamsService },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()
    })

    describe('Create mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, {
                useValue: mockRoute(null)
            })
            fixture = TestBed.createComponent(TeamDetailComponent)
            component = fixture.componentInstance
            router = TestBed.inject(Router)
            fixture.detectChanges()
        })

        it('should create', () => {
            expect(component).toBeTruthy()
        })

        it('should mark name, abbreviation, and location as required', () => {
            const nameControl = component.teamForm.get('name')
            const abbreviationControl = component.teamForm.get('abbreviation')
            const locationControl = component.teamForm.get('location')

            // Name
            nameControl?.setValue('')
            expect(nameControl?.valid).toBeFalse()
            nameControl?.setValue('Team One')
            expect(nameControl?.valid).toBeTrue()

            // Abbreviation
            abbreviationControl?.setValue('')
            expect(abbreviationControl?.valid).toBeFalse()
            abbreviationControl?.setValue('T1')
            expect(abbreviationControl?.valid).toBeTrue()

            // Location
            locationControl?.setValue('')
            expect(locationControl?.valid).toBeFalse()
            locationControl?.setValue('City A')
            expect(locationControl?.valid).toBeTrue()
        })

        it('should not submit form if invalid', () => {
            component.submitForm()
            expect(notificationService.showWarnAndLog).toHaveBeenCalledWith(
                'Form Error',
                'Team form is invalid'
            )
        })

        it('should call createTeam when form is valid in create mode', () => {
            component.isEditMode = false
            component.teamForm.setValue({ name: 'Team One', abbreviation: 'T1', location: 'City A' })
            component.submitForm()

            expect(teamsService.createTeam).toHaveBeenCalledWith({
                name: 'Team One',
                abbreviation: 'T1',
                location: 'City A'
            })
        })

        it('should show error notification if createTeam fails', () => {
            const mockError = new Error('Failed to create')
            teamsService.createTeam.and.returnValue(throwError(() => mockError))

            component.isEditMode = false
            component.teamForm.patchValue({ name: 'Team One', abbreviation: 'T1', location: 'City A' })
            component.submitForm()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Create Error',
                'Failed to create team',
                mockError
            )
        })

        it('should navigate to team list after successful createTeam', () => {
            const routerSpy = spyOn(router, 'navigateByUrl')
            component.isEditMode = false
            component.teamForm.setValue({ name: 'Team One', abbreviation: 'T1', location: 'City A' })
            component.submitForm()

            const call = routerSpy.calls.all()[0].args[0].toString()
            expect(call).toEqual('/admin/teams')
        })
    })

    describe('Edit mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute('team1') })
            fixture = TestBed.createComponent(TeamDetailComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should load team data when in edit mode', () => {
            expect(teamsService.getTeam).toHaveBeenCalledWith('team1')
            expect(component.teamForm.value).toEqual({
                name: mockTeam1.name,
                abbreviation: mockTeam1.abbreviation,
                location: mockTeam1.location
            })
        })

        it('should show error notification if loading team fails', () => {
            const mockError = new Error('Failed to load')
            teamsService.getTeam.and.returnValue(throwError(() => mockError))

            component['loadTeam']('123')
            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Load Error',
                'Failed to load team',
                mockError
            )
        })

        it('should call updateTeam with form data when valid', () => {
            component.teamForm.patchValue({ name: 'test' })
            component.submitForm()

            expect(teamsService.updateTeam).toHaveBeenCalledWith('team1', {
                name: 'test',
                abbreviation: mockTeam1.abbreviation,
                location: mockTeam1.location
            })
        })

        it('should show error notification if updateTeam fails', () => {
            const mockError = new Error('Failed to update')
            teamsService.updateTeam.and.returnValue(throwError(() => mockError))

            component.teamForm.patchValue({ name: 1 })
            component.submitForm()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Update Error',
                'Failed to update team',
                mockError
            )
        })

        it('should call deleteTeam and show snackbar when deletion is confirmed', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(teamsService.deleteTeam).toHaveBeenCalledWith('team1')
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Team deleted successfully')
        })

        it('should not call deleteTeam or show snackbar when deletion is cancelled', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(false) } as any)
            component.confirmDelete()
            expect(teamsService.deleteTeam).not.toHaveBeenCalled()
            expect(notificationService.showSnackbar).not.toHaveBeenCalled()
        })

        it('should show error notification if deleteTeam fails', () => {
            const mockError = new Error('Failed')
            teamsService.deleteTeam.and.returnValue(throwError(() => mockError))
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)

            component.confirmDelete()
            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Delete Error',
                'Failed to delete team',
                mockError
            )
        })
    })
})
