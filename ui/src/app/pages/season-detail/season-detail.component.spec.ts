import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ComponentFixture, TestBed } from '@angular/core/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { ActivatedRoute, provideRouter, Router } from '@angular/router'
import { of, throwError } from 'rxjs'

import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { TeamsService } from '../../services/teams/teams.service'
import { Season, Team } from '../../types/api'
import { SeasonListComponent } from '../season-list/season-list.component'
import { SeasonDetailComponent } from './season-detail.component'

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
        start_date: new Date('2025-01-01T00:00:00Z'),
        end_date: new Date('2026-01-01T12:59:00Z'),
        teams: mockTeams,
        created_at: new Date('2024-12-01T00:00:00Z'),
        updated_at: new Date('2024-12-01T00:00:00Z')
    }

    const mockSeason2: Season = {
        id: 'season2',
        competition_id: 'comp1',
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
                    get: (key: string) =>
                        ({ 'competition-id': competitionId, 'season-id': seasonId })[key] ?? null
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
            'showErrorAndLog',
            'showWarnAndLog',
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

        it('should mark teams invalid if fewer than 2 selected', () => {
            const teams = component.seasonForm.get('teams')
            teams?.setValue([])
            expect(teams?.valid).toBeFalse()
            teams?.setValue(['team1'])
            expect(teams?.valid).toBeFalse()
            teams?.setValue(['team1', 'team2'])
            expect(teams?.valid).toBeTrue()
        })

        it('should mark form invalid if end_datetime is before start_datetime', () => {
            const start = new Date('2025-01-02T10:00:00')
            const end = new Date('2025-01-01T09:00:00')

            component.seasonForm.setValue({
                start_datetime: start,
                end_datetime: end,
                teams: ['team1', 'team2']
            })

            expect(component.seasonForm.valid).toBeFalse()
            expect(component.seasonForm.errors?.['endBeforeStart']).toBeTrue()
        })

        it('should mark teams invalid if more than max allowed', () => {
            const validList = Array.from({ length: 20 }, (_, i) => `${i}`)

            component.seasonForm.get('teams')?.setValue(validList)
            expect(component.seasonForm.get('teams')?.valid).toBeTrue()

            const invalidList = [...validList, 'extra']

            component.seasonForm.get('teams')?.setValue(invalidList)
            expect(component.seasonForm.get('teams')?.valid).toBeFalse()
        })

        it('should require start/end datetimes and rounds', () => {
            const startControl = component.seasonForm.get('start_datetime')
            const endControl = component.seasonForm.get('end_datetime')

            // Start datetime required
            startControl?.setValue(null)
            expect(startControl?.valid).toBeFalse()
            startControl?.setValue(new Date('2025-01-01T10:00:00'))
            expect(startControl?.valid).toBeTrue()

            // End datetime required
            endControl?.setValue(null)
            expect(endControl?.valid).toBeFalse()
            endControl?.setValue(new Date('2025-01-02T10:00:00'))
            expect(endControl?.valid).toBeTrue()
        })

        it('should not call createSeason or updateSeason if form is invalid', () => {
            component.seasonForm.setValue({
                start_datetime: null,
                end_datetime: null,
                teams: []
            })
            component.isEditMode = false

            component.submitForm()

            expect(seasonsService.createSeason).not.toHaveBeenCalled()
            expect(seasonsService.updateSeason).not.toHaveBeenCalled()
        })

        it('should pass start_datetime and end_datetime correctly on submit', () => {
            const startDatetime = new Date('2025-01-01T13:30:00')
            const endDatetime = new Date('2025-12-31T15:45:00')

            component.isEditMode = false
            component.competitionId = 'comp1'
            component.seasonForm.setValue({
                start_datetime: startDatetime,
                end_datetime: endDatetime,
                teams: ['team1', 'team2']
            })

            component.submitForm()

            const calledSeason = seasonsService.createSeason.calls.mostRecent().args[1]
            expect(calledSeason.start_date?.getTime()).toBe(startDatetime.getTime())
            expect(calledSeason.end_date?.getTime()).toBe(endDatetime.getTime())
        })

        it('should not submit if competitionId is null', () => {
            component.competitionId = null
            component.seasonForm.setValue({
                start_datetime: mockSeason1.start_date,
                end_datetime: mockSeason1.end_date,
                teams: ['team1', 'team2']
            })

            component.submitForm()

            expect(notificationService.showWarnAndLog).toHaveBeenCalledWith(
                'Form Error',
                'No competition selected'
            )
        })
        it('should call createSeason when in create mode and form is valid', () => {
            const startDatetime = new Date('2025-01-01T13:00:00')
            const endDatetime = new Date('2026-01-02T01:59:00')

            component.isEditMode = false
            component.competitionId = 'comp1'
            component.seasonForm.setValue({
                start_datetime: startDatetime,
                end_datetime: endDatetime,
                teams: ['team1', 'team2']
            })

            component.submitForm()

            expect(seasonsService.createSeason).toHaveBeenCalledWith(
                'comp1',
                jasmine.objectContaining({
                    start_date: startDatetime,
                    end_date: endDatetime,
                    teams: ['team1', 'team2']
                })
            )
        })

        it('should show error if createSeason fails', () => {
            const mockError = new Error('Failed to create')
            seasonsService.createSeason.and.returnValue(throwError(() => mockError))

            component.isEditMode = false
            component.competitionId = 'comp1'
            component.seasonForm.patchValue({
                start_datetime: mockSeason1.start_date,
                end_datetime: mockSeason1.end_date,
                teams: ['team1', 'team2']
            })

            component.submitForm()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Create Error',
                'Failed to create season',
                mockError
            )
        })

        it('should navigate after createSeason', () => {
            const routerSpy = spyOn(router, 'navigate')
            component.isEditMode = false
            component.competitionId = 'comp1'
            component.seasonForm.setValue({
                start_datetime: mockSeason1.start_date,
                end_datetime: mockSeason1.end_date,
                teams: ['team1', 'team2']
            })

            component.submitForm()

            expect(routerSpy).toHaveBeenCalledWith(['/admin/competitions', 'comp1', 'seasons'])
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
            expect(component.seasonForm.value.start_datetime).toEqual(mockSeason1.start_date)
            expect(component.seasonForm.value.end_datetime).toEqual(mockSeason1.end_date)
            expect(component.teams).toEqual(mockTeams)
        })

        it('should load season correctly into start_datetime and end_datetime', () => {
            const seasonFromApi: Season = {
                id: 'season1',
                competition_id: 'comp1',
                start_date: new Date('2025-01-01T09:30:00Z'),
                end_date: new Date('2025-12-31T18:45:00Z'),
                teams: mockTeams,
                created_at: new Date(),
                updated_at: new Date()
            }

            seasonsService.getSeason.and.returnValue(of(seasonFromApi))
            component['loadSeason']('comp1', 'season1')

            const formValue = component.seasonForm.value

            expect(formValue.start_datetime.getTime()).toBe(new Date(seasonFromApi.start_date).getTime())
            expect(formValue.end_datetime.getTime()).toBe(new Date(seasonFromApi.end_date).getTime())
        })

        it('should show error if loadSeason fails', () => {
            const mockError = new Error('Failed to load')
            seasonsService.getSeason.and.returnValue(throwError(() => mockError))
            component['loadSeason']('123', '456') // trigger loadSeason again
            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Load Error',
                'Failed to load season',
                mockError
            )
        })

        it('should populate teams control with IDs in edit mode', () => {
            component['loadSeason']('comp1', 'season1')
            expect(component.seasonForm.value.teams).toEqual(['team1', 'team2', 'team3', 'team4'])
        })

        it('should show error if loadTeams fails', () => {
            const mockError = new Error('Teams failed')
            teamsService.getTeams.and.returnValue(throwError(() => mockError))
            component['loadTeams']()
            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Load Error',
                'Failed to load teams',
                mockError
            )
        })

        it('should call updateSeason when in edit mode and form is valid', () => {
            const startDatetime = new Date('2025-01-01T13:00:00')
            const endDatetime = new Date('2026-01-02T01:59:00')

            component.isEditMode = true
            component.competitionId = 'comp1'
            ;(component as any).seasonId = 'season1'

            component.seasonForm.setValue({
                start_datetime: startDatetime,
                end_datetime: endDatetime,
                teams: ['team1', 'team2']
            })

            component.submitForm()

            expect(seasonsService.updateSeason).toHaveBeenCalledWith(
                'comp1',
                'season1',
                jasmine.objectContaining({
                    start_date: startDatetime,
                    end_date: endDatetime,
                    teams: ['team1', 'team2']
                })
            )
        })

        it('should show error if updateSeason fails', () => {
            const mockError = new Error('Failed to update')
            seasonsService.updateSeason.and.returnValue(throwError(() => mockError))
            component.seasonForm.patchValue({
                start_date: mockSeason1.start_date,
                start_time: new Date('1970-01-01T13:30:00Z'),
                end_date: mockSeason1.end_date,
                end_time: new Date('1970-01-01T15:30:00Z'),
                teams: ['team1', 'team2']
            })
            component.submitForm()
            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Update Error',
                'Failed to update season',
                mockError
            )
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
            const mockError = new Error('Failed')
            seasonsService.deleteSeason.and.returnValue(throwError(() => mockError))
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Delete Error',
                'Failed to delete season',
                mockError
            )
        })
    })
})
