import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Season } from '../../types/api'
import { environment } from '../../../environments/environment'
import { Observable } from 'rxjs'

@Injectable({
    providedIn: 'root'
})
export class SeasonsService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getSeasons(competitionId: string): Observable<Season[]> {
        return this.http.get<Season[]>(`${this.path}/v1/competitions/${competitionId}/seasons`)
    }

    getSeason(competitionId: string, id: string): Observable<Season> {
        return this.http.get<Season>(`${this.path}/v1/competitions/${competitionId}/seasons/${id}`)
    }

    createSeason(competitionId: string, season: Partial<Season>): Observable<Season> {
        return this.http.post<Season>(`${this.path}/v1/competitions/${competitionId}/seasons`, season)
    }

    updateSeason(competitionId: string, id: string, season: Partial<Season>): Observable<Season> {
        return this.http.put<Season>(`${this.path}/v1/competitions/${competitionId}/seasons/${id}`, season)
    }

    deleteSeason(competitionId: string, id: string): Observable<void> {
        return this.http.delete<void>(`${this.path}/v1/competitions/${competitionId}/seasons/${id}`)
    }
}
