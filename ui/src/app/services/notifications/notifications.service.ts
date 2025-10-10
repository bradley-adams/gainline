import { Injectable, inject } from '@angular/core'
import { MatSnackBar } from '@angular/material/snack-bar'
import { MatDialog, MatDialogRef } from '@angular/material/dialog'
import { ConfirmComponent } from '../../components/dialog/confirm/confirm.component'
import { ErrorComponent } from '../../components/dialog/error/error.component'

@Injectable({
    providedIn: 'root'
})
export class NotificationService {
    private readonly dialog = inject(MatDialog)
    private readonly snackBar = inject(MatSnackBar)

    showSnackbar(message: string, buttonLabel = '', duration = 500000): void {
        this.snackBar.open(message, buttonLabel, { duration })
    }

    showConfirm(title: string, message: string) {
        return this.dialog.open(ConfirmComponent, {
            data: { title, message },
            disableClose: true
        })
    }

    showError(title: string, message: string) {
        return this.dialog.open(ErrorComponent, {
            data: { title, message },
            disableClose: true
        })
    }
}
