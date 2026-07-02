import { TestBed } from '@angular/core/testing'
import { environment } from '../../../environments/environment'
import { GameState } from '../../types/api'
import { LiveGameService } from './live-game.service'

const mockCompetitionID = 'comp1'
const mockSeasonID = 'season1'
const mockGameID = 'game1'

const mockGameState: GameState = {
    game_id: mockGameID,
    home_score: 2,
    away_score: 1,
    status: 'playing',
    minute: 35
}

class MockEventSource {
    static instance: MockEventSource
    url: string
    listeners: Record<string, ((event: MessageEvent) => void)[]> = {}
    onerror: (() => void) | null = null

    constructor(url: string) {
        this.url = url
        MockEventSource.instance = this
    }

    addEventListener(event: string, handler: (event: MessageEvent) => void) {
        if (!this.listeners[event]) {
            this.listeners[event] = []
        }
        this.listeners[event].push(handler)
    }

    emit(event: string, data: unknown) {
        const msg = { data: JSON.stringify(data) } as MessageEvent
        this.listeners[event]?.forEach((fn) => fn(msg))
    }

    close() {}
}

describe('LiveGameService', () => {
    let service: LiveGameService
    const baseUrl = environment.apiUrl

    beforeEach(() => {
        ;(window as unknown as Record<string, unknown>)['EventSource'] = MockEventSource

        TestBed.configureTestingModule({})
        service = TestBed.inject(LiveGameService)
    })

    it('should be created', () => {
        expect(service).toBeTruthy()
    })

    it('should construct the correct SSE url', () => {
        service.watchGame(mockCompetitionID, mockSeasonID, mockGameID).subscribe()

        expect(MockEventSource.instance.url).toBe(
            `${baseUrl}/v1/competitions/${mockCompetitionID}/seasons/${mockSeasonID}/games/${mockGameID}/live`
        )
    })

    it('should emit game state updates', (done) => {
        service.watchGame(mockCompetitionID, mockSeasonID, mockGameID).subscribe((state) => {
            expect(state).toEqual(mockGameState)
            done()
        })

        MockEventSource.instance.emit('update', mockGameState)
    })

    it('should error the observable on SSE error', (done) => {
        service.watchGame(mockCompetitionID, mockSeasonID, mockGameID).subscribe({
            error: (err) => {
                expect(err).toBe('SSE connection error')
                done()
            }
        })

        MockEventSource.instance.onerror?.()
    })
})
