import { Component, inject, OnInit } from '@angular/core'
import { GamesService } from '../../services/games/games.service'
import { MatTableDataSource } from '@angular/material/table'
import { Game } from '../../types/api'
import { CommonModule } from '@angular/common'
import { MaterialModule } from '../../shared/material/material.module'
import { ActivatedRoute, RouterModule } from '@angular/router'

@Component({
    selector: 'app-game-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule],
    templateUrl: './game-list.component.html',
    styleUrl: './game-list.component.scss'
})
export class GameListComponent implements OnInit {
    private readonly route = inject(ActivatedRoute)
    private readonly gamesService = inject(GamesService)

    dataSource = new MatTableDataSource<Game>()

    public competitionId: string | null = null
    public seasonId: string | null = null

    ngOnInit(): void {
        this.competitionId = this.route.snapshot.paramMap.get('competition-id')
        this.seasonId = this.route.snapshot.paramMap.get('season-id')

        if (this.competitionId && this.seasonId) {
            this.loadGames(this.competitionId, this.seasonId)
        }
    }

    private loadGames(competitionId: string, seasonId: string): void {
        this.gamesService.getGames(competitionId, seasonId).subscribe({
            next: (games) => {
                this.dataSource.data = games
            },
            error: (err) => console.error('Error loading games:', err)
        })
    }
}
