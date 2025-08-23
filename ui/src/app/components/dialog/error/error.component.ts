import { Component, inject } from '@angular/core'
import { MAT_DIALOG_DATA } from '@angular/material/dialog'
import { MaterialModule } from '../../../shared/material/material.module'

@Component({
    selector: 'app-error',
    standalone: true,
    imports: [MaterialModule],
    templateUrl: './error.component.html',
    styleUrls: ['./error.component.scss']
})
export class ErrorComponent {
    private readonly data = inject(MAT_DIALOG_DATA)

    title: string = this.data?.title ?? 'Error'
    message: string = this.data?.message ?? 'Sorry, an unexpected error occurred. Please try again later.'
}
