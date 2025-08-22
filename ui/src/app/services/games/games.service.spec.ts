import { TestBed } from '@angular/core/testing'
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing'
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http'
import { GamesService } from './games.service'
import { environment } from '../../../environments/environment'
import { Game, GameStatus } from '../../types/api'

describe('GamesService', () => {
    let service: GamesService
    let httpMock: HttpTestingController
    const baseUrl = environment.apiUrl

    const mockCompetitionID = 'comp1'
    const mockSeasonID = 'season1'
    const mockGameID = 'game1'

    const mockGame: Game = {
        id: mockGameID,
        season_id: mockSeasonID,
        round: 1,
        date: new Date('2025-05-01T15:00:00Z'),
        home_team_id: 'team1',
        away_team_id: 'team2',
        home_score: 2,
        away_score: 1,
        status: GameStatus.SCHEDULED,
        created_at: new Date('2025-04-01T00:00:00Z'),
        updated_at: new Date('2025-04-01T00:00:00Z')
    }

    const mockGames: Game[] = [
        mockGame,
        {
            id: 'game2',
            season_id: mockSeasonID,
            round: 2,
            date: new Date('2025-05-08T15:00:00Z'),
            home_team_id: 'team3',
            away_team_id: 'team4',
            home_score: 0,
            away_score: 0,
            status: GameStatus.SCHEDULED,
            created_at: new Date('2025-04-02T00:00:00Z'),
            updated_at: new Date('2025-04-02T00:00:00Z')
        }
    ]

    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [provideHttpClient(withInterceptorsFromDi()), provideHttpClientTesting()]
        })
        service = TestBed.inject(GamesService)
        httpMock = TestBed.inject(HttpTestingController)
    })

    afterEach(() => {
        httpMock.verify()
    })

    it('should be created', () => {
        expect(service).toBeTruthy()
    })

    it('should get all games for a season', () => {
        service.getGames(mockCompetitionID, mockSeasonID).subscribe((games) => {
            expect(games).toEqual(mockGames)
        })

        const req = httpMock.expectOne(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}/games`
        )
        expect(req.request.method).toBe('GET')
        req.flush(mockGames)
    })

    it('should get a game by id', () => {
        service.getGame(mockCompetitionID, mockSeasonID, mockGameID).subscribe((game) => {
            expect(game).toEqual(mockGame)
        })

        const req = httpMock.expectOne(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}/games/${mockGameID}`
        )
        expect(req.request.method).toBe('GET')
        req.flush(mockGame)
    })

    it('should create a game', () => {
        const newGame: Partial<Game> = {
            round: 1,
            date: new Date('2025-05-01T15:00:00Z'),
            home_team_id: 'team1',
            away_team_id: 'team2',
            home_score: 2,
            away_score: 1,
            status: GameStatus.SCHEDULED
        }

        service.createGame(mockCompetitionID, mockSeasonID, newGame).subscribe((game) => {
            expect(game).toEqual(mockGame)
        })

        const req = httpMock.expectOne(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}/games`
        )
        expect(req.request.method).toBe('POST')
        expect(req.request.body).toEqual(newGame)
        req.flush(mockGame)
    })

    it('should update a game', () => {
        const update: Partial<Game> = {
            home_score: 3,
            away_score: 2,
            status: GameStatus.SCHEDULED
        }

        service.updateGame(mockCompetitionID, mockSeasonID, mockGameID, update).subscribe((game) => {
            expect(game).toEqual({ ...mockGame, ...update })
        })

        const req = httpMock.expectOne(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}/games/${mockGameID}`
        )
        expect(req.request.method).toBe('PUT')
        expect(req.request.body).toEqual(update)
        req.flush({ ...mockGame, ...update })
    })

    it('should delete a game', () => {
        service.deleteGame(mockCompetitionID, mockSeasonID, mockGameID).subscribe((res) => {
            expect(res).toBeNull()
        })

        const req = httpMock.expectOne(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}/games/${mockGameID}`
        )
        expect(req.request.method).toBe('DELETE')
        req.flush(null)
    })
})
