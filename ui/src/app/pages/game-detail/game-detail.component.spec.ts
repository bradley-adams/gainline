import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ComponentFixture, TestBed } from '@angular/core/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { ActivatedRoute, provideRouter, Router } from '@angular/router'
import { of, throwError } from 'rxjs'

import { GameDetailComponent } from './game-detail.component'
import { GameListComponent } from '../game-list/game-list.component'
import { GamesService } from '../../services/games/games.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { Game, GameStatus, Season, Team } from '../../types/api'
import { Validators } from '@angular/forms'

describe('GameDetailComponent', () => {
    let component: GameDetailComponent
    let fixture: ComponentFixture<GameDetailComponent>
    let router: Router

    let seasonsService: jasmine.SpyObj<SeasonsService>
    let gamesService: jasmine.SpyObj<GamesService>
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
        }
    ]

    const mockSeason1: Season = {
        id: 'season1',
        competition_id: 'comp1',
        rounds: 3,
        start_date: new Date('2025-01-01T00:00:00Z'),
        end_date: new Date('2025-12-31T23:59:59Z'),
        teams: mockTeams,
        created_at: new Date('2024-12-01T00:00:00Z'),
        updated_at: new Date('2024-12-01T00:00:00Z')
    }

    const mockGame1: Game = {
        id: 'game1',
        season_id: 'season1',
        round: 1,
        date: new Date('2025-02-01T15:00:00Z'),
        home_team_id: 'team1',
        away_team_id: 'team2',
        home_score: 2,
        away_score: 1,
        status: GameStatus.FINISHED,
        created_at: new Date('2025-01-20T10:00:00Z'),
        updated_at: new Date('2025-02-01T17:00:00Z')
    }

    const mockGame2: Game = {
        id: 'game2',
        season_id: 'season1',
        round: 2,
        date: new Date('2025-03-01T15:00:00Z'),
        home_team_id: 'team3',
        away_team_id: 'team4',
        home_score: 0,
        away_score: 0,
        status: GameStatus.SCHEDULED,
        created_at: new Date('2025-02-20T10:00:00Z'),
        updated_at: new Date('2025-02-21T10:00:00Z')
    }

    function mockRoute(compId: string | null, seasonId: string | null, gameId: string | null) {
        return {
            snapshot: {
                paramMap: {
                    get: (key: string) =>
                        ({ 'competition-id': compId, 'season-id': seasonId, 'game-id': gameId })[key] ?? null
                }
            }
        }
    }

    beforeEach(async () => {
        gamesService = jasmine.createSpyObj('GamesService', [
            'getGame',
            'createGame',
            'updateGame',
            'deleteGame'
        ])
        gamesService.getGame.and.returnValue(of(mockGame1))
        gamesService.createGame.and.returnValue(of(mockGame1))
        gamesService.updateGame.and.returnValue(of(mockGame2))
        gamesService.deleteGame.and.returnValue(of(undefined))

        seasonsService = jasmine.createSpyObj('SeasonsService', ['getSeason'])
        seasonsService.getSeason.and.returnValue(of(mockSeason1))

        notificationService = jasmine.createSpyObj('NotificationService', [
            'showConfirm',
            'showErrorAndLog',
            'showWarnAndLog',
            'showSnackbar'
        ])

        await TestBed.configureTestingModule({
            imports: [GameDetailComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/:competition-id/seasons/:season-id/games',
                        component: GameListComponent
                    }
                ]),
                { provide: SeasonsService, useValue: seasonsService },
                { provide: GamesService, useValue: gamesService },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()
    })

    describe('Create mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, {
                useValue: mockRoute('comp1', 'season1', null)
            })
            fixture = TestBed.createComponent(GameDetailComponent)
            component = fixture.componentInstance
            router = TestBed.inject(Router)
            fixture.detectChanges()
        })

        it('should create', () => {
            expect(component).toBeTruthy()
        })

        it('should mark round, datetime, home_team_id, away_team_id and status as required', () => {
            const roundControl = component.gameForm.get('round')
            const datetimeControl = component.gameForm.get('datetime')
            const homeTeamControl = component.gameForm.get('home_team_id')
            const awayTeamControl = component.gameForm.get('away_team_id')
            const statusControl = component.gameForm.get('status')

            roundControl?.setValue(null)
            expect(roundControl?.valid).toBeFalse()
            roundControl?.setValue(1)
            expect(roundControl?.valid).toBeTrue()

            datetimeControl?.setValue(null)
            expect(datetimeControl?.valid).toBeFalse()
            datetimeControl?.setValue(new Date('2025-08-21T10:00:00'))
            expect(datetimeControl?.valid).toBeTrue()

            homeTeamControl?.setValue(null)
            expect(homeTeamControl?.valid).toBeFalse()
            homeTeamControl?.setValue('team1')
            expect(homeTeamControl?.valid).toBeTrue()

            awayTeamControl?.setValue(null)
            expect(awayTeamControl?.valid).toBeFalse()
            awayTeamControl?.setValue('team2')
            expect(awayTeamControl?.valid).toBeTrue()

            statusControl?.setValue(null)
            expect(statusControl?.valid).toBeFalse()
            statusControl?.setValue('scheduled')
            expect(statusControl?.valid).toBeTrue()
        })

        it('should mark form as invalid if datetime is outside the season', () => {
            component.seasonStart = new Date('2025-01-01T00:00:00Z')
            component.seasonEnd = new Date('2025-11-30T23:59:59Z')

            // Before season
            component.gameForm.patchValue({ datetime: new Date('2024-12-31T12:00:00Z') })
            expect(component.gameForm.errors).toEqual({ outOfSeason: true })

            // After season
            component.gameForm.patchValue({ datetime: new Date('2025-12-31T12:00:00Z') })
            expect(component.gameForm.errors).toEqual({ outOfSeason: true })

            // Within season
            component.gameForm.patchValue({ datetime: new Date('2025-06-15T12:00:00Z') })
            expect(component.gameForm.errors).toBeNull()
        })

        it('should show warning if competitionId or seasonId is missing', () => {
            component.competitionId = null
            component.submitForm()

            expect(notificationService.showWarnAndLog).toHaveBeenCalledWith(
                'Form Error',
                'Game form is invalid or competition/season not selected'
            )
        })

        it('should use datetime directly on submit', () => {
            const datetime = new Date('2025-01-01T13:30:00')

            component.isEditMode = false
            component.gameForm.setValue({
                round: mockGame1.round,
                datetime: datetime,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.SCHEDULED
            })

            component.submitForm()

            const calledGame = gamesService.createGame.calls.mostRecent().args[2]
            expect(calledGame.date?.getTime()).toBe(datetime.getTime())
        })

        it('should not submit if competitionId is null', () => {
            component.competitionId = null
            component.gameForm.setValue({
                round: 1,
                datetime: new Date('2025-08-21T10:00:00'),
                home_team_id: 'team1',
                away_team_id: 'team2',
                home_score: 0,
                away_score: 0,
                status: 'scheduled'
            })

            component.submitForm()

            expect(notificationService.showWarnAndLog).toHaveBeenCalledWith(
                'Form Error',
                'Game form is invalid or competition/season not selected'
            )
        })

        it('should populate rounds based on season', () => {
            expect(component.rounds).toEqual([1, 2, 3])
        })

        it('should populate home and away teams', () => {
            component.ngOnInit()
            expect(component.teams.length).toBe(2)
            expect(component.teams[0].name).toBe('Team One')
            expect(component.teams[1].name).toBe('Team Two')
        })

        it('should remove required validator from scores when status is "scheduled"', () => {
            const homeScoreCtrl = component.gameForm.get('home_score')!
            const awayScoreCtrl = component.gameForm.get('away_score')!

            component.gameForm.patchValue({ status: 'playing' })
            expect(homeScoreCtrl.hasValidator(Validators.required)).toBeTrue()
            expect(awayScoreCtrl.hasValidator(Validators.required)).toBeTrue()

            component.gameForm.patchValue({ status: 'scheduled' })
            expect(homeScoreCtrl.hasValidator(Validators.required)).toBeFalse()
            expect(awayScoreCtrl.hasValidator(Validators.required)).toBeFalse()
        })

        it('should add required validator when status is "playing" or "finished"', () => {
            const homeScoreCtrl = component.gameForm.get('home_score')!
            const awayScoreCtrl = component.gameForm.get('away_score')!

            component.gameForm.patchValue({ status: 'scheduled' })
            expect(homeScoreCtrl.hasValidator(Validators.required)).toBeFalse()

            component.gameForm.patchValue({ status: 'playing' })
            expect(homeScoreCtrl.hasValidator(Validators.required)).toBeTrue()

            component.gameForm.patchValue({ status: 'finished' })
            expect(awayScoreCtrl.hasValidator(Validators.required)).toBeTrue()
        })

        it('should not include scores when status is "scheduled"', () => {
            component.isEditMode = false
            component.gameForm.patchValue({
                round: 1,
                datetime: new Date('2025-01-01T13:30:00'),
                home_team_id: 'team1',
                away_team_id: 'team2',
                home_score: 5,
                away_score: 4,
                status: 'scheduled'
            })

            component.submitForm()

            const calledGame = gamesService.createGame.calls.mostRecent().args[2]
            expect(calledGame.home_score).toBeNull()
            expect(calledGame.away_score).toBeNull()
        })

        it('should call createGame when in create mode and form is valid', () => {
            component.isEditMode = false
            component.gameForm.setValue({
                round: mockGame1.round,
                datetime: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.FINISHED
            })

            component.submitForm()

            expect(gamesService.createGame).toHaveBeenCalledWith('comp1', 'season1', {
                round: mockGame1.round,
                date: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.FINISHED,
                season_id: mockGame1.season_id
            })
        })

        it('should show error if createGame fails', () => {
            const mockError = new Error('Failed to create')
            gamesService.createGame.and.returnValue(throwError(() => mockError))

            component.isEditMode = false
            component.gameForm.patchValue({
                round: mockGame1.round,
                datetime: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.SCHEDULED
            })
            component.submitForm()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Create Error',
                'Failed to create game',
                mockError
            )
        })

        it('should navigate to games list after successful createGame', () => {
            const routerSpy = spyOn(router, 'navigateByUrl')

            component.isEditMode = false
            component.gameForm.setValue({
                round: mockGame1.round,
                datetime: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.SCHEDULED
            })
            component.submitForm()

            const call = routerSpy.calls.all()[0].args[0].toString()
            expect(call).toEqual('/admin/competitions/comp1/seasons/season1/games')
        })
    })

    describe('Edit mode', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute('comp1', 'season1', 'game1') })
            fixture = TestBed.createComponent(GameDetailComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should load game when in edit mode', () => {
            expect(gamesService.getGame).toHaveBeenCalledWith('comp1', 'season1', 'game1')

            expect(component.gameForm.value).toEqual({
                round: mockGame1.round,
                datetime: new Date(mockGame1.date),
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: mockGame1.home_score,
                away_score: mockGame1.away_score,
                status: mockGame1.status
            })

            expect(component.teams).toEqual(mockTeams)
        })

        it('should load datetime correctly when loading a game', () => {
            const gameFromApi: Game = {
                id: 'game1',
                season_id: 'season1',
                round: 1,
                date: new Date('2025-06-15T14:45:00Z'),
                home_team_id: 'team1',
                away_team_id: 'team2',
                home_score: 3,
                away_score: 2,
                status: GameStatus.FINISHED,
                created_at: new Date('2025-05-01T10:00:00Z'),
                updated_at: new Date('2025-06-15T16:00:00Z')
            }

            gamesService.getGame.and.returnValue(of(gameFromApi))
            component['loadGame']('comp1', 'season1', 'game1')

            const formValue = component.gameForm.value

            expect(formValue.datetime instanceof Date).toBeTrue()
            expect(formValue.datetime.toISOString()).toBe(gameFromApi.date.toISOString())
        })

        it('should show error if loadGame fails', () => {
            const mockError = new Error('Failed to load')
            gamesService.getGame.and.returnValue(throwError(() => mockError))

            component['loadGame']('123', '456', '789')

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Load Error',
                'Failed to load game',
                mockError
            )
        })

        it('should show error if loadSeason fails', () => {
            const mockError = new Error('Failed to load season')
            seasonsService.getSeason.and.returnValue(throwError(() => mockError))

            component.ngOnInit()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Load Error',
                'Failed to load season',
                mockError
            )
        })

        it('should update game when in edit mode and form is valid', () => {
            component.gameForm.patchValue({ round: 5 })

            component.submitForm()

            expect(gamesService.updateGame).toHaveBeenCalledWith('comp1', 'season1', 'game1', {
                round: 5,
                date: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: mockGame1.home_score,
                away_score: mockGame1.away_score,
                status: mockGame1.status,
                season_id: mockGame1.season_id
            })
        })

        it('should show error if updateGame fails', () => {
            const mockError = new Error('Failed to update')
            gamesService.updateGame.and.returnValue(throwError(() => mockError))

            component.gameForm.patchValue({ round: 'Bad Update' })
            component.submitForm()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Update Error',
                'Failed to update game',
                mockError
            )
        })

        it('should call deleteGame when confirmed', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(gamesService.deleteGame).toHaveBeenCalledWith('comp1', 'season1', 'game1')
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Game deleted successfully')
        })

        it('should not call deleteGame when cancelled', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(false) } as any)
            component.confirmDelete()
            expect(gamesService.deleteGame).not.toHaveBeenCalled()
            expect(notificationService.showSnackbar).not.toHaveBeenCalled()
        })

        it('should show error if deleteGame fails', () => {
            const mockError = new Error('Failed')
            gamesService.deleteGame.and.returnValue(throwError(() => mockError))

            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()

            expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
                'Delete Error',
                'Failed to delete game',
                mockError
            )
        })
    })
})
