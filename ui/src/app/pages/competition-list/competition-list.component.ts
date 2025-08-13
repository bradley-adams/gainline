import { Component, inject } from '@angular/core'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { MatTableDataSource } from '@angular/material/table'
import { Competition } from '../../types/api'
import { CommonModule } from '@angular/common'
import { MaterialModule } from '../../shared/material/material.module'
import { RouterModule } from '@angular/router'

@Component({
    selector: 'app-competition-list',
    standalone: true,
    imports: [CommonModule, MaterialModule, RouterModule],
    templateUrl: './competition-list.component.html',
    styleUrl: './competition-list.component.scss'
})
export class CompetitionListComponent {
    private readonly competitionsService = inject(CompetitionsService)

    public dataSource = new MatTableDataSource<Competition>([])

    ngOnInit(): void {
        this.loadCompetitions()
    }

    private loadCompetitions(): void {
        this.competitionsService.getCompetitions().subscribe({
            next: (competitions) => {
                this.dataSource.data = competitions
            },
            error: (err) => console.error('Error loading competitions:', err)
        })
    }
}
