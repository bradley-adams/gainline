import { HttpClient } from '@angular/common/http'
import { inject, Injectable } from '@angular/core'
import { Observable } from 'rxjs'
import { environment } from '../../../environments/environment'
import { Competition, PaginatedResponse } from '../../types/api'

@Injectable({
    providedIn: 'root'
})
export class CompetitionsService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getCompetitions(): Observable<Competition[]> {
        return this.http.get<Competition[]>(`${this.path}/v1/competitions`)
    }

    getPaginatedCompetitions(page: number, pageSize: number): Observable<PaginatedResponse<Competition>> {
        return this.http.get<PaginatedResponse<Competition>>(`${this.path}/v1/competitions2`, {
            params: {
                page: page.toString(),
                page_size: pageSize.toString()
            }
        })
    }

    getCompetition(id: string): Observable<Competition> {
        return this.http.get<Competition>(`${this.path}/v1/competitions/${id}`)
    }

    createCompetition(competition: Partial<Competition>): Observable<Competition> {
        return this.http.post<Competition>(`${this.path}/v1/competitions`, competition)
    }

    updateCompetition(id: string, competition: Partial<Competition>): Observable<Competition> {
        return this.http.put<Competition>(`${this.path}/v1/competitions/${id}`, competition)
    }

    deleteCompetition(id: string) {
        return this.http.delete<void>(`${this.path}/v1/competitions/${id}`)
    }
}
