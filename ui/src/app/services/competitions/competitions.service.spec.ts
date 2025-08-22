import { TestBed } from '@angular/core/testing'
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing'
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http'
import { CompetitionsService } from './competitions.service'
import { environment } from '../../../environments/environment'
import { Competition } from '../../types/api'

describe('CompetitionsService', () => {
    let service: CompetitionsService
    let httpMock: HttpTestingController
    const baseUrl = environment.apiUrl

    const mockCompetitionID = 'comp1'
    const mockCompetition: Competition = {
        id: 'comp1',
        name: 'Competition 1',
        created_at: new Date('2023-01-01T00:00:00Z'),
        updated_at: new Date('2023-01-02T00:00:00Z')
    }

    const mockCompetitions: Competition[] = [
        mockCompetition,
        {
            id: 'comp2',
            name: 'Competition 2',
            created_at: new Date('2023-01-03T00:00:00Z'),
            updated_at: new Date('2023-01-04T00:00:00Z')
        }
    ]

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [],
            providers: [provideHttpClient(withInterceptorsFromDi()), provideHttpClientTesting()]
        })
        service = TestBed.inject(CompetitionsService)
        httpMock = TestBed.inject(HttpTestingController)
    })

    afterEach(() => {
        httpMock.verify()
    })

    it('should be created', () => {
        expect(service).toBeTruthy()
    })

    it('should get all competitions', () => {
        service.getCompetitions().subscribe((competitions) => {
            expect(competitions).toEqual(mockCompetitions)
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/competitions`)
        expect(req.request.method).toBe('GET')
        req.flush(mockCompetitions)
    })

    it('should get a competition by id', () => {
        service.getCompetition('comp1').subscribe((competition) => {
            expect(competition).toEqual(mockCompetition)
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/competitions/comp1`)
        expect(req.request.method).toBe('GET')
        req.flush(mockCompetition)
    })

    it('should create a competition', () => {
        const newComp: Partial<Competition> = { name: 'Competition 1' }

        service.createCompetition(newComp).subscribe((competition) => {
            expect(competition).toEqual(mockCompetition)
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/competitions`)
        expect(req.request.method).toBe('POST')
        expect(req.request.body).toEqual(newComp)
        req.flush(mockCompetition)
    })

    it('should update a competition', () => {
        const update: Partial<Competition> = { name: 'Updated Name' }

        service.updateCompetition('comp1', update).subscribe((competition) => {
            expect(competition).toEqual({ ...mockCompetition, ...update })
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/competitions/comp1`)
        expect(req.request.method).toBe('PUT')
        expect(req.request.body).toEqual(update)
        req.flush({ ...mockCompetition, ...update })
    })

    it('should delete a competition', () => {
        service.deleteCompetition('comp1').subscribe((res) => {
            expect(res).toBeNull()
        })

        const req = httpMock.expectOne(`${baseUrl}/v1/competitions/comp1`)
        expect(req.request.method).toBe('DELETE')
        req.flush(null)
    })
})
