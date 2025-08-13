import { ComponentFixture, TestBed } from '@angular/core/testing'

import { CompetitionListComponent } from './competition-list.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { CompetitionDetailComponent } from '../competition-detail/competition-detail.component'

describe('CompetitionListComponent', () => {
    let component: CompetitionListComponent
    let fixture: ComponentFixture<CompetitionListComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [CompetitionListComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/create',
                        component: CompetitionDetailComponent
                    },
                    {
                        path: 'admin/competitions/:competition-id',
                        component: CompetitionDetailComponent
                    }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(CompetitionListComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })
})
