import { ComponentFixture, TestBed } from '@angular/core/testing'
import { SnackbarComponent } from './snackbar.component'
import { By } from '@angular/platform-browser'
import { MatSnackBarRef, MAT_SNACK_BAR_DATA } from '@angular/material/snack-bar'

describe('SnackbarComponent', () => {
    let component: SnackbarComponent
    let fixture: ComponentFixture<SnackbarComponent>
    let mockSnackBarRef: jasmine.SpyObj<MatSnackBarRef<SnackbarComponent>>

    beforeEach(async () => {
        mockSnackBarRef = jasmine.createSpyObj('MatSnackBarRef', ['dismiss'])

        await TestBed.configureTestingModule({
            imports: [SnackbarComponent],
            providers: [
                { provide: MatSnackBarRef, useValue: mockSnackBarRef },
                { provide: MAT_SNACK_BAR_DATA, useValue: 'Test message' }
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(SnackbarComponent)
        component = fixture.componentInstance
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should display the message passed via MAT_SNACK_BAR_DATA', () => {
        const text = fixture.debugElement.query(By.css('span')).nativeElement.textContent
        expect(text).toContain('Test message')
    })

    it('should display a close icon', () => {
        const icon = fixture.debugElement.query(By.css('mat-icon'))
        expect(icon).toBeTruthy()
        expect(icon.nativeElement.textContent.trim()).toBe('clear')
    })

    it('should dismiss the snackbar when close icon is clicked', () => {
        const button = fixture.debugElement.query(By.css('button'))
        button.triggerEventHandler('click')
        expect(mockSnackBarRef.dismiss).toHaveBeenCalled()
    })
})
