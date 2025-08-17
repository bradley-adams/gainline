import { ComponentFixture, TestBed } from '@angular/core/testing'

import { SeasonDetailComponent } from './season-detail.component'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { SeasonListComponent } from '../season-list/season-list.component'

describe('SeasonDetailComponent', () => {
    let component: SeasonDetailComponent
    let fixture: ComponentFixture<SeasonDetailComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [SeasonDetailComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/:competition-id/seasons',
                        component: SeasonListComponent
                    }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(SeasonDetailComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })
})
