import { ComponentFixture, TestBed } from '@angular/core/testing'
import { GameDetailComponent } from './game-detail.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ActivatedRoute, provideRouter, Router } from '@angular/router'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { GameListComponent } from '../game-list/game-list.component'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { GamesService } from '../../services/games/games.service'
import { Game, GameStatus, Season, Team } from '../../types/api'
import { of, throwError } from 'rxjs'
import { NotificationService } from '../../services/notifications/notifications.service'

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

    function mockRoute(competitionId: string | null, seasonId: string | null, gameId: string | null) {
        return {
            snapshot: {
                paramMap: {
                    get: (key: string) => {
                        if (key === 'competition-id') return competitionId
                        if (key === 'season-id') return seasonId
                        if (key === 'game-id') return gameId
                        return null
                    }
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
            'showError',
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

        it('should mark round, date, home_team_id, away_team_id and status as required', () => {
            const roundControl = component.gameForm.get('round')
            const dateControl = component.gameForm.get('date')
            const timeControl = component.gameForm.get('time')
            const homeTeamControl = component.gameForm.get('home_team_id')
            const awayTeamControl = component.gameForm.get('away_team_id')
            const statusControl = component.gameForm.get('status')

            roundControl?.setValue(null)
            expect(roundControl?.valid).toBeFalse()
            roundControl?.setValue(1)
            expect(roundControl?.valid).toBeTrue()

            dateControl?.setValue(null)
            expect(dateControl?.valid).toBeFalse()
            dateControl?.setValue('2025-08-21')
            expect(dateControl?.valid).toBeTrue()

            timeControl?.setValue(null)
            expect(timeControl?.valid).toBeFalse()
            timeControl?.setValue('10:00')
            expect(timeControl?.valid).toBeTrue()

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

        it('should not submit if form is invalid', () => {
            spyOn(console, 'error')
            component.submitForm()
            expect(console.error).toHaveBeenCalledWith('game form is invalid')
        })

        it('should combine date and time correctly on submit', () => {
            const date = new Date('2025-01-01T00:00:00')
            const time = new Date('1970-01-01T13:30:00')

            component.isEditMode = false
            component.gameForm.setValue({
                round: mockGame1.round,
                date: date,
                time: time,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.SCHEDULED
            })

            const expectedStart = (component as any).combineDateAndTime(date, time)

            component.submitForm()

            const calledGame = gamesService.createGame.calls.mostRecent().args[2]
            expect(calledGame.date?.getTime()).toBe(expectedStart.getTime())
        })

        it('should not submit if competitionId is null', () => {
            spyOn(console, 'error')
            component.competitionId = null
            component.gameForm.setValue({
                round: 1,
                date: '2025-08-21',
                time: '10:00',
                home_team_id: 'team1',
                away_team_id: 'team2',
                home_score: 0,
                away_score: 0,
                status: 'scheduled'
            })

            component.submitForm()

            const call = (console.error as jasmine.Spy).calls.all()[0].args[0]
            expect(call).toEqual('game form is invalid')
        })

        it('should populate rounds based on season', () => {
            expect(component.rounds).toEqual([1, 2, 3])
        })

        it('should have home and away teams populated 2', () => {
            component.ngOnInit()
            expect(component.teams.length).toBe(2)
            expect(component.teams[0].name).toBe('Team One')
            expect(component.teams[1].name).toBe('Team Two')
        })

        it('should call createGame when in create mode and form is valid', () => {
            component.isEditMode = false
            component.gameForm.setValue({
                round: mockGame1.round,
                date: mockGame1.date,
                time: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.SCHEDULED
            })

            component.submitForm()

            expect(gamesService.createGame).toHaveBeenCalledWith('comp1', 'season1', {
                round: mockGame1.round,
                date: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.SCHEDULED,
                season_id: mockGame1.season_id
            })
        })

        it('should log error if createGame fails', () => {
            const error = new Error('Failed to create')
            spyOn(console, 'error')
            gamesService.createGame.and.returnValue(throwError(() => error))

            component.isEditMode = false
            component.gameForm.patchValue({
                round: mockGame1.round,
                date: mockGame1.date,
                time: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: 10,
                away_score: 5,
                status: GameStatus.SCHEDULED
            })
            component.submitForm()

            expect(console.error).toHaveBeenCalledWith('Error creating game:', error)
        })

        it('should navigate after createGame', () => {
            const routerSpy = spyOn(router, 'navigateByUrl')

            component.isEditMode = false
            component.gameForm.setValue({
                round: mockGame1.round,
                date: mockGame1.date,
                time: mockGame1.date,
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
                date: mockGame1.date,
                time: mockGame1.date,
                home_team_id: mockGame1.home_team_id,
                away_team_id: mockGame1.away_team_id,
                home_score: mockGame1.home_score,
                away_score: mockGame1.away_score,
                status: mockGame1.status
            })

            expect(component.teams).toEqual(mockTeams)
        })

        it('should separate date and time correctly when loading a game', () => {
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

            expect(formValue.date.getFullYear()).toBe(2025)
            expect(formValue.date.getMonth()).toBe(5)
            expect(formValue.date.getDate()).toBe(16)

            expect(formValue.time.getHours()).toBe(new Date(gameFromApi.date).getHours())
            expect(formValue.time.getMinutes()).toBe(new Date(gameFromApi.date).getMinutes())
        })

        it('should log error if loadGame fails', () => {
            const error = new Error('Failed to load')
            spyOn(console, 'error')
            gamesService.getGame.and.returnValue(throwError(() => error))

            component['loadGame']('123', '456', '789')

            expect(console.error).toHaveBeenCalledWith('Error loading game:', error)
        })

        it('should log error if loadSeason fails', () => {
            const error = new Error('Failed to load season')
            spyOn(console, 'error')
            seasonsService.getSeason.and.returnValue(throwError(() => error))

            component.ngOnInit()

            expect(console.error).toHaveBeenCalledWith('Error loading season:', error)
        })

        it('should call updateGame when in edit mode and form is valid', () => {
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

        it('should log error if updateGame fails', () => {
            const error = new Error('Failed to update')
            spyOn(console, 'error')
            gamesService.updateGame.and.returnValue(throwError(() => error))

            component.gameForm.patchValue({ round: 'Bad Update' })
            component.submitForm()

            expect(console.error).toHaveBeenCalledWith('Error updating game:', error)
        })

        it('should call deleteGame when confirmed', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(gamesService.deleteGame).toHaveBeenCalledWith('comp1', 'season1', 'game1')
            expect(notificationService.showSnackbar).toHaveBeenCalledWith('Game deleted successfully', 'OK')
        })

        it('should not call deleteGame when cancelled', () => {
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(false) } as any)
            component.confirmDelete()
            expect(gamesService.deleteGame).not.toHaveBeenCalled()
            expect(notificationService.showSnackbar).not.toHaveBeenCalled()
        })

        it('should show error if deleteGame fails', () => {
            gamesService.deleteGame.and.returnValue(throwError(() => new Error('Failed')))
            notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
            component.confirmDelete()
            expect(notificationService.showError).toHaveBeenCalledWith(
                'Delete Error',
                'Failed to delete game'
            )
        })
    })
})
