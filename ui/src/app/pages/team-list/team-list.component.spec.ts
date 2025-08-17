import { ComponentFixture, TestBed } from '@angular/core/testing'

import { TeamListComponent } from './team-list.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { TeamDetailComponent } from '../team-detail/team-detail.component'

describe('TeamListComponent', () => {
    let component: TeamListComponent
    let fixture: ComponentFixture<TeamListComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [TeamListComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/teams/create',
                        component: TeamDetailComponent
                    },
                    {
                        path: 'admin/teams/:team-id',
                        component: TeamDetailComponent
                    }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(TeamListComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })
})
