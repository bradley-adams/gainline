import { inject, Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { Competition } from '../../types/api'
import { environment } from '../../../environments/environment'

@Injectable({
    providedIn: 'root'
})
export class CompetitionsService {
        private readonly path = environment.apiUrl
    private readonly http = inject(HttpClient)

    getCompetitions() {
        return this.http.get<Competition[]>(`${this.path}/v1/competitions`,)
    }
}
