import { Component, inject } from '@angular/core'
import { CommonModule } from '@angular/common'
import { RouterModule } from '@angular/router'
import { MatTableDataSource } from '@angular/material/table'
import { MaterialModule } from '../../shared/material/material.module'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { Competition } from '../../types/api'

@Component({
    selector: 'app-competition-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule],
    templateUrl: './competition-list.component.html',
    styleUrl: './competition-list.component.scss'
})
export class CompetitionListComponent {
    private readonly competitionsService = inject(CompetitionsService)
    private readonly notificationService = inject(NotificationService)

    public dataSource = new MatTableDataSource<Competition>([])

    ngOnInit(): void {
        this.loadCompetitions()
    }

    private loadCompetitions(): void {
        this.competitionsService.getCompetitions().subscribe({
            next: (competitions) => {
                this.dataSource.data = competitions
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load competitions', err)
            }
        })
    }

    public confirmDelete(comp: Competition): void {
        const confirmationMessage = `Are you sure you want to delete competition "${comp.name}"?`
        this.notificationService
            .showConfirm('Confirm Delete', confirmationMessage)
            .afterClosed()
            .subscribe((confirmed) => {
                if (confirmed) {
                    this.removeCompetition(comp.id)
                }
            })
    }

    private removeCompetition(id: string): void {
        this.competitionsService.deleteCompetition(id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Competition deleted successfully')
                this.loadCompetitions()
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Delete Error', 'Failed to delete competition', err)
            }
        })
    }
}
