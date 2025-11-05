import { Component, inject } from '@angular/core'
import { MatTableDataSource } from '@angular/material/table'
import { Team } from '../../types/api'
import { CommonModule } from '@angular/common'
import { MaterialModule } from '../../shared/material/material.module'
import { RouterModule } from '@angular/router'
import { TeamsService } from '../../services/teams/teams.service'
import { NotificationService } from '../../services/notifications/notifications.service'

@Component({
    selector: 'app-team-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule],
    templateUrl: './team-list.component.html',
    styleUrl: './team-list.component.scss'
})
export class TeamListComponent {
    private readonly teamsService = inject(TeamsService)
    private readonly notificationService = inject(NotificationService)

    public dataSource = new MatTableDataSource<Team>([])

    ngOnInit(): void {
        this.loadTeams()
    }

    private loadTeams(): void {
        this.teamsService.getTeams().subscribe({
            next: (teams) => {
                this.dataSource.data = teams
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
