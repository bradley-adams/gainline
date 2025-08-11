import { Routes } from '@angular/router';
import { ScheduleComponent } from './pages/schedule/schedule.component';
import { CompetitionComponent } from './pages/competition/competition.component';

export const routes: Routes = [
  { path: '', redirectTo: '/schedule', pathMatch: 'full' },

  { path: 'schedule', component: ScheduleComponent },

  { path: 'competitions/create', component: CompetitionComponent },
    
  { path: 'competitions/:competition-id', component: CompetitionComponent },
]

