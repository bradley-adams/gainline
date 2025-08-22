import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Team } from '../../types/api'
import { environment } from '../../../environments/environment'
import { Observable } from 'rxjs'

@Injectable({
    providedIn: 'root'
})
export class TeamsService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getTeams(): Observable<Team[]> {
        return this.http.get<Team[]>(`${this.path}/v1/teams`)
    }

    getTeam(id: string): Observable<Team> {
        return this.http.get<Team>(`${this.path}/v1/teams/${id}`)
    }

    createTeam(team: Partial<Team>): Observable<Team> {
        return this.http.post<Team>(`${this.path}/v1/teams`, team)
    }

    updateTeam(id: string, team: Partial<Team>): Observable<Team> {
        return this.http.put<Team>(`${this.path}/v1/teams/${id}`, team)
    }

    deleteTeam(id: string): Observable<void> {
        return this.http.delete<void>(`${this.path}/v1/teams/${id}`)
    }
}
