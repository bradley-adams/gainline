import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { MatPaginator, PageEvent } from '@angular/material/paginator'
import { MatTableDataSource } from '@angular/material/table'
import { RouterModule } from '@angular/router'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { MaterialModule } from '../../shared/material/material.module'
import { Competition } from '../../types/api'

@Component({
    selector: 'app-competition-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule, BreadcrumbComponent, MatPaginator],
    templateUrl: './competition-list.component.html',
    styleUrl: './competition-list.component.scss'
})
export class CompetitionListComponent {
    private readonly competitionsService = inject(CompetitionsService)
    private readonly notificationService = inject(NotificationService)

    public dataSource = new MatTableDataSource<Competition>([])
    public total = 0
    public page = 1
    public pageSize = 10

    ngOnInit(): void {
        this.loadCompetitions()
    }

    public onPageChange(event: PageEvent): void {
        this.page = event.pageIndex + 1
        this.pageSize = event.pageSize
        this.loadCompetitions()
    }

    private loadCompetitions(): void {
        this.competitionsService.getCompetitions(this.page, this.pageSize).subscribe({
            next: (response) => {
                this.dataSource.data = response.data
                this.total = response.pagination.total
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
