import { ComponentFixture, TestBed } from '@angular/core/testing'
import { By } from '@angular/platform-browser'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'

import { AdminDashboardComponent } from './admin-dashboard.component'
import { TeamListComponent } from '../team-list/team-list.component'
import { CompetitionListComponent } from '../competition-list/competition-list.component'

describe('AdminDashboardComponent', () => {
    let fixture: ComponentFixture<AdminDashboardComponent>
    let component: AdminDashboardComponent

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [AdminDashboardComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    { path: 'admin/teams', component: TeamListComponent },
                    { path: 'admin/competitions', component: CompetitionListComponent }
                ])
            ]
        }).compileComponents()
    })

    beforeEach(() => {
        fixture = TestBed.createComponent(AdminDashboardComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should render a button linking to the competitions page', () => {
        const button = fixture.debugElement.query(By.css('button[routerLink="/admin/competitions"]'))
        expect(button).toBeTruthy()
        expect(button.nativeElement.textContent).toContain('Manage')
    })

    it('should render a button linking to the teams page', () => {
        const button = fixture.debugElement.query(By.css('button[routerLink="/admin/teams"]'))
        expect(button).toBeTruthy()
        expect(button.nativeElement.textContent).toContain('Manage')
    })
})
