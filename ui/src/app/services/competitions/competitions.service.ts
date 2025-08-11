import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Competition } from '../../types/api'
import { environment } from '../../../environments/environment'
import { Observable } from 'rxjs'

@Injectable({
    providedIn: 'root'
})
export class CompetitionsService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getCompetitions(): Observable<Competition[]> {
        return this.http.get<Competition[]>(`${this.path}/v1/competitions`)
    }

    getCompetition(id: string): Observable<Competition> {
        return this.http.get<Competition>(`${this.path}/v1/competitions/${id}`)
    }

    createCompetition(competition: Partial<Competition>): Observable<Competition> {
        console.log('Posting competition:', competition)
        return this.http.post<Competition>(`${this.path}/v1/competitions`, competition)
    }

    updateCompetition(id: string, competition: Partial<Competition>): Observable<Competition> {
        return this.http.put<Competition>(`${this.path}/v1/competitions/${id}`, competition)
    }

    deleteCompetition(id: string): Observable<void> {
        return this.http.delete<void>(`${this.path}/v1/competitions/${id}`)
    }
}
