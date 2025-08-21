import { ComponentFixture, TestBed } from '@angular/core/testing'

import { GameListComponent } from './game-list.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ActivatedRoute, convertToParamMap, provideRouter, Router } from '@angular/router'
import { GameDetailComponent } from '../game-detail/game-detail.component'
import { Game, GameStatus } from '../../types/api'
import { GamesService } from '../../services/games/games.service'
import { of, throwError } from 'rxjs'
import { By } from '@angular/platform-browser'

describe('GameListComponent', () => {
    let component: GameListComponent
    let fixture: ComponentFixture<GameListComponent>
    let router: Router

    let gamesService: jasmine.SpyObj<GamesService>

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
        gamesService = jasmine.createSpyObj('GamesService', ['getGames'])
        gamesService.getGames.and.returnValue(of(mockGames))

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
                {
                    provide: ActivatedRoute,
                    useValue: {
                        snapshot: {
                            paramMap: convertToParamMap({ 'competition-id': 'comp1', 'season-id': 'season1' })
                        }
                    }
                }
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

    it('should load games into the table', () => {
        const rows = fixture.nativeElement.querySelectorAll('tr')
        expect(rows.length).toBe(3)

        // first game row
        expect(rows[1].textContent).toContain('02/02/2025')
        expect(rows[1].textContent).toContain('team1')
        expect(rows[1].textContent).toContain('team2')
        expect(rows[1].textContent).toContain('finished')

        // second game row
        expect(rows[2].textContent).toContain('02/03/2025')
        expect(rows[2].textContent).toContain('team3')
        expect(rows[2].textContent).toContain('team4')
        expect(rows[2].textContent).toContain('scheduled')
    })

    it('should show error when games fail to load', () => {
        const error = new Error('Failed to load')
        spyOn(console, 'error')

        gamesService.getGames.and.returnValue(throwError(() => error))

        component.ngOnInit()
        expect(console.error).toHaveBeenCalledWith('Error loading games:', error)
    })

    it('should navigate when "Create Game" button is clicked', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')
        fixture.detectChanges()

        const button = fixture.debugElement.query(By.css('.actions button'))

        button.nativeElement.click()
        const call = routerSpy.calls.all()[0].args[0].toString()
        expect(call).toEqual('/admin/competitions/comp1/seasons/season1/games/create')
    })

    it('should display table correctly', () => {
        const tableRows = fixture.nativeElement.querySelectorAll('tr')
        expect(tableRows.length).toBe(mockGames.length + 1)

        // Header row
        const headerRow = tableRows[0]
        expect(headerRow.cells[0].innerHTML).toBe('Date')
        expect(headerRow.cells[1].innerHTML).toBe('Home Team')
        expect(headerRow.cells[2].innerHTML).toBe('Home Score')
        expect(headerRow.cells[3].innerHTML).toBe('Away Score')
        expect(headerRow.cells[4].innerHTML).toBe('Away Team')
        expect(headerRow.cells[5].innerHTML).toBe('Status')
        expect(headerRow.cells[6].innerHTML).toBe('Actions')

        expect(tableRows[1].cells[0].textContent).toBe('02/02/2025 04:00')
        expect(tableRows[1].cells[1].textContent).toBe('team1')
        expect(tableRows[1].cells[2].textContent).toBe('2')
        expect(tableRows[1].cells[3].textContent).toBe('1')
        expect(tableRows[1].cells[4].textContent).toBe('team2')
        expect(tableRows[1].cells[5].textContent).toBe('finished')
        expect(tableRows[1].cells[6].textContent).toBe('edit')

        expect(tableRows[2].cells[0].textContent).toBe('02/03/2025 04:00')
        expect(tableRows[2].cells[1].textContent).toBe('team3')
        expect(tableRows[2].cells[2].textContent).toBe('0')
        expect(tableRows[2].cells[3].textContent).toBe('0')
        expect(tableRows[2].cells[4].textContent).toBe('team4')
        expect(tableRows[2].cells[5].textContent).toBe('scheduled')
        expect(tableRows[2].cells[6].textContent).toBe('edit')
    })

    it('edit buttons navigate correctly', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')

        const editButton = fixture.debugElement.query(By.css('[data-testid="edit-button"]'))

        editButton.nativeElement.click()
        const call = routerSpy.calls.all()[0].args[0].toString()
        expect(call).toEqual('/admin/competitions/comp1/seasons/season1/games/game1')
    })
})
