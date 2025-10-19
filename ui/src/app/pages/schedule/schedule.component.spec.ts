import { ComponentFixture, TestBed } from '@angular/core/testing'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { of, throwError } from 'rxjs'

import { ScheduleComponent } from './schedule.component'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { GamesService } from '../../services/games/games.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { Competition, Game, GameStatus, Season, Team } from '../../types/api'

describe('ScheduleComponent', () => {
    let component: ScheduleComponent
    let fixture: ComponentFixture<ScheduleComponent>

    let competitionsService: jasmine.SpyObj<CompetitionsService>
    let seasonsService: jasmine.SpyObj<SeasonsService>
    let gamesService: jasmine.SpyObj<GamesService>
    let notificationService: jasmine.SpyObj<NotificationService>

    const mockCompetitions: Competition[] = [
        {
            id: 'comp1',
            name: 'Competition 1',
            created_at: new Date('2023-01-01T00:00:00Z'),
            updated_at: new Date('2023-01-02T00:00:00Z')
        },
        {
            id: 'comp2',
            name: 'Competition 2',
            created_at: new Date('2023-01-03T00:00:00Z'),
            updated_at: new Date('2023-01-04T00:00:00Z')
        }
    ]

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

    const mockSeasons: Season[] = [
        {
            id: 'season1',
            competition_id: 'comp1',
            rounds: 3,
            start_date: new Date('2025-01-01T00:00:00Z'),
            end_date: new Date('2025-12-31T23:59:59Z'),
            teams: mockTeams,
            created_at: new Date('2024-12-01T00:00:00Z'),
            updated_at: new Date('2024-12-01T00:00:00Z')
        },
        {
            id: 'season2',
            competition_id: 'comp1',
            rounds: 0,
            start_date: new Date('2024-01-01T00:00:00Z'),
            end_date: new Date('2024-12-31T23:59:59Z'),
            teams: mockTeams,
            created_at: new Date('2023-12-01T00:00:00Z'),
            updated_at: new Date('2023-12-01T00:00:00Z')
        }
    ]

    const mockGames: Game[] = [
        {
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
        },
        {
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
    ]

    beforeEach(async () => {
        competitionsService = jasmine.createSpyObj('CompetitionsService', ['getCompetitions'])
        competitionsService.getCompetitions.and.returnValue(of(mockCompetitions))

        seasonsService = jasmine.createSpyObj('SeasonsService', ['getSeasons'])
        seasonsService.getSeasons.and.returnValue(of(mockSeasons))

        gamesService = jasmine.createSpyObj('GamesService', ['getGames'])
        gamesService.getGames.and.returnValue(of(mockGames))

        notificationService = jasmine.createSpyObj('NotificationService', ['showErrorAndLog'])

        await TestBed.configureTestingModule({
            imports: [ScheduleComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                { provide: CompetitionsService, useValue: competitionsService },
                { provide: SeasonsService, useValue: seasonsService },
                { provide: GamesService, useValue: gamesService },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(ScheduleComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should load competitions on init and select the first competition', () => {
        expect(competitionsService.getCompetitions).toHaveBeenCalled()
        expect(component.competitions.length).toBe(2)
        expect(component.scheduleForm.get('competition')!.value).toBe('comp1')
    })

    it('should load seasons and select first season when competition changes', () => {
        const compId = mockCompetitions[1].id

        component.scheduleForm.get('competition')!.setValue(compId)

        expect(seasonsService.getSeasons).toHaveBeenCalledWith(compId)
        expect(component.seasons.length).toBe(mockSeasons.length)
        expect(component.scheduleForm.get('season')!.value).toBe(mockSeasons[0].id)
    })

    it('should load games and update dataSource when season or round changes', () => {
        const compId = mockCompetitions[0].id
        const seasonId = mockSeasons[0].id
        const round = 1

        component.scheduleForm.get('competition')!.setValue(compId)
        component.scheduleForm.get('season')!.setValue(seasonId)
        component.scheduleForm.get('round')!.setValue(round)

        expect(gamesService.getGames).toHaveBeenCalledWith(compId, seasonId)
        const filteredGames = component.games.filter((g) => g.round === round)
        expect(filteredGames.length).toBeGreaterThan(0)
        expect(component.games.every((g) => g.round === round)).toBeTrue()
        expect(component.dataSource.data.length).toBe(filteredGames.length)
        expect(component.dataSource.data).toEqual(filteredGames)
    })

    it('should reset rounds and clear round selection when season has zero rounds', () => {
        const compId = 'new-comp-id'

        seasonsService.getSeasons.and.callFake((id: string) => {
            if (id === 'comp1' || id === 'comp2') return of(mockSeasons)
            return of([])
        })

        component.scheduleForm.get('competition')!.setValue(compId)

        expect(seasonsService.getSeasons).toHaveBeenCalledWith(compId)
        expect(component.seasons.length).toBe(0)
        expect(component.rounds.length).toBe(0)
        expect(component.scheduleForm.get('round')!.value).toBe('')
        expect(component.games.length).toBe(0)
        expect(component.dataSource.data.length).toBe(0)
        expect(component.dataSource.data).toEqual([])
    })

    it('should display "No games found" message when no games are available', () => {
        component.games = []
        component.dataSource.data = []

        const noDataRow: HTMLElement = fixture.nativeElement.querySelector('tr.mat-row td.mat-cell')
        expect(noDataRow).toBeTruthy()
        expect(noDataRow.textContent).toContain('No games found')
        expect(component.dataSource.data.length).toBe(0)
        expect(component.dataSource.data).toEqual([])
    })

    it('should show error notification if loading competitions fails', () => {
        const mockError = new Error('Failed')
        competitionsService.getCompetitions.and.returnValue(throwError(() => mockError))

        component.ngOnInit()

        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Load Error',
            'Failed to load competitions',
            mockError
        )
    })
})
