import { Injectable, inject } from '@angular/core'
import { MatSnackBar } from '@angular/material/snack-bar'
import { MatDialog, MatDialogRef } from '@angular/material/dialog'
import { ConfirmComponent } from '../../components/dialog/confirm/confirm.component'
import { ErrorComponent } from '../../components/dialog/error/error.component'
import { SnackbarComponent } from '../../components/snackbar/snackbar.component'

@Injectable({
    providedIn: 'root'
})
export class NotificationService {
    private readonly dialog = inject(MatDialog)
    private readonly snackBar = inject(MatSnackBar)

    showSnackbar(message: string, duration = 5000): void {
        this.snackBar.openFromComponent(SnackbarComponent, {
            data: message,
            duration
        })
    }

    showConfirm(title: string, message: string) {
        return this.dialog.open(ConfirmComponent, {
            data: { title, message },
            disableClose: true
        })
    }

    showErrorAndLog(title: string, message: string, error: Error) {
        const sanitizedMessage = message.toLowerCase().replace(/[^\w\s]/g, '')
        console.error(sanitizedMessage, error)

        return this.dialog.open(ErrorComponent, {
            data: { title, message },
            disableClose: true
        })
    }

    showWarnAndLog(title: string, message: string) {
        const sanitizedMessage = message.toLowerCase().replace(/[^\w\s]/g, '')
        console.warn(sanitizedMessage)

        return this.dialog.open(ErrorComponent, {
            data: { title, message },
            disableClose: true
        })
    }
}
