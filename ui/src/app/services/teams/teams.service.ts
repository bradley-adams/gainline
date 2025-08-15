import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Team } from '../../types/api'
import { environment } from '../../../environments/environment'

@Injectable({
    providedIn: 'root'
})
export class TeamsService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getTeams() {
        return this.http.get<Team[]>(`${this.path}/v1/teams `)
    }
}
