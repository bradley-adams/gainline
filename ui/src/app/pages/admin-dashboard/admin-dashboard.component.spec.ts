import { ComponentFixture, TestBed } from '@angular/core/testing'

import { AdminDashboardComponent } from './admin-dashboard.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { TeamListComponent } from '../team-list/team-list.component'
import { CompetitionListComponent } from '../competition-list/competition-list.component'
import { By } from '@angular/platform-browser'

describe('AdminDashboardComponent', () => {
    let component: AdminDashboardComponent
    let fixture: ComponentFixture<AdminDashboardComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [AdminDashboardComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/teams',
                        component: TeamListComponent
                    },
                    {
                        path: 'admin/competitions',
                        component: CompetitionListComponent
                    }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(AdminDashboardComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should have a link to manage competitions', () => {
        const button = fixture.debugElement.query(By.css('button[routerLink="/admin/competitions"]'))
        expect(button).toBeTruthy()
        expect(button.nativeElement.textContent).toContain('Manage')
    })

    it('should have a link to manage teams', () => {
        const button = fixture.debugElement.query(By.css('button[routerLink="/admin/teams"]'))
        expect(button).toBeTruthy()
        expect(button.nativeElement.textContent).toContain('Manage')
    })
})
