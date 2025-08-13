import { Routes } from '@angular/router'
import { ScheduleComponent } from './pages/schedule/schedule.component'
import { CompetitionListComponent } from './pages/competition-list/competition-list.component'
import { CompetitionDetailComponent } from './pages/competition-detail/competition-detail.component'

export const routes: Routes = [
    { path: '', redirectTo: '/schedule', pathMatch: 'full' },

    { path: 'schedule', component: ScheduleComponent },

    { path: 'admin', component: CompetitionListComponent },

    { path: 'admin/competitions', component: CompetitionListComponent },
    { path: 'admin/competitions/create', component: CompetitionDetailComponent },
    { path: 'admin/competitions/:competition-id', component: CompetitionDetailComponent }
]
