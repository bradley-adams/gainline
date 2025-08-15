import { Routes } from '@angular/router'
import { ScheduleComponent } from './pages/schedule/schedule.component'
import { CompetitionListComponent } from './pages/competition-list/competition-list.component'
import { CompetitionDetailComponent } from './pages/competition-detail/competition-detail.component'
import { SeasonListComponent } from './pages/season-list/season-list.component'
import { SeasonDetailComponent } from './pages/season-detail/season-detail.component'

// prettier-ignore
export const routes: Routes = [
    { path: '', redirectTo: '/schedule', pathMatch: 'full' },

    { path: 'schedule', component: ScheduleComponent },

    // Admin - Competitions
    { path: 'admin', component: CompetitionListComponent },
    { path: 'admin/competitions', component: CompetitionListComponent },
    { path: 'admin/competitions/create', component: CompetitionDetailComponent },
    { path: 'admin/competitions/:competition-id', component: CompetitionDetailComponent },

    // Admin - Seasons
    { path: 'admin/competitions/:competition-id/seasons', component: SeasonListComponent },
    { path: 'admin/competitions/:competition-id/seasons/create', component: SeasonDetailComponent },
    { path: 'admin/competitions/:competition-id/seasons/:season-id', component: SeasonDetailComponent }
];
