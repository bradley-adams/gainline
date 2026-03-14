import { HttpClient } from '@angular/common/http'
import { inject, Injectable } from '@angular/core'
import { Observable } from 'rxjs'
import { environment } from '../../../environments/environment'
import { PaginatedResponse, Season } from '../../types/api'

@Injectable({
    providedIn: 'root'
})
export class SeasonsService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getSeasons(competitionId: string): Observable<Season[]> {
        return this.http.get<Season[]>(`${this.path}/v1/competitions/${competitionId}/seasons`)
    }

    getPaginatedSeasons(
        competitionId: string,
        page = 1,
        pageSize = 10
    ): Observable<PaginatedResponse<Season>> {
        return this.http.get<PaginatedResponse<Season>>(
            `${this.path}/v1/competitions/${competitionId}/seasons`,
            {
                params: {
                    page,
                    page_size: pageSize
                }
            }
        )
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
