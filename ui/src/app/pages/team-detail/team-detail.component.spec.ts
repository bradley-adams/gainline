import { ComponentFixture, TestBed } from '@angular/core/testing'

import { TeamDetailComponent } from './team-detail.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { TeamListComponent } from '../team-list/team-list.component'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'

describe('TeamDetailComponent', () => {
    let component: TeamDetailComponent
    let fixture: ComponentFixture<TeamDetailComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [TeamDetailComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/teams',
                        component: TeamListComponent
                    }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(TeamDetailComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })
})
