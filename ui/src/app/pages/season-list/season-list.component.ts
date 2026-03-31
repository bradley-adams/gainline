import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { MatPaginator, PageEvent } from '@angular/material/paginator'
import { MatTableDataSource } from '@angular/material/table'
import { ActivatedRoute, RouterModule } from '@angular/router'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { MaterialModule } from '../../shared/material/material.module'
import { Season } from '../../types/api'

@Component({
    selector: 'app-season-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule, BreadcrumbComponent, MatPaginator],
    templateUrl: './season-list.component.html',
    styleUrl: './season-list.component.scss'
})
export class SeasonListComponent {
    private readonly route = inject(ActivatedRoute)
    private readonly seasonsService = inject(SeasonsService)
    private readonly notificationService = inject(NotificationService)

    public dataSource = new MatTableDataSource<Season>([])
    public competitionId: string | null = null

    public total = 0
    public page = 1
    public pageSize = 10

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')

        if (this.competitionId) {
            this.loadSeasons(this.competitionId)
        }
    }

    public onPageChange(event: PageEvent): void {
        this.page = event.pageIndex + 1
        this.pageSize = event.pageSize
        if (this.competitionId) {
            this.loadSeasons(this.competitionId)
        }
    }

    private loadSeasons(competitionId: string): void {
        this.seasonsService.getSeasons(competitionId, this.page, this.pageSize).subscribe({
            next: (response) => {
                this.dataSource.data = response.data
                this.total = response.pagination.total
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load seasons', err)
            }
        })
    }

    confirmDelete(season: Season): void {
        this.notificationService
            .showConfirm('Confirm Delete', 'Are you sure you want to delete this season?')
            .afterClosed()
            .subscribe((confirmed) => {
                if (confirmed && this.competitionId) {
                    this.deleteSeason(this.competitionId, season.id)
                }
            })
    }

    private deleteSeason(competitionId: string, seasonId: string): void {
        this.seasonsService.deleteSeason(competitionId, seasonId).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Season deleted successfully')
                this.loadSeasons(competitionId)
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Delete Error', 'Failed to delete season', err)
            }
        })
    }
}
