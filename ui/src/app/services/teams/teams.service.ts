import { HttpClient } from '@angular/common/http'
import { inject, Injectable } from '@angular/core'
import { Observable } from 'rxjs'
import { environment } from '../../../environments/environment'
import { PaginatedResponse, Team } from '../../types/api'

@Injectable({
    providedIn: 'root'
})
export class TeamsService {
    private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getTeams(): Observable<Team[]> {
        return this.http.get<Team[]>(`${this.path}/v1/teams`)
    }

    getTeamsPaginated(page: number, pageSize: number): Observable<PaginatedResponse<Team>> {
        return this.http.get<PaginatedResponse<Team>>(`${this.path}/v1/teamspaginated`, {
            params: {
                page,
                page_size: pageSize
            }
        })
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
