import { ComponentFixture, TestBed } from '@angular/core/testing'
import { By } from '@angular/platform-browser'
import { MAT_DIALOG_DATA } from '@angular/material/dialog'
import { ErrorComponent } from './error.component'

describe('ErrorComponent', () => {
    let component: ErrorComponent
    let fixture: ComponentFixture<ErrorComponent>

    const mockData = {
        title: 'Test Error',
        message: 'An unexpected error occurred.'
    }

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [ErrorComponent],
            providers: [{ provide: MAT_DIALOG_DATA, useValue: mockData }]
        }).compileComponents()

        fixture = TestBed.createComponent(ErrorComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should display the given title', () => {
        const titleEl = fixture.debugElement.query(By.css('h1')).nativeElement
        expect(titleEl.textContent).toContain(mockData.title)
    })

    it('should display the given message', () => {
        const messageEl = fixture.debugElement.query(By.css('mat-dialog-content')).nativeElement
        expect(messageEl.textContent).toContain(mockData.message)
    })

    it('should have a cancel button with text "OK"', () => {
        const cancelBtn = fixture.debugElement.query(By.css('[data-testid="cancel-button"]')).nativeElement
        expect(cancelBtn.textContent).toContain('OK')
    })
})
