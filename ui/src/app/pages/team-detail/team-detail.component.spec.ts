import { ComponentFixture, TestBed } from '@angular/core/testing'

import { TeamDetailComponent } from './team-detail.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ActivatedRoute, provideRouter, Router } from '@angular/router'
import { TeamListComponent } from '../team-list/team-list.component'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { TeamsService } from '../../services/teams/teams.service'
import { Team } from '../../types/api'
import { of, throwError } from 'rxjs'

describe('TeamDetailComponent', () => {
    let component: TeamDetailComponent
    let fixture: ComponentFixture<TeamDetailComponent>
    let router: Router

    let teamsService: jasmine.SpyObj<TeamsService>

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
        teamsService = jasmine.createSpyObj('TeamsService', ['getTeam', 'createTeam', 'updateTeam'])
        teamsService.getTeam.and.returnValue(of(mockTeam1))
        teamsService.createTeam.and.returnValue(of(mockTeam1))
        teamsService.updateTeam.and.returnValue(of(mockTeam2))

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
                { provide: TeamsService, useValue: teamsService }
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

        it('should not submit if form is invalid', () => {
            spyOn(console, 'error')
            component.submitForm()
            expect(console.error).toHaveBeenCalledWith('team form is invalid')
        })

        it('should call createTeam when in create mode and form is valid', () => {
            component.isEditMode = false
            component.teamForm.setValue({
                name: 'Team One',
                abbreviation: 'T1',
                location: 'City A'
            })

            component.submitForm()

            expect(teamsService.createTeam).toHaveBeenCalledWith({
                name: 'Team One',
                abbreviation: 'T1',
                location: 'City A'
            })
        })

        it('should log error if createTeam fails', () => {
            const error = new Error('Failed to create')
            spyOn(console, 'error')
            teamsService.createTeam.and.returnValue(throwError(() => error))

            component.isEditMode = false
            component.teamForm.patchValue({
                name: 'Team One',
                abbreviation: 'T1',
                location: 'City A'
            })
            component.submitForm()

            expect(console.error).toHaveBeenCalledWith('Error creating team:', error)
        })

        it('should navigate after createTeam', () => {
            const routerSpy = spyOn(router, 'navigateByUrl')

            component.isEditMode = false
            component.teamForm.setValue({
                name: 'Team One',
                abbreviation: 'T1',
                location: 'City A'
            })
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

        it('should load team when in edit mode', () => {
            expect(teamsService.getTeam).toHaveBeenCalledWith('team1')

            expect(component.teamForm.value).toEqual({
                name: mockTeam1.name,
                abbreviation: mockTeam1.abbreviation,
                location: mockTeam1.location
            })
        })

        it('should log error if loadTeam fails', () => {
            const error = new Error('Failed to load')
            spyOn(console, 'error')
            teamsService.getTeam.and.returnValue(throwError(() => error))

            component['loadTeam']('123')

            expect(console.error).toHaveBeenCalledWith('Error loading team:', error)
        })

        it('should call updateTeam when in edit mode and form is valid', () => {
            component.teamForm.patchValue({ name: 'test' })

            component.submitForm()

            expect(teamsService.updateTeam).toHaveBeenCalledWith('team1', {
                name: 'test',
                abbreviation: mockTeam1.abbreviation,
                location: mockTeam1.location
            })
        })

        it('should log error if updateTeam fails', () => {
            const error = new Error('Failed to update')
            spyOn(console, 'error')
            teamsService.updateTeam.and.returnValue(throwError(() => error))

            component.teamForm.patchValue({ name: 1 })
            component.submitForm()

            expect(console.error).toHaveBeenCalledWith('Error updating team:', error)
        })
    })
})
