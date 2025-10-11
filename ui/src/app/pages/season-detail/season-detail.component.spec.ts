import { ComponentFixture, TestBed } from '@angular/core/testing'

import { SeasonDetailComponent } from './season-detail.component'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ActivatedRoute, provideRouter, Router } from '@angular/router'
import { SeasonListComponent } from '../season-list/season-list.component'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { TeamsService } from '../../services/teams/teams.service'
import { Season, Team } from '../../types/api'
import { of, throwError } from 'rxjs'
import { NotificationService } from '../../services/notifications/notifications.service'

describe('SeasonDetailComponent', () => {
    let component: SeasonDetailComponent
    let fixture: ComponentFixture<SeasonDetailComponent>
    let router: Router

    let seasonsService: jasmine.SpyObj<SeasonsService>
    let teamsService: jasmine.SpyObj<TeamsService>
    let notificationService: jasmine.SpyObj<NotificationService>

    const mockTeams: Team[] = [
        {
            id: 'team1',
            abbreviation: 'T1',
            location: 'City A',
            name: 'Team One',
            created_at: new Date('2024-01-01T00:00:00Z'),
            updated_at: new Date('2024-01-01T00:00:00Z')
        },
        {
            id: 'team2',
            abbreviation: 'T2',
            location: 'City B',
            name: 'Team Two',
            created_at: new Date('2024-01-02T00:00:00Z'),
            updated_at: new Date('2024-01-02T00:00:00Z')
        },
        {
            id: 'team3',
            abbreviation: 'T3',
            location: 'City C',
            name: 'Team Three',
            created_at: new Date('2024-01-03T00:00:00Z'),
            updated_at: new Date('2024-01-03T00:00:00Z')
        },
        {
            id: 'team4',
            abbreviation: 'T4',
            location: 'City D',
            name: 'Team Four',
            created_at: new Date('2024-01-04T00:00:00Z'),
            updated_at: new Date('2024-01-04T00:00:00Z')
        }
    ]

    const mockSeason1: Season = {
        id: 'season1',
        competition_id: 'comp1',
        rounds: 3,
        start_date: new Date('2025-01-01T00:00:00Z'),
        end_date: new Date('2026-01-01T12:59:00Z'),
        teams: mockTeams,
        created_at: new Date('2024-12-01T00:00:00Z'),
        updated_at: new Date('2024-12-01T00:00:00Z')
    }

    const mockSeason2: Season = {
        id: 'season2',
        competition_id: 'comp1',
        rounds: 0,
        start_date: new Date('2024-01-01T00:00:00Z'),
        end_date: new Date('2024-12-31T23:59:59Z'),
        teams: mockTeams,
        created_at: new Date('2023-12-01T00:00:00Z'),
        updated_at: new Date('2023-12-01T00:00:00Z')
    }

    function mockRoute(competitionId: string | null, seasonId: string | null) {
        return {
            snapshot: {
                paramMap: {
                    get: (key: string) => {
                        if (key === 'competition-id') return competitionId
                        if (key === 'season-id') return seasonId
                        return null
                    }
                }
            }
        }
    }

    beforeEach(async () => {
        seasonsService = jasmine.createSpyObj('SeasonsService', [
            'getSeason',
            'createSeason',
            'updateSeason',
            'deleteSeason'
        ])
        seasonsService.getSeason.and.returnValue(of(mockSeason1))
        seasonsService.createSeason.and.returnValue(of(mockSeason1))
        seasonsService.updateSeason.and.returnValue(of(mockSeason2))
        seasonsService.deleteSeason.and.returnValue(of(undefined))

        teamsService = jasmine.createSpyObj('TeamsService', ['getTeams'])
        teamsService.getTeams.and.returnValue(of(mockTeams))

        notificationService = jasmine.createSpyObj('NotificationService', [
            'showConfirm',
            'showError',
            'showSnackbar'
        ])

        await TestBed.configureTestingModule({
            imports: [SeasonDetailComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/:competition-id/seasons',
                        component: SeasonListComponent
                    }
                ]),
                { provide: SeasonsService, useValue: seasonsService },
                { provide: TeamsService, useValue: teamsService },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()
    })

    describe('Create mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, {
                useValue: mockRoute('comp1', null)
            })
            fixture = TestBed.createComponent(SeasonDetailComponent)
            component = fixture.componentInstance
            router = TestBed.inject(Router)
            fixture.detectChanges()
        })

        it('should create', () => {
            expect(component).toBeTruthy()
        })

        it('should mark start_date, start_time, end_date, end_time and rounds as required', () => {
            const startControl = component.seasonForm.get('start_date')
            const startTimeControl = component.seasonForm.get('start_time')
            const endControl = component.seasonForm.get('end_date')
            const endTimeControl = component.seasonForm.get('end_time')
            const roundsControl = component.seasonForm.get('rounds')

            startControl?.setValue('')
            expect(startControl?.valid).toBeFalse()
            startControl?.setValue(mockSeason1.start_date)
            expect(startControl?.valid).toBeTrue()

            startTimeControl?.setValue(null)
            expect(startTimeControl?.valid).toBeFalse()
            startTimeControl?.setValue(mockSeason1.start_date)
            expect(startTimeControl?.valid).toBeTrue()

            endControl?.setValue('')
            expect(endControl?.valid).toBeFalse()
            endControl?.setValue(mockSeason1.end_date)
            expect(endControl?.valid).toBeTrue()

            endTimeControl?.setValue(null)
            expect(endTimeControl?.valid).toBeFalse()
            endTimeControl?.setValue(mockSeason1.end_date)
            expect(endTimeControl?.valid).toBeTrue()

            roundsControl?.setValue('')
            expect(roundsControl?.valid).toBeFalse()
            roundsControl?.setValue(mockSeason1.rounds)
            expect(roundsControl?.valid).toBeTrue()
        })

        it('should not submit if form is invalid', () => {
            spyOn(console, 'error')
            component.submitForm()
            expect(console.error).toHaveBeenCalledWith('Season form is invalid')
        })

        it('should combine start_date and start_time correctly on submit', () => {
            const startDate = new Date('2025-01-01T00:00:00')
            const startTime = new Date('1970-01-01T13:30:00')
            const endDate = new Date('2025-12-31T00:00:00')
            const endTime = new Date('1970-01-01T15:45:00')

            component.isEditMode = false
            component.seasonForm.setValue({
                start_date: startDate,
                start_time: startTime,
                end_date: endDate,
                end_time: endTime,
                rounds: 5,
                teams: ['team1']
            })

            const expectedStart = (component as any).combineDateAndTime(startDate, startTime)
            const expectedEnd = (component as any).combineDateAndTime(endDate, endTime)

            component.submitForm()

            const calledSeason = seasonsService.createSeason.calls.mostRecent().args[1]
            expect(calledSeason.start_date?.getTime()).toBe(expectedStart.getTime())
            expect(calledSeason.end_date?.getTime()).toBe(expectedEnd.getTime())
        })

        it('should mark rounds invalid if less than 1 or greater than 50', () => {
            const rounds = component.seasonForm.get('rounds')
            rounds?.setValue(0)
            expect(rounds?.valid).toBeFalse()
            rounds?.setValue(51)
            expect(rounds?.valid).toBeFalse()
        })

        it('should not submit if competitionId is null', () => {
            spyOn(console, 'error')
            component.competitionId = null
            component.seasonForm.setValue({
                start_date: mockSeason1.start_date,
                start_time: new Date('1970-01-01T13:30:00Z'),
                end_date: mockSeason1.end_date,
                end_time: new Date('1970-01-01T15:30:00Z'),
                rounds: mockSeason1.rounds,
                teams: ['team1', 'team2']
            })
            component.submitForm()
            expect(console.error).toHaveBeenCalledWith('Season form is invalid')
        })

        it('should call createSeason when in create mode and form is valid', () => {
            const startDate = new Date('2025-01-01T00:00:00')
            const startTime = new Date('1970-01-01T13:00:00')
            const endDate = new Date('2026-01-02T00:00:00')
            const endTime = new Date('1970-01-01T01:59:00')

            component.isEditMode = false
            component.seasonForm.setValue({
                start_date: startDate,
                start_time: startTime,
                end_date: endDate,
                end_time: endTime,
                rounds: 20,
                teams: ['team1', 'team2']
            })

            const expectedStart = (component as any).combineDateAndTime(startDate, startTime)
            const expectedEnd = (component as any).combineDateAndTime(endDate, endTime)

            component.submitForm()

            expect(seasonsService.createSeason).toHaveBeenCalledWith(
                'comp1',
                jasmine.objectContaining({
                    start_date: expectedStart,
                    end_date: expectedEnd,
                    rounds: 20,
                    teams: ['team1', 'team2']
                })
            )
        })

        it('should log error if createSeason fails', () => {
            const error = new Error('Failed to create')
            spyOn(console, 'error')
            seasonsService.createSeason.and.returnValue(throwError(() => error))

            component.isEditMode = false
            component.seasonForm.patchValue({
                start_date: mockSeason1.start_date,
                start_time: new Date('1970-01-01T13:30:00Z'),
                end_date: mockSeason1.end_date,
                end_time: new Date('1970-01-01T15:30:00Z'),
                rounds: mockSeason1.rounds
            })
            component.submitForm()

            expect(console.error).toHaveBeenCalledWith('Error creating season:', error)
        })

        it('should navigate after createSeason', () => {
            const routerSpy = spyOn(router, 'navigateByUrl')

            component.isEditMode = false
            component.seasonForm.setValue({
                start_date: mockSeason1.start_date,
                start_time: new Date('1970-01-01T13:30:00Z'),
                end_date: mockSeason1.end_date,
                end_time: new Date('1970-01-01T15:30:00Z'),
                rounds: mockSeason1.rounds,
                teams: ['team1', 'team2']
            })
            component.submitForm()

            const call = routerSpy.calls.all()[0].args[0].toString()
            expect(call).toEqual('/admin/competitions/comp1/seasons')
        })
    })

    describe('Edit mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute('comp1', 'season1') })
            fixture = TestBed.createComponent(SeasonDetailComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should load season and teams when in edit mode', () => {
            expect(seasonsService.getSeason).toHaveBeenCalledWith('comp1', 'season1')
            expect(component.seasonForm.value.start_date).toEqual(mockSeason1.start_date)
            expect(component.seasonForm.value.end_date).toEqual(mockSeason1.end_date)
            expect(component.seasonForm.value.rounds).toEqual(mockSeason1.rounds)
            expect(component.teams).toEqual(mockTeams)
        })

        it('should separate start_date and start_time correctly when loading a season', () => {
            const seasonFromApi: Season = {
                id: 'season1',
                competition_id: 'comp1',
                rounds: 3,
                start_date: new Date('2025-01-01T09:30:00Z'),
                end_date: new Date('2025-12-31T18:45:00Z'),
                teams: mockTeams,
                created_at: new Date(),
                updated_at: new Date()
            }

            // Spy getSeason to return our specific test season
            seasonsService.getSeason.and.returnValue(of(seasonFromApi))
            component['loadSeason']('comp1', 'season1')

            const formValue = component.seasonForm.value

            // Date part
            expect(formValue.start_date.getFullYear()).toBe(2025)
            expect(formValue.start_date.getMonth()).toBe(0)
            expect(formValue.start_date.getDate()).toBe(1)

            expect(formValue.end_date.getFullYear()).toBe(2026)
            expect(formValue.end_date.getMonth()).toBe(0)
            expect(formValue.end_date.getDate()).toBe(1)

            // Time part
            expect(formValue.start_time.getHours()).toBe(new Date(seasonFromApi.start_date).getHours())
            expect(formValue.start_time.getMinutes()).toBe(new Date(seasonFromApi.start_date).getMinutes())
            expect(formValue.end_time.getHours()).toBe(new Date(seasonFromApi.end_date).getHours())
            expect(formValue.end_time.getMinutes()).toBe(new Date(seasonFromApi.end_date).getMinutes())
        })

        it('should log error if loadSeason fails', () => {
            const error = new Error('Failed to load')
            spyOn(console, 'error')
            seasonsService.getSeason.and.returnValue(throwError(() => error))

            component['loadSeason']('123', '456') // trigger loadSeason again

            expect(console.error).toHaveBeenCalledWith('Error loading season:', error)
        })

        it('should populate teams control with IDs in edit mode', () => {
            component['loadSeason']('comp1', 'season1')
            expect(component.seasonForm.value.teams).toEqual(['team1', 'team2', 'team3', 'team4'])
        })

        it('should log error if loadTeams fails', () => {
            const error = new Error('Teams failed')
            spyOn(console, 'error')
            teamsService.getTeams.and.returnValue(throwError(() => error))

            component['loadTeams']()
            expect(console.error).toHaveBeenCalledWith('Error loading teams:', error)
        })

        it('should call updateSeason when in edit mode and form is valid', () => {
            const startDate = new Date('2025-01-01T00:00:00')
            const startTime = new Date('1970-01-01T13:00:00')
            const endDate = new Date('2026-01-02T00:00:00')
            const endTime = new Date('1970-01-01T01:59:00')

            component.isEditMode = true
            component.competitionId = 'comp1'
            ;(component as any).seasonId = 'season1'

            component.seasonForm.setValue({
                start_date: startDate,
                start_time: startTime,
                end_date: endDate,
                end_time: endTime,
                rounds: 20,
                teams: ['team1', 'team2']
            })

            const expectedStart = (component as any).combineDateAndTime(startDate, startTime)
            const expectedEnd = (component as any).combineDateAndTime(endDate, endTime)

            component.submitForm()

            expect(seasonsService.updateSeason).toHaveBeenCalledWith(
                'comp1',
                'season1',
                jasmine.objectContaining({
                    start_date: expectedStart,
                    end_date: expectedEnd,
                    rounds: 20,
                    teams: ['team1', 'team2']
                })
            )
        })

        it('should log error if updateSeason fails', () => {
            const error = new Error('Failed to update')
            spyOn(console, 'error')
            seasonsService.updateSeason.and.returnValue(throwError(() => error))

            component.seasonForm.patchValue({ rounds: 'Bad Update' })
            component.submitForm()

            expect(console.error).toHaveBeenCalledWith('Error updating season:', error)
        })

        it('should call deleteSeason when confirmed', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(seasonsService.deleteSeason).toHaveBeenCalledWith('comp1', 'season1')
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Season deleted successfully')
        })

        it('should not call deleteSeason when cancelled', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(false) } as any)
            component.confirmDelete()
            expect(seasonsService.deleteSeason).not.toHaveBeenCalled()
            expect(notificationService.showSnackbar).not.toHaveBeenCalled()
        })

        it('should show error if deleteSeason fails', () => {
            seasonsService.deleteSeason.and.returnValue(throwError(() => new Error('Failed')))
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(notificationService.showError).toHaveBeenCalledWith(
                'Delete Error',
                'Failed to delete season'
            )
        })
    })
})
