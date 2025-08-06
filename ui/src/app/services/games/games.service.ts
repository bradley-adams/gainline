import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Game } from '../../types/api'
import { environment } from '../../../environments/environment'

@Injectable({
    providedIn: 'root'
})
export class GamesService {
        private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getGames(competitionId: string, seasonId: string) {
        return this.http.get<Game[]>(`${this.path}/v1/competitions/${competitionId}/seasons/${seasonId}/games `,)
    }
}

