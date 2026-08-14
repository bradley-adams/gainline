import { provideHttpClient, withInterceptors } from '@angular/common/http'
import { ApplicationConfig } from '@angular/core'
import { provideAnimations } from '@angular/platform-browser/animations'
import { provideRouter } from '@angular/router'
import { provideAuth0 } from '@auth0/auth0-angular'

import { environment } from '../environments/environment'
import { routes } from './app.routes'
import { authInterceptor } from './interceptors/auth.interceptor'

export const appConfig: ApplicationConfig = {
    providers: [
        provideRouter(routes),
        provideAnimations(),
        provideHttpClient(withInterceptors([authInterceptor])),
        provideAuth0({
            domain: environment.auth0Domain,
            clientId: environment.auth0ClientId,
            authorizationParams: {
                redirect_uri: window.location.origin,
                audience: environment.auth0Audience
            }
        })
    ]
}
