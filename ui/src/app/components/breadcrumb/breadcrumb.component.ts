import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { ActivatedRoute, RouterModule } from '@angular/router'

interface BreadcrumbItem {
    label: string
    url: string[]
}

@Component({
    selector: 'app-breadcrumb',
    standalone: true,
    imports: [CommonModule, RouterModule],
    templateUrl: './breadcrumb.component.html',
    styleUrls: ['./breadcrumb.component.scss']
})
export class BreadcrumbComponent {
    private readonly activatedRoute: ActivatedRoute = inject(ActivatedRoute)

    public competitionID: string | null = null
    public seasonId: string | null = null
    public gameID: string | null = null

    public readonly breadcrumbItems: BreadcrumbItem[] = []

    ngOnInit(): void {
        this.competitionID = this.activatedRoute.snapshot.paramMap.get('competition-id')
        this.seasonId = this.activatedRoute.snapshot.paramMap.get('season-id')
        this.gameID = this.activatedRoute.snapshot.paramMap.get('game-id')

        this.buildBreadcrumbs()
    }

    private buildBreadcrumbs(): void {
        this.breadcrumbItems.push({ label: 'Admin', url: ['/admin'] })

        if (!this.competitionID) return
        this.breadcrumbItems.push({ label: 'Competitions', url: ['/admin/competitions'] })

        if (!this.seasonId) return
        this.breadcrumbItems.push({
            label: 'Seasons',
            url: ['/admin/competitions', this.competitionID!, 'seasons']
        })

        if (!this.gameID) return
        this.breadcrumbItems.push({
            label: 'Games',
            url: ['/admin/competitions', this.competitionID!, 'seasons', this.seasonId!, 'games']
        })
    }
}
