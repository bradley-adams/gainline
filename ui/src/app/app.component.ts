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
      <app-header></app-header>
      <router-outlet></router-outlet>
    `,
    styles: `
        .contents {
            box-sizing: border-box;
            background-color: transparent;
            margin: 0;
            position: absolute;
            height: auto;
            inset: 0;
            padding: 0;
        }
    `
})
export class AppComponent {}
