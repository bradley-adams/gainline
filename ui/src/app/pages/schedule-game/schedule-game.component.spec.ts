import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ComponentFixture, fakeAsync, TestBed, tick } from '@angular/core/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { ActivatedRoute } from '@angular/router'
import { of, Subject, throwError } from 'rxjs'

import { GamesService } from '../../services/games/games.service'
import { LiveGameService } from '../../services/live-game/live-game.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { Game, GameState, GameStatus, Season, StageType, Team } from '../../types/api'
import { ScheduleGameComponent } from './schedule-game.component'

describe('ScheduleGameComponent', () => {
    let component: ScheduleGameComponent
    let fixture: ComponentFixture<ScheduleGameComponent>

    let gamesService: jasmine.SpyObj<GamesService>
    let seasonsService: jasmine.SpyObj<SeasonsService>
    let liveGameService: jasmine.SpyObj<LiveGameService>
    let notificationService: jasmine.SpyObj<NotificationService>

    const mockTeams: Team[] = [
        {
            id: 'team1',
            name: 'Home Team',
            abbreviation: 'HT',
            location: 'City A',
            created_at: new Date('2024-01-01T00:00:00Z'),
            updated_at: new Date('2024-01-01T00:00:00Z')
        },
        {
            id: 'team2',
            name: 'Away Team',
            abbreviation: 'AT',
            location: 'City B',
            created_at: new Date('2024-01-01T00:00:00Z'),
            updated_at: new Date('2024-01-01T00:00:00Z')
        }
    ]

    const mockSeason: Season = {
        id: 'season1',
        competition_id: 'comp1',
        start_date: new Date('2025-01-01T00:00:00Z'),
        end_date: new Date('2025-12-31T23:59:59Z'),
        stages: [
            {
                id: 'stage1',
                name: 'Round 1',
                stage_type: StageType.StageTypeRegular,
                order_index: 1,
                created_at: new Date('2025-01-01T00:00:00Z'),
                updated_at: new Date('2025-01-01T00:00:00Z')
            }
        ],
        teams: mockTeams,
        created_at: new Date('2024-12-01T00:00:00Z'),
        updated_at: new Date('2024-12-01T00:00:00Z')
    }

    const mockGame: Game = {
        id: 'game1',
        season_id: 'season1',
        stage_id: 'stage1',
        date: new Date('2025-02-01T15:00:00Z'),
        home_team_id: 'team1',
        away_team_id: 'team2',
        home_score: 1,
        away_score: 0,
        status: GameStatus.FINISHED,
        created_at: new Date('2025-01-20T10:00:00Z'),
        updated_at: new Date('2025-02-01T17:00:00Z')
    }

    const mockPlayingGame: Game = {
        ...mockGame,
        status: GameStatus.PLAYING,
        home_score: 0,
        away_score: 0
    }

    const mockGameState: GameState = {
        game_id: 'game1',
        home_score: 2,
        away_score: 1,
        status: 'playing',
        minute: 55
    }

    const mockRoute = {
        snapshot: {
            paramMap: {
                get: (key: string) => {
                    const params: Record<string, string> = {
                        'competition-id': 'comp1',
                        'season-id': 'season1',
                        'game-id': 'game1'
                    }
                    return params[key] ?? null
                }
            }
        }
    }

    beforeEach(async () => {
        gamesService = jasmine.createSpyObj('GamesService', ['getGame'])
        seasonsService = jasmine.createSpyObj('SeasonsService', ['getSeason'])
        liveGameService = jasmine.createSpyObj('LiveGameService', ['watchGame'])
        notificationService = jasmine.createSpyObj('NotificationService', ['showErrorAndLog'])

        gamesService.getGame.and.returnValue(of(mockGame))
        seasonsService.getSeason.and.returnValue(of(mockSeason))

        await TestBed.configureTestingModule({
            imports: [ScheduleGameComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                { provide: ActivatedRoute, useValue: mockRoute },
                { provide: GamesService, useValue: gamesService },
                { provide: SeasonsService, useValue: seasonsService },
                { provide: LiveGameService, useValue: liveGameService },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(ScheduleGameComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should load game on init', () => {
        expect(gamesService.getGame).toHaveBeenCalledWith('comp1', 'season1', 'game1')
        expect(component.game).toEqual(mockGame)
    })

    it('should load season on init', () => {
        expect(seasonsService.getSeason).toHaveBeenCalledWith('comp1', 'season1')
    })

    it('should resolve team names from season', () => {
        expect(component.homeTeamName).toBe('Home Team')
        expect(component.awayTeamName).toBe('Away Team')
    })

    it('should not open live connection for a finished game', () => {
        expect(liveGameService.watchGame).not.toHaveBeenCalled()
    })

    it('should open live connection when game is playing', fakeAsync(() => {
        const liveSubject = new Subject<GameState>()
        gamesService.getGame.and.returnValue(of(mockPlayingGame))
        liveGameService.watchGame.and.returnValue(liveSubject.asObservable())

        fixture = TestBed.createComponent(ScheduleGameComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
        tick()

        expect(liveGameService.watchGame).toHaveBeenCalledWith('comp1', 'season1', 'game1')
    }))

    it('should update live state when SSE event arrives', fakeAsync(() => {
        const liveSubject = new Subject<GameState>()
        gamesService.getGame.and.returnValue(of(mockPlayingGame))
        liveGameService.watchGame.and.returnValue(liveSubject.asObservable())

        fixture = TestBed.createComponent(ScheduleGameComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
        tick()

        liveSubject.next(mockGameState)
        fixture.detectChanges()

        expect(component.liveState).toEqual(mockGameState)
    }))

    it('should prefer live state scores over db scores', fakeAsync(() => {
        const liveSubject = new Subject<GameState>()
        gamesService.getGame.and.returnValue(of(mockPlayingGame))
        liveGameService.watchGame.and.returnValue(liveSubject.asObservable())

        fixture = TestBed.createComponent(ScheduleGameComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
        tick()

        liveSubject.next(mockGameState)
        fixture.detectChanges()

        expect(component.homeScore).toBe(mockGameState.home_score)
        expect(component.awayScore).toBe(mockGameState.away_score)
    }))

    it('should fall back to db scores when no live state', () => {
        expect(component.homeScore).toBe(mockGame.home_score)
        expect(component.awayScore).toBe(mockGame.away_score)
    })

    it('should show error notification if loading game fails', () => {
        const mockError = new Error('Failed')
        gamesService.getGame.and.returnValue(throwError(() => mockError))

        component.ngOnInit()

        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Load Error',
            'Failed to load game',
            mockError
        )
    })

    it('should show error notification if season loading fails', () => {
        const mockError = new Error('Failed')
        seasonsService.getSeason.and.returnValue(throwError(() => mockError))

        component.ngOnInit()

        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Load Error',
            'Failed to load season',
            mockError
        )
    })

    it('should show error notification if live connection fails', fakeAsync(() => {
        const mockError = new Error('SSE error')
        gamesService.getGame.and.returnValue(of(mockPlayingGame))
        liveGameService.watchGame.and.returnValue(throwError(() => mockError))

        fixture = TestBed.createComponent(ScheduleGameComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
        tick()

        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Live Error',
            'Lost live connection',
            mockError
        )
    }))

    it('should unsubscribe from live state on destroy', fakeAsync(() => {
        const liveSubject = new Subject<GameState>()
        gamesService.getGame.and.returnValue(of(mockPlayingGame))
        liveGameService.watchGame.and.returnValue(liveSubject.asObservable())

        fixture = TestBed.createComponent(ScheduleGameComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
        tick()

        fixture.destroy()

        liveSubject.next(mockGameState)

        expect(component.liveState).toBeNull()
    }))
})
