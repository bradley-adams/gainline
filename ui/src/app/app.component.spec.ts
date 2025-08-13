import { TestBed } from '@angular/core/testing'
import { AppComponent } from './app.component'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter } from '@angular/router'
import { CompetitionDetailComponent } from './pages/competition-detail/competition-detail.component'

describe('AppComponent', () => {
    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [AppComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    { path: 'competitions/create', component: CompetitionDetailComponent }
                ])
            ]
        }).compileComponents()
    })

    it('should create the app', () => {
        const fixture = TestBed.createComponent(AppComponent)
        const app = fixture.componentInstance
        expect(app).toBeTruthy()
    })

    it('should render the header component', () => {
        const fixture = TestBed.createComponent(AppComponent)
        fixture.detectChanges()
        const headerEl = fixture.nativeElement.querySelector('app-header')
        expect(headerEl).toBeTruthy()
    })

    it('should include a router outlet', () => {
        const fixture = TestBed.createComponent(AppComponent)
        fixture.detectChanges()
        const outlet = fixture.nativeElement.querySelector('router-outlet')
        expect(outlet).toBeTruthy()
    })
})
