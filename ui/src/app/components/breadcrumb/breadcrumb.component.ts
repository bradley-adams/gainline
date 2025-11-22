import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { ActivatedRoute, RouterModule } from '@angular/router'
import { CompetitionsService } from '../../services/competitions/competitions.service'

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
    private readonly route = inject(ActivatedRoute)
    private readonly competitionsService = inject(CompetitionsService)

    public breadcrumbItems: BreadcrumbItem[] = []

    private competitionId: string = ''
    private seasonId: string = ''
    private gameId: string = ''

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id') ?? ''
        this.seasonId = this.route.snapshot.paramMap.get('season-id') ?? ''
        this.gameId = this.route.snapshot.paramMap.get('game-id') ?? ''

        this.breadcrumbItems.push({ label: 'Admin', url: ['/admin'] })
        this.breadcrumbItems.push({ label: 'Competitions', url: ['/admin/competitions'] })

        if (!this.competitionId) return

        this.competitionsService.getCompetition(this.competitionId).subscribe({
            next: (competition) => {
                const compName = competition.name ?? this.competitionId
                this.breadcrumbItems.push({
                    label: compName,
                    url: ['/admin/competitions', this.competitionId]
                })
                this.buildBreadcrumb()
            },
            error: () => {
                this.breadcrumbItems.push({
                    label: this.competitionId,
                    url: ['/admin/competitions', this.competitionId]
                })
                this.buildBreadcrumb()
            }
        })
    }

    private buildBreadcrumb(): void {
        const path = this.route.snapshot.routeConfig?.path ?? ''

        const isSeasonsListPage = path.endsWith('seasons') && !this.seasonId && !this.gameId
        const isSeasonCreatePage = path.includes('seasons') && path.includes('create')
        const isGamesListPage = path.endsWith('games') && !this.gameId
        const isGameCreatePage = path.includes('games') && path.includes('create')

        if (isSeasonsListPage || isSeasonCreatePage || this.seasonId || this.gameId) {
            this.breadcrumbItems.push({
                label: 'Seasons',
                url: ['/admin/competitions', this.competitionId, 'seasons']
            })
        }

        if (this.seasonId) {
            this.breadcrumbItems.push({
                label: this.seasonId,
                url: ['/admin/competitions', this.competitionId, 'seasons', this.seasonId]
            })
        }

        if (isGamesListPage || isGameCreatePage || this.gameId) {
            this.breadcrumbItems.push({
                label: 'Games',
                url: ['/admin/competitions', this.competitionId, 'seasons', this.seasonId, 'games']
            })
        }

        if (this.gameId) {
            this.breadcrumbItems.push({
                label: this.gameId,
                url: [
                    '/admin/competitions',
                    this.competitionId,
                    'seasons',
                    this.seasonId,
                    'games',
                    this.gameId
                ]
            })
        }
    }
}
