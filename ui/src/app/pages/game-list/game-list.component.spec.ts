import { ComponentFixture, TestBed } from '@angular/core/testing'

import { GameListComponent } from './game-list.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { GameDetailComponent } from '../game-detail/game-detail.component'

describe('GameListComponent', () => {
    let component: GameListComponent
    let fixture: ComponentFixture<GameListComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [GameListComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/:competition-id/seasons/:season-id/games/create',
                        component: GameDetailComponent
                    },
                    {
                        path: 'admin/competitions/:competition-id/seasons/:season-id/games/:game-id',
                        component: GameDetailComponent
                    }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(GameListComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })
})
