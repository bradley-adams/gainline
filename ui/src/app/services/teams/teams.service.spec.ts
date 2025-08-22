import { TestBed } from '@angular/core/testing'
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing'
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http'
import { TeamsService } from './teams.service'
import { environment } from '../../../environments/environment'
import { Team } from '../../types/api'

describe('TeamsService', () => {
    let service: TeamsService
    let httpMock: HttpTestingController
    const baseUrl = environment.apiUrl

    const mockTeamID = 'team1'

    const mockTeam: Team = {
        id: mockTeamID,
        abbreviation: 'T1',
        location: 'City A',
        name: 'Team One',
        created_at: new Date('2024-01-01T00:00:00Z'),
        updated_at: new Date('2024-01-01T00:00:00Z')
    }

    const mockTeams: Team[] = [
        mockTeam,
        {
            id: 'team2',
            abbreviation: 'T2',
            location: 'City B',
            name: 'Team Two',
            created_at: new Date('2024-01-02T00:00:00Z'),
            updated_at: new Date('2024-01-02T00:00:00Z')
        }
    ]

    beforeEach(() => {
        TestBed.configureTestingModule({
            providers: [provideHttpClient(withInterceptorsFromDi()), provideHttpClientTesting()]
        })
        service = TestBed.inject(TeamsService)
        httpMock = TestBed.inject(HttpTestingController)
    })

    afterEach(() => {
        httpMock.verify()
    })

    it('should be created', () => {
        expect(service).toBeTruthy()
    })

    it('should get all teams', () => {
        service.getTeams().subscribe((teams) => {
            expect(teams).toEqual(mockTeams)
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/teams`)
        expect(req.request.method).toBe('GET')
        req.flush(mockTeams)
    })

    it('should get a team by id', () => {
        service.getTeam(mockTeamID).subscribe((team) => {
            expect(team).toEqual(mockTeam)
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/teams/${mockTeamID}`)
        expect(req.request.method).toBe('GET')
        req.flush(mockTeam)
    })

    it('should create a team', () => {
        const newTeam: Partial<Team> = {
            abbreviation: 'T3',
            location: 'City C',
            name: 'Team Three'
        }

        service.createTeam(newTeam).subscribe((team) => {
            expect(team).toEqual(mockTeam)
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/teams`)
        expect(req.request.method).toBe('POST')
        expect(req.request.body).toEqual(newTeam)
        req.flush(mockTeam)
    })

    it('should update a team', () => {
        const update: Partial<Team> = { location: 'Updated City' }

        service.updateTeam(mockTeamID, update).subscribe((team) => {
            expect(team).toEqual({ ...mockTeam, ...update })
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/teams/${mockTeamID}`)
        expect(req.request.method).toBe('PUT')
        expect(req.request.body).toEqual(update)
        req.flush({ ...mockTeam, ...update })
    })

    it('should delete a team', () => {
        service.deleteTeam(mockTeamID).subscribe((res) => {
            expect(res).toBeNull()
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/teams/${mockTeamID}`)
        expect(req.request.method).toBe('DELETE')
        req.flush(null)
    })
})
