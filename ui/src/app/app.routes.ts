import { Routes } from '@angular/router'
import { AdminDashboardComponent } from './pages/admin-dashboard/admin-dashboard.component'
import { CompetitionDetailComponent } from './pages/competition-detail/competition-detail.component'
import { CompetitionListComponent } from './pages/competition-list/competition-list.component'
import { GameDetailComponent } from './pages/game-detail/game-detail.component'
import { GameListComponent } from './pages/game-list/game-list.component'
import { ScheduleGameComponent } from './pages/schedule-game/schedule-game.component'
import { ScheduleComponent } from './pages/schedule/schedule.component'
import { SeasonDetailComponent } from './pages/season-detail/season-detail.component'
import { SeasonListComponent } from './pages/season-list/season-list.component'
import { TeamDetailComponent } from './pages/team-detail/team-detail.component'
import { TeamListComponent } from './pages/team-list/team-list.component'

// prettier-ignore
export const routes: Routes = [
    { path: '', redirectTo: '/schedule', pathMatch: 'full' },

    { path: 'schedule', component: ScheduleComponent },
    { path: 'schedule/competitions/:competition-id/seasons/:season-id/games/:game-id', component: ScheduleGameComponent },

    { path: 'admin', component: AdminDashboardComponent },
    
    // Admin - Competitions
    { path: 'admin/competitions', component: CompetitionListComponent },
    { path: 'admin/competitions/create', component: CompetitionDetailComponent },
    { path: 'admin/competitions/:competition-id', component: CompetitionDetailComponent },

    // Admin - Seasons
    { path: 'admin/competitions/:competition-id/seasons', component: SeasonListComponent },
    { path: 'admin/competitions/:competition-id/seasons/create', component: SeasonDetailComponent },
    { path: 'admin/competitions/:competition-id/seasons/:season-id', component: SeasonDetailComponent },

    // Admin - Games
    { path: 'admin/competitions/:competition-id/seasons/:season-id/games', component: GameListComponent },
    { path: 'admin/competitions/:competition-id/seasons/:season-id/games/create', component: GameDetailComponent },
    { path: 'admin/competitions/:competition-id/seasons/:season-id/games/:game-id', component: GameDetailComponent },

    // Admin - Teams
    { path: 'admin/teams', component: TeamListComponent },
    { path: 'admin/teams/create', component: TeamDetailComponent },
    { path: 'admin/teams/:team-id', component: TeamDetailComponent },
]
