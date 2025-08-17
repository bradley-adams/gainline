import { ComponentFixture, TestBed } from '@angular/core/testing'

import { SeasonListComponent } from './season-list.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { SeasonDetailComponent } from '../season-detail/season-detail.component'

describe('SeasonListComponent', () => {
    let component: SeasonListComponent
    let fixture: ComponentFixture<SeasonListComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [SeasonListComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/:competition-id/seasons/create',
                        component: SeasonDetailComponent
                    },
                    {
                        path: 'admin/competitions/:competition-id/seasons/:season-id',
                        component: SeasonDetailComponent
                    }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(SeasonListComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })
})
