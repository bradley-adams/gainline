import { Routes } from '@angular/router';
import { ScheduleComponent } from './pages/schedule/schedule.component';

export const routes: Routes = [
  { path: '', redirectTo: '/schedule', pathMatch: 'full' },
  { path: 'schedule', component: ScheduleComponent },
]

