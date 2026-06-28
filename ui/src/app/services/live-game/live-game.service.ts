import { Injectable } from '@angular/core'
import { Observable } from 'rxjs'
import { environment } from '../../../environments/environment'
import { GameState } from '../../types/api'

@Injectable({
    providedIn: 'root'
})
export class LiveGameService {
    private readonly path = environment.apiUrl

    watchGame(competitionId: string, seasonId: string, gameId: string): Observable<GameState> {
        return new Observable((observer) => {
            const url = `${this.path}/v1/competitions/${competitionId}/seasons/${seasonId}/games/${gameId}/live`
            const eventSource = new EventSource(url)

            eventSource.addEventListener('update', (event: MessageEvent) => {
                const state: GameState = JSON.parse(event.data)
                observer.next(state)
            })

            eventSource.onerror = () => {
                observer.error('SSE connection error')
                eventSource.close()
            }

            return () => eventSource.close()
        })
    }
}
