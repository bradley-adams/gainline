import { HttpInterceptorFn } from '@angular/common/http'
import { inject } from '@angular/core'
import { AuthService } from '@auth0/auth0-angular'
import { catchError, switchMap } from 'rxjs'

export const authInterceptor: HttpInterceptorFn = (req, next) => {
    const auth = inject(AuthService)

    return auth.getAccessTokenSilently().pipe(
        switchMap((token) => {
            const authReq = req.clone({
                setHeaders: {
                    Authorization: `Bearer ${token}`
                }
            })
            return next(authReq)
        }),
        catchError(() => next(req))
    )
}
