import { ComponentFixture, TestBed } from '@angular/core/testing'
import { By } from '@angular/platform-browser'
import { MAT_DIALOG_DATA } from '@angular/material/dialog'
import { ConfirmComponent } from './confirm.component'

describe('ConfirmComponent', () => {
    let component: ConfirmComponent
    let fixture: ComponentFixture<ConfirmComponent>

    const mockData = {
        title: 'Test Confirm',
        message: 'Are you sure you want to continue?'
    }

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [ConfirmComponent],
            providers: [{ provide: MAT_DIALOG_DATA, useValue: mockData }]
        }).compileComponents()

        fixture = TestBed.createComponent(ConfirmComponent)
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

    it('should have a cancel button with text "CANCEL"', () => {
        const cancelBtn = fixture.debugElement.query(By.css('[data-testid="cancel-button"]')).nativeElement
        expect(cancelBtn.textContent).toContain('CANCEL')
    })

    it('should have a confirm button with text "Accept"', () => {
        const confirmBtn = fixture.debugElement.query(By.css('[data-testid="confirm-button"]')).nativeElement
        expect(confirmBtn.textContent).toContain('Accept')
    })
})
