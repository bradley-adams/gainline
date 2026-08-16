import { Routes } from '@angular/router'
import { AuthGuard } from '@auth0/auth0-angular'
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

export const routes: Routes = [
    { path: '', redirectTo: '/schedule', pathMatch: 'full' },

    { path: 'schedule', component: ScheduleComponent, canActivate: [AuthGuard] },
    {
        path: 'schedule/competitions/:competition-id/seasons/:season-id/games/:game-id',
        component: ScheduleGameComponent,
        canActivate: [AuthGuard]
    },

    { path: 'admin', component: AdminDashboardComponent, canActivate: [AuthGuard] },

    // Admin - Competitions
    { path: 'admin/competitions', component: CompetitionListComponent, canActivate: [AuthGuard] },
    { path: 'admin/competitions/create', component: CompetitionDetailComponent, canActivate: [AuthGuard] },
    {
        path: 'admin/competitions/:competition-id',
        component: CompetitionDetailComponent,
        canActivate: [AuthGuard]
    },

    // Admin - Seasons
    {
        path: 'admin/competitions/:competition-id/seasons',
        component: SeasonListComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'admin/competitions/:competition-id/seasons/create',
        component: SeasonDetailComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'admin/competitions/:competition-id/seasons/:season-id',
        component: SeasonDetailComponent,
        canActivate: [AuthGuard]
    },

    // Admin - Games
    {
        path: 'admin/competitions/:competition-id/seasons/:season-id/games',
        component: GameListComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'admin/competitions/:competition-id/seasons/:season-id/games/create',
        component: GameDetailComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'admin/competitions/:competition-id/seasons/:season-id/games/:game-id',
        component: GameDetailComponent,
        canActivate: [AuthGuard]
    },

    // Admin - Teams
    { path: 'admin/teams', component: TeamListComponent, canActivate: [AuthGuard] },
    { path: 'admin/teams/create', component: TeamDetailComponent, canActivate: [AuthGuard] },
    { path: 'admin/teams/:team-id', component: TeamDetailComponent, canActivate: [AuthGuard] }
]
