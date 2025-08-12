import { ComponentFixture, TestBed } from '@angular/core/testing'

import { HeaderComponent } from './header.component'
import { By } from '@angular/platform-browser'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { CompetitionComponent } from '../../pages/competition/competition.component'

describe('HeaderComponent', () => {
    let component: HeaderComponent
    let fixture: ComponentFixture<HeaderComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [HeaderComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([{ path: 'competitions/create', component: CompetitionComponent }])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(HeaderComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should display the title "Gainline"', () => {
        const toolbarText = fixture.debugElement.nativeElement.textContent
        expect(toolbarText).toContain('Gainline')
    })

    it('should have an icon linking to /competitions/create', () => {
        const iconDebugEl = fixture.debugElement.query(By.css('mat-icon'))
        expect(iconDebugEl).toBeTruthy()

        const routerLink = iconDebugEl.attributes['ng-reflect-router-link']
        expect(routerLink).toBe('/competitions/create')
    })
})
