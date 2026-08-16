import { CommonModule } from '@angular/common'
import { Component, inject } from '@angular/core'
import { MatButtonModule } from '@angular/material/button'
import { MatIconModule } from '@angular/material/icon'
import { MatToolbarModule } from '@angular/material/toolbar'
import { RouterModule } from '@angular/router'
import { AuthService } from '@auth0/auth0-angular'

@Component({
    selector: 'app-header',
    standalone: true,
    imports: [CommonModule, RouterModule, MatToolbarModule, MatButtonModule, MatIconModule],
    templateUrl: './header.component.html',
    styleUrls: ['./header.component.scss']
})
export class HeaderComponent {
    public readonly auth = inject(AuthService)

    public login(): void {
        this.auth.loginWithRedirect()
    }

    public logout(): void {
        this.auth.logout({
            logoutParams: {
                returnTo: window.location.origin
            }
        })
    }
}
