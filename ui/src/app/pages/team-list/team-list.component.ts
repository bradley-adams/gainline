import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { MatPaginator, PageEvent } from '@angular/material/paginator'
import { MatTableDataSource } from '@angular/material/table'
import { RouterModule } from '@angular/router'
import { NotificationService } from '../../services/notifications/notifications.service'
import { TeamsService } from '../../services/teams/teams.service'
import { MaterialModule } from '../../shared/material/material.module'
import { Team } from '../../types/api'

@Component({
    selector: 'app-team-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule, MatPaginator],
    templateUrl: './team-list.component.html',
    styleUrl: './team-list.component.scss'
})
export class TeamListComponent {
    private readonly teamsService = inject(TeamsService)
    private readonly notificationService = inject(NotificationService)

    public dataSource = new MatTableDataSource<Team>([])

    public total = 0
    public page = 1
    public pageSize = 10

    ngOnInit(): void {
        this.loadTeams()
    }

    public onPageChange(event: PageEvent): void {
        this.page = event.pageIndex + 1
        this.pageSize = event.pageSize
        this.loadTeams()
    }

    private loadTeams(): void {
        this.teamsService.getTeamsPaginated(this.page, this.pageSize).subscribe({
            next: (response) => {
                this.dataSource.data = response.data
                this.total = response.pagination.total
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Load Error', 'Failed to load teams', err)
            }
        })
    }

    confirmDelete(team: Team): void {
        const confirmationMessage = `Are you sure you want to delete ${team.name}?`

        this.notificationService
            .showConfirm('Confirm Delete', confirmationMessage)
            .afterClosed()
            .subscribe((confirmed) => {
                if (confirmed) {
                    this.deleteTeam(team.id)
                }
            })
    }

    private deleteTeam(id: string): void {
        this.teamsService.deleteTeam(id).subscribe({
            next: () => {
                this.notificationService.showSnackbar('Team deleted successfully')
                this.loadTeams()
            },
            error: (err) => {
                this.notificationService.showErrorAndLog('Delete Error', 'Failed to delete team', err)
            }
        })
    }
}
