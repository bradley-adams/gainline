import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { MaterialModule } from '../../shared/material/material.module'
import { ActivatedRoute, RouterModule } from '@angular/router'
import { MatTableDataSource } from '@angular/material/table'
import { Season } from '../../types/api'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { BreadcrumbComponent } from '../../components/breadcrumb/breadcrumb.component'
import { NotificationService } from '../../services/notifications/notifications.service'

@Component({
    selector: 'app-season-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule, BreadcrumbComponent],
    templateUrl: './season-list.component.html',
    styleUrl: './season-list.component.scss'
})
export class SeasonListComponent {
    private readonly route = inject(ActivatedRoute)
    private readonly seasonsService = inject(SeasonsService)
    private readonly notificationService = inject(NotificationService)

    public dataSource = new MatTableDataSource<Season>([])

    public competitionId: string | null = null

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')

        if (this.competitionId) {
            this.loadSeasons(this.competitionId)
        }
    }

    private loadSeasons(competitionId: string): void {
        this.seasonsService.getSeasons(competitionId).subscribe({
            next: (seasons) => {
                this.dataSource.data = seasons
            },
            error: (err) => {
                console.error('Error loading seasons:', err)
                this.notificationService.showError('Load Error', 'Failed to load seasons')
            }
        })
    }
}
