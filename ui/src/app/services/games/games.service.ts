import { HttpClient } from '@angular/common/http'
import { inject, Injectable } from '@angular/core'
import { Observable } from 'rxjs'
import { environment } from '../../../environments/environment'
import { Game, PaginatedResponse } from '../../types/api'

@Injectable({
    providedIn: 'root'
})
export class GamesService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getGames(
        competitionId: string,
        seasonId: string,
        page = 1,
        pageSize = 10
    ): Observable<PaginatedResponse<Game>> {
        return this.http.get<PaginatedResponse<Game>>(
            `${this.path}/v1/competitions/${competitionId}/seasons/${seasonId}/games`,
            {
                params: {
                    page,
                    page_size: pageSize
                }
            }
        )
    }

    getGame(competitionId: string, seasonId: string, id: string): Observable<Game> {
        return this.http.get<Game>(
            `${this.path}/v1/competitions/${competitionId}/seasons/${seasonId}/games/${id}`
        )
    }

    createGame(competitionId: string, seasonId: string, game: Partial<Game>): Observable<Game> {
        return this.http.post<Game>(
            `${this.path}/v1/competitions/${competitionId}/seasons/${seasonId}/games`,
            game
        )
    }

    updateGame(competitionId: string, seasonId: string, id: string, game: Partial<Game>): Observable<Game> {
        return this.http.put<Game>(
            `${this.path}/v1/competitions/${competitionId}/seasons/${seasonId}/games/${id}`,
            game
        )
    }

    deleteGame(competitionId: string, seasonId: string, id: string): Observable<void> {
        return this.http.delete<void>(
            `${this.path}/v1/competitions/${competitionId}/seasons/${seasonId}/games/${id}`
        )
    }
}
