import { Component, inject } from '@angular/core'
import { MatSnackBarRef, MAT_SNACK_BAR_DATA } from '@angular/material/snack-bar'
import { MatIconModule } from '@angular/material/icon'
import { MatButtonModule } from '@angular/material/button'

@Component({
    selector: 'app-snackbar',
    standalone: true,
    imports: [MatIconModule, MatButtonModule],
    templateUrl: './snackbar.component.html',
    styleUrls: ['./snackbar.component.scss']
})
export class SnackbarComponent {
    snackBarRef = inject(MatSnackBarRef<SnackbarComponent>)
    data = inject(MAT_SNACK_BAR_DATA)

    close(): void {
        this.snackBarRef.dismiss()
    }
}
