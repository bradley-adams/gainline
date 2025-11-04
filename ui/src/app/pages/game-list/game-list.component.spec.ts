import { ComponentFixture, TestBed } from '@angular/core/testing'
import { By } from '@angular/platform-browser'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ActivatedRoute, convertToParamMap, provideRouter, Router } from '@angular/router'
import { of, throwError } from 'rxjs'

import { GameListComponent } from './game-list.component'
import { GameDetailComponent } from '../game-detail/game-detail.component'
import { GamesService } from '../../services/games/games.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { Game, GameStatus, Season, Team } from '../../types/api'

describe('GameListComponent', () => {
    let component: GameListComponent
    let fixture: ComponentFixture<GameListComponent>
    let router: Router

    let gamesService: jasmine.SpyObj<GamesService>
    let seasonsService: jasmine.SpyObj<SeasonsService>
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

    const mockSeason: Season = {
        id: 'season1',
        competition_id: 'comp1',
        rounds: 3,
        start_date: new Date('2025-01-01T00:00:00Z'),
        end_date: new Date('2025-12-31T23:59:59Z'),
        teams: mockTeams,
        created_at: new Date('2024-12-01T00:00:00Z'),
        updated_at: new Date('2024-12-01T00:00:00Z')
    }

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
        gamesService = jasmine.createSpyObj('GamesService', ['getGames', 'deleteGame'])
        seasonsService = jasmine.createSpyObj('SeasonsService', ['getSeason'])
        notificationService = jasmine.createSpyObj('NotificationService', [
            'showConfirm',
            'showErrorAndLog',
            'showWarnAndLog',
            'showSnackbar'
        ])

        gamesService.getGames.and.returnValue(of(mockGames))
        gamesService.deleteGame.and.returnValue(of(undefined))
        seasonsService.getSeason.and.returnValue(of(mockSeason))

        await TestBed.configureTestingModule({
            imports: [GameListComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/:competition-id/seasons/:season-id/games/create',
                        component: GameDetailComponent
                    },
                    {
                        path: 'admin/competitions/:competition-id/seasons/:season-id/games/:game-id',
                        component: GameDetailComponent
                    }
                ]),
                { provide: GamesService, useValue: gamesService },
                { provide: SeasonsService, useValue: seasonsService },
                {
                    provide: ActivatedRoute,
                    useValue: {
                        snapshot: {
                            paramMap: convertToParamMap({
                                'competition-id': 'comp1',
                                'season-id': 'season1'
                            })
                        }
                    }
                },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(GameListComponent)
        component = fixture.componentInstance
        router = TestBed.inject(Router)
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should load games with team names', () => {
        const rows = fixture.nativeElement.querySelectorAll('tr')
        expect(rows.length).toBe(3) // header + 2 games

        expect(rows[1].textContent).toContain('Team One')
        expect(rows[1].textContent).toContain('Team Two')
        expect(rows[1].textContent).toContain('finished')

        expect(rows[2].textContent).toContain('Team Three')
        expect(rows[2].textContent).toContain('Team Four')
        expect(rows[2].textContent).toContain('scheduled')
    })

    it('should display games table with correct headers and data', () => {
        const tableRows = fixture.nativeElement.querySelectorAll('tr')
        expect(tableRows.length).toBe(mockGames.length + 1)

        const headerRow = tableRows[0]
        expect(headerRow.cells[0].innerHTML).toBe('Date')
        expect(headerRow.cells[1].innerHTML).toBe('Home Team')
        expect(headerRow.cells[2].innerHTML).toBe('Home Score')
        expect(headerRow.cells[3].innerHTML).toBe('Away Score')
        expect(headerRow.cells[4].innerHTML).toBe('Away Team')
        expect(headerRow.cells[5].innerHTML).toBe('Status')
        expect(headerRow.cells[6].innerHTML).toBe('Actions')

        expect(tableRows[1].cells[0].textContent).toBe('02/02/2025 04:00')
        expect(tableRows[2].cells[0].textContent).toBe('02/03/2025 04:00')

        expect(tableRows[1].cells[1].textContent).toBe('Team One')
        expect(tableRows[2].cells[1].textContent).toBe('Team Three')

        expect(tableRows[1].cells[2].textContent).toBe('2')
        expect(tableRows[2].cells[2].textContent).toBe('0')

        expect(tableRows[1].cells[3].textContent).toBe('1')
        expect(tableRows[2].cells[3].textContent).toBe('0')

        expect(tableRows[1].cells[4].textContent).toBe('Team Two')
        expect(tableRows[2].cells[4].textContent).toBe('Team Four')

        expect(tableRows[1].cells[5].textContent).toBe('finished')
        expect(tableRows[2].cells[5].textContent).toBe('scheduled')

        expect(tableRows[1].cells[6].textContent).toContain('edit')
        expect(tableRows[1].cells[6].textContent).toContain('delete')

        expect(tableRows[2].cells[6].textContent).toContain('edit')
        expect(tableRows[2].cells[6].textContent).toContain('delete')
    })

    it('should display "No games found" row when dataSource is empty', () => {
        component.dataSource.data = []
        const noDataRow: HTMLElement = fixture.nativeElement.querySelector('tr.mat-row td.mat-cell')

        expect(noDataRow).toBeTruthy()
        expect(noDataRow.textContent).toContain('No games found')
    })

    it('should show error when season fails to load', () => {
        const mockError = new Error('Failed')
        seasonsService.getSeason.and.returnValue(throwError(() => mockError))

        component.ngOnInit()

        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Load Error',
            'Failed to load season',
            mockError
        )
    })

    it('should show error when games fail to load', () => {
        const mockError = new Error('Failed')
        gamesService.getGames.and.returnValue(throwError(() => mockError))

        component.ngOnInit()

        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Load Error',
            'Failed to load games',
            mockError
        )
    })

    it('should navigate to create game page when create button is clicked', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')

        const button = fixture.debugElement.query(By.css('.actions button'))
        button.nativeElement.click()

        const call = routerSpy.calls.all()[0].args[0].toString()
        expect(call).toEqual('/admin/competitions/comp1/seasons/season1/games/create')
    })

    it('should navigate to the correct edit page when edit button is clicked', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')

        const editButton = fixture.debugElement.query(By.css('[data-testid="edit-button"]'))
        editButton.nativeElement.click()

        const call = routerSpy.calls.all()[0].args[0].toString()
        expect(call).toEqual('/admin/competitions/comp1/seasons/season1/games/game1')
    })

    it('should call deleteGame when confirmed', () => {
        notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
        component.confirmDelete(mockGames[0])
        expect(gamesService.deleteGame).toHaveBeenCalledWith('comp1', 'season1', 'game1')
        expect(notificationService.showSnackbar).toHaveBeenCalledWith('Game deleted successfully')
    })

    it('should not call deleteGame when cancelled', () => {
        notificationService.showConfirm.and.returnValue({ afterClosed: () => of(false) } as any)
        component.confirmDelete(mockGames[0])
        expect(gamesService.deleteGame).not.toHaveBeenCalled()
        expect(notificationService.showSnackbar).not.toHaveBeenCalled()
    })

    it('should show error if deleteGame fails', () => {
        const mockError = new Error('Failed')
        gamesService.deleteGame.and.returnValue(throwError(() => mockError))

        notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
        component.confirmDelete(mockGames[0])

        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Delete Error',
            'Failed to delete game',
            mockError
        )
    })
})
