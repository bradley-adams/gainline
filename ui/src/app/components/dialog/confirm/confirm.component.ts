import { Component, inject } from '@angular/core'
import { MAT_DIALOG_DATA } from '@angular/material/dialog'
import { MaterialModule } from '../../../shared/material/material.module'

@Component({
    selector: 'app-confirm',
    standalone: true,
    imports: [MaterialModule],
    templateUrl: './confirm.component.html',
    styleUrls: ['./confirm.component.scss']
})
export class ConfirmComponent {
    private readonly data = inject(MAT_DIALOG_DATA)

    title: string = this.data?.title ?? 'Confirm'
    message: string = this.data?.message ?? 'Are you sure?'
}
