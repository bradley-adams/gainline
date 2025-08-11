import { ComponentFixture, TestBed } from '@angular/core/testing'
import { CompetitionComponent } from './competition.component'
import { provideRouter } from '@angular/router'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { provideHttpClient } from '@angular/common/http'

describe('CompetitionComponent', () => {
    let component: CompetitionComponent
    let fixture: ComponentFixture<CompetitionComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [CompetitionComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    { path: 'competitions/:competition-id', component: CompetitionComponent }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(CompetitionComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })
})
