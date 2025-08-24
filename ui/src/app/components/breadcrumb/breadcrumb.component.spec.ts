import { ComponentFixture, TestBed } from '@angular/core/testing'
import { BreadcrumbComponent } from './breadcrumb.component'
import { ActivatedRoute, provideRouter, Router } from '@angular/router'
import { By } from '@angular/platform-browser'

describe('BreadcrumbComponent', () => {
    let component: BreadcrumbComponent
    let fixture: ComponentFixture<BreadcrumbComponent>
    let router: Router

    function mockRoute(competitionId: string | null, seasonId: string | null, gameId: string | null) {
        return {
            snapshot: {
                paramMap: {
                    get: (key: string) => {
                        if (key === 'competition-id') return competitionId
                        if (key === 'season-id') return seasonId
                        if (key === 'game-id') return gameId
                        return null
                    }
                }
            }
        }
    }

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [BreadcrumbComponent],
            providers: [provideRouter([{ path: 'admin', component: BreadcrumbComponent }])]
        }).compileComponents()
    })

    describe('with no route params', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute(null, null, null) })
            fixture = TestBed.createComponent(BreadcrumbComponent)
            component = fixture.componentInstance
            router = TestBed.inject(Router)
            fixture.detectChanges()
        })

        it('should create', () => {
            expect(component).toBeTruthy()
        })

        it('should render only "Admin" breadcrumb', () => {
            const items = fixture.debugElement.queryAll(By.css('.breadcrumb-item'))
            expect(items.length).toBe(1)
            expect(items[0].nativeElement.textContent).toContain('Admin')
        })
    })

    describe('with competitionID only', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute('comp1', null, null) })
            fixture = TestBed.createComponent(BreadcrumbComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should render Admin and Competitions breadcrumbs', () => {
            const items = fixture.debugElement.queryAll(By.css('.breadcrumb-item'))
            expect(items.length).toBe(2)
            expect(items[0].nativeElement.textContent).toContain('Admin')
            expect(items[1].nativeElement.textContent).toContain('Competitions')
        })
    })

    describe('with full route params', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, { useValue: mockRoute('comp1', 'season1', 'game1') })
            fixture = TestBed.createComponent(BreadcrumbComponent)
            component = fixture.componentInstance
            router = TestBed.inject(Router)
            fixture.detectChanges()
        })

        it('should render all four breadcrumbs', () => {
            const items = fixture.debugElement.queryAll(By.css('.breadcrumb-item'))
            expect(items.length).toBe(4)
            expect(items[0].nativeElement.textContent).toContain('Admin')
            expect(items[1].nativeElement.textContent).toContain('Competitions')
            expect(items[2].nativeElement.textContent).toContain('Seasons')
            expect(items[3].nativeElement.textContent).toContain('Games')
        })

        it('should navigate correctly when breadcrumb links are clicked', () => {
            const routerSpy = spyOn(router, 'navigateByUrl')

            const links: HTMLElement[] = Array.from(
                fixture.nativeElement.querySelectorAll('.breadcrumb-item a')
            )

            // Admin link
            links[0].click()
            expect(routerSpy.calls.all()[0].args[0].toString()).toBe('/admin')

            // Competitions link
            links[1].click()
            expect(routerSpy.calls.all()[1].args[0].toString()).toBe('/admin/competitions')

            // Seasons link
            links[2].click()
            expect(routerSpy.calls.all()[2].args[0].toString()).toBe('/admin/competitions/comp1/seasons')

            // Games link
            links[3].click()
            expect(routerSpy.calls.all()[3].args[0].toString()).toBe(
                '/admin/competitions/comp1/seasons/season1/games'
            )
        })
    })
})
