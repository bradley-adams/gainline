import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Season } from '../../types/api'
import { environment } from '../../../environments/environment'

@Injectable({
    providedIn: 'root'
})
export class SeasonsService {
        private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getSeasons(competitionId: string) {
        return this.http.get<Season[]>(`${this.path}/v1/competitions/${competitionId}/seasons`,)
    }
}
