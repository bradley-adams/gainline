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

        // Competitions list is always included
        this.breadcrumbItems.push({
            label: 'Competitions',
            url: ['/admin/competitions']
        })

        if (!this.competitionID) return

        // Specific Competition
        this.breadcrumbItems.push({
            label: this.competitionID,
            url: ['/admin/competitions', this.competitionID]
        })

        // Seasons list for this competition
        this.breadcrumbItems.push({
            label: 'Seasons',
            url: ['/admin/competitions', this.competitionID, 'seasons']
        })

        if (!this.seasonId) return

        // Specific Season
        this.breadcrumbItems.push({
            label: this.seasonId,
            url: ['/admin/competitions', this.competitionID, 'seasons', this.seasonId]
        })

        // Games list for this season
        this.breadcrumbItems.push({
            label: 'Games',
            url: ['/admin/competitions', this.competitionID, 'seasons', this.seasonId, 'games']
        })

        if (!this.gameID) return

        // Specific Game
        this.breadcrumbItems.push({
            label: this.gameID,
            url: ['/admin/competitions', this.competitionID, 'seasons', this.seasonId, 'games', this.gameID]
        })
    }
}
