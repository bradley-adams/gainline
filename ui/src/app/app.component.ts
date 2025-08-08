import { Component, inject } from '@angular/core'
import { CommonModule } from '@angular/common'
import { RouterOutlet } from '@angular/router'
import { MatSidenavModule } from '@angular/material/sidenav'
import { HeaderComponent } from './components/header/header.component'

@Component({
    selector: 'app-root',
    standalone: true,
    imports: [CommonModule, RouterOutlet, MatSidenavModule, HeaderComponent],
    template: `
        <div class="mat-typography app-frame mat-app-background">
            <app-header></app-header>
            <router-outlet></router-outlet>
        </div>
    `,
})
export class AppComponent {}
