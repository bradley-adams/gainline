import { ComponentFixture, TestBed } from '@angular/core/testing'
import { provideRouter } from '@angular/router'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { NoopAnimationsModule } from '@angular/platform-browser/animations'
import { provideHttpClient } from '@angular/common/http'
import { CompetitionDetailComponent } from './competition-detail.component'
import { CompetitionListComponent } from '../competition-list/competition-list.component'

describe('CompetitionDetailComponent', () => {
    let component: CompetitionDetailComponent
    let fixture: ComponentFixture<CompetitionDetailComponent>

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [CompetitionDetailComponent, NoopAnimationsModule],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions',
                        component: CompetitionListComponent
                    }
                ])
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(CompetitionDetailComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })
})
