import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http'
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing'
import { TestBed } from '@angular/core/testing'
import { environment } from '../../../environments/environment'
import { Season, Stage, StageType, Team } from '../../types/api'
import { SeasonsService } from './seasons.service'

describe('SeasonsService', () => {
    let service: SeasonsService
    let httpMock: HttpTestingController
    const baseUrl = environment.apiUrl

    const mockCompetitionID = 'comp1'
    const mockSeasonID = 'season1'

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

    const MockStages: Stage[] = [
        {
            id: 'stage-round-1',
            name: 'Round 1',
            stageType: StageType.StageTypeRegular,
            orderIndex: 1,
            created_at: new Date('2025-01-01T00:00:00Z'),
            updated_at: new Date('2025-01-01T00:00:00Z')
        },
        {
            id: 'stage-round-2',
            name: 'Round 2',
            stageType: StageType.StageTypeRegular,
            orderIndex: 2,
            created_at: new Date('2025-01-01T00:00:00Z'),
            updated_at: new Date('2025-01-01T00:00:00Z')
        },
        {
            id: 'stage-round-3',
            name: 'Final',
            stageType: StageType.StageTypeFinals,
            orderIndex: 3,
            created_at: new Date('2025-01-01T00:00:00Z'),
            updated_at: new Date('2025-01-01T00:00:00Z')
        }
    ]

    const mockSeason: Season = {
        id: 'season1',
        competition_id: 'comp1',
        start_date: new Date('2025-01-01T00:00:00Z'),
        end_date: new Date('2025-12-31T23:59:59Z'),
        stages: MockStages,
        teams: mockTeams,
        created_at: new Date('2024-12-01T00:00:00Z'),
        updated_at: new Date('2024-12-01T00:00:00Z')
    }

    const mockSeasons: Season[] = [
        mockSeason,
        {
            id: 'season2',
            competition_id: 'comp1',
            start_date: new Date('2024-01-01T00:00:00Z'),
            end_date: new Date('2024-12-31T23:59:59Z'),
            stages: MockStages,
            teams: mockTeams,
            created_at: new Date('2023-12-01T00:00:00Z'),
            updated_at: new Date('2023-12-01T00:00:00Z')
        }
    ]

    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [provideHttpClient(withInterceptorsFromDi()), provideHttpClientTesting()]
        })
        service = TestBed.inject(SeasonsService)
        httpMock = TestBed.inject(HttpTestingController)
    })

    afterEach(() => {
        httpMock.verify()
    })

    it('should be created', () => {
        expect(service).toBeTruthy()
    })

    it('should get all seasons for a competition', () => {
        service.getSeasons(mockCompetitionID).subscribe((seasons) => {
            expect(seasons).toEqual(mockSeasons)
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/competitions/${mockCompetitionID}/seasons`)
        expect(req.request.method).toBe('GET')
        req.flush(mockSeasons)
    })

    it('should get a season by id', () => {
        service.getSeason(mockCompetitionID, mockSeasonID).subscribe((season) => {
            expect(season).toEqual(mockSeason)
        })

        const req = httpMock.expectOne(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}`
        )
        expect(req.request.method).toBe('GET')
        req.flush(mockSeason)
    })

    it('should create a season', () => {
        const newSeason: Partial<Season> = {
            competition_id: 'comp1',
            start_date: new Date('2024-01-01T00:00:00Z'),
            end_date: new Date('2024-12-31T23:59:59Z'),
            teams: ['team1', 'team2', 'team3', 'team4']
        }

        service.createSeason(mockCompetitionID, newSeason).subscribe((season) => {
            expect(season).toEqual(mockSeason)
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/competitions/${mockCompetitionID}/seasons`)
        expect(req.request.method).toBe('POST')
        expect(req.request.body).toEqual(newSeason)
        req.flush(mockSeason)
    })

    it('should update a season', () => {
        const update: Partial<Season> = {
            competition_id: 'comp1',
            start_date: new Date('2024-01-01T00:00:00Z'),
            end_date: new Date('2024-12-31T23:59:59Z'),
            teams: ['team1', 'team2', 'team3', 'team4']
        }

        service.updateSeason(mockCompetitionID, mockSeasonID, update).subscribe((season) => {
            expect(season).toEqual({ ...mockSeason, ...update })
        })

        const req = httpMock.expectOne(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}`
        )
        expect(req.request.method).toBe('PUT')
        expect(req.request.body).toEqual(update)
        req.flush({ ...mockSeason, ...update })
    })

    it('should delete a season', () => {
        service.deleteSeason(mockCompetitionID, mockSeasonID).subscribe((res) => {
            expect(res).toBeNull()
        })

        const req = httpMock.expectOne(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}`
        )
        expect(req.request.method).toBe('DELETE')
        req.flush(null)
    })
})
