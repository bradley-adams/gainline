import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Game } from '../../types/api'
import { environment } from '../../../environments/environment'
import { Observable } from 'rxjs'

@Injectable({
    providedIn: 'root'
})
export class GamesService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getGames(competitionId: string, seasonId: string): Observable<Game[]> {
        return this.http.get<Game[]>(
            `${this.path}/v1/competitions/${competitionId}/seasons/${seasonId}/games`
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
