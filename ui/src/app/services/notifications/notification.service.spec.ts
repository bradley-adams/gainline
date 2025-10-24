import { TestBed } from '@angular/core/testing'
import { MatSnackBar } from '@angular/material/snack-bar'
import { MatDialog } from '@angular/material/dialog'

import { NotificationService } from './notifications.service'
import { ConfirmComponent } from '../../components/dialog/confirm/confirm.component'
import { ErrorComponent } from '../../components/dialog/error/error.component'
import { SnackbarComponent } from '../../components/snackbar/snackbar.component'

describe('NotificationService', () => {
    let service: NotificationService
    let snackBarSpy: jasmine.SpyObj<MatSnackBar>
    let dialogSpy: jasmine.SpyObj<MatDialog>

    beforeEach(() => {
        snackBarSpy = jasmine.createSpyObj('MatSnackBar', ['openFromComponent'])
        dialogSpy = jasmine.createSpyObj('MatDialog', ['open'])

        TestBed.configureTestingModule({
            providers: [
                NotificationService,
                { provide: MatSnackBar, useValue: snackBarSpy },
                { provide: MatDialog, useValue: dialogSpy }
            ]
        })

        service = TestBed.inject(NotificationService)
    })

    it('should be created', () => {
        expect(service).toBeTruthy()
    })

    it('should show snackbar with message and default duration', () => {
        service.showSnackbar('Test message')

        expect(snackBarSpy.openFromComponent).toHaveBeenCalledWith(SnackbarComponent, {
            data: 'Test message',
            duration: 5000
        })
    })

    it('should show snackbar with custom duration', () => {
        service.showSnackbar('Custom message', 10000)

        expect(snackBarSpy.openFromComponent).toHaveBeenCalledWith(SnackbarComponent, {
            data: 'Custom message',
            duration: 10000
        })
    })

    it('should open confirm dialog with provided title and message', () => {
        service.showConfirm('Confirm Title', 'Confirm Message')

        expect(dialogSpy.open).toHaveBeenCalledWith(ConfirmComponent, {
            data: { title: 'Confirm Title', message: 'Confirm Message' },
            disableClose: true
        })
    })

    it('should open error dialog and log sanitized error', () => {
        const consoleErrorSpy = spyOn(console, 'error')
        const mockError = new Error('Test error')

        service.showErrorAndLog('Error Title', 'Something went wrong!!!', mockError)

        expect(consoleErrorSpy).toHaveBeenCalledWith('something went wrong', mockError)
        expect(dialogSpy.open).toHaveBeenCalledWith(ErrorComponent, {
            data: { title: 'Error Title', message: 'Something went wrong!!!' },
            disableClose: true
        })
    })

    it('should open warning dialog and log sanitized warning', () => {
        const consoleWarnSpy = spyOn(console, 'warn')

        service.showWarnAndLog('Warning Title', 'Be careful! Something looks off...')

        expect(consoleWarnSpy).toHaveBeenCalledWith('be careful something looks off')
        expect(dialogSpy.open).toHaveBeenCalledWith(ErrorComponent, {
            data: { title: 'Warning Title', message: 'Be careful! Something looks off...' },
            disableClose: true
        })
    })
})
