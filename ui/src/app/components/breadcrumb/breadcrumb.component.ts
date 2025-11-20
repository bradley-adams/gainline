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
        const path = this.activatedRoute.snapshot.routeConfig?.path ?? ''

        const isSeasonsListPage = path.endsWith('seasons') && !this.seasonId && !this.gameID
        const isSeasonCreatePage = path.includes('seasons') && path.includes('create')
        const isGamesListPage = path.endsWith('games') && !this.gameID
        const isGameCreatePage = path.includes('games') && path.includes('create')

        this.breadcrumbItems.push({ label: 'Admin', url: ['/admin'] })
        this.breadcrumbItems.push({ label: 'Competitions', url: ['/admin/competitions'] })

        if (!this.competitionID) return

        this.breadcrumbItems.push({
            label: this.competitionID,
            url: ['/admin/competitions', this.competitionID]
        })

        if (isSeasonsListPage || isSeasonCreatePage || this.seasonId || this.gameID) {
            this.breadcrumbItems.push({
                label: 'Seasons',
                url: ['/admin/competitions', this.competitionID, 'seasons']
            })
        }

        if (!this.seasonId) return

        this.breadcrumbItems.push({
            label: this.seasonId,
            url: ['/admin/competitions', this.competitionID, 'seasons', this.seasonId]
        })

        if (isGamesListPage || isGameCreatePage || this.gameID) {
            this.breadcrumbItems.push({
                label: 'Games',
                url: ['/admin/competitions', this.competitionID, 'seasons', this.seasonId, 'games']
            })
        }

        if (!this.gameID) return

        this.breadcrumbItems.push({
            label: this.gameID,
            url: ['/admin/competitions', this.competitionID, 'seasons', this.seasonId, 'games', this.gameID]
        })
    }
}
