import { ComponentFixture, TestBed } from '@angular/core/testing'
import { By } from '@angular/platform-browser'
import { ActivatedRoute, provideRouter, Router } from '@angular/router'
import { of } from 'rxjs'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { BreadcrumbComponent } from './breadcrumb.component'

describe('BreadcrumbComponent', () => {
    let component: BreadcrumbComponent
    let fixture: ComponentFixture<BreadcrumbComponent>
    let router: Router

    const mockCompetitionsService = {
        getCompetition: (id: string) => of({ id, name: `Competition ${id}` })
    }

    function mockRoute(
        competitionId: string | null,
        seasonId: string | null,
        gameId: string | null,
        path = ''
    ) {
        return {
            snapshot: {
                paramMap: {
                    get: (key: string) => {
                        if (key === 'competition-id') return competitionId
                        if (key === 'season-id') return seasonId
                        if (key === 'game-id') return gameId
                        return null
                    }
                },
                routeConfig: {
                    path
                }
            }
        }
    }

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [BreadcrumbComponent],
            providers: [
                provideRouter([{ path: 'admin', component: BreadcrumbComponent }]),
                { provide: CompetitionsService, useValue: mockCompetitionsService }
            ]
        }).compileComponents()
    })

    describe('with no route params', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, {
                useValue: mockRoute(null, null, null, 'admin')
            })
            fixture = TestBed.createComponent(BreadcrumbComponent)
            component = fixture.componentInstance
            router = TestBed.inject(Router)
            fixture.detectChanges()
        })

        it('should create', () => {
            expect(component).toBeTruthy()
        })

        it('should render Admin and Competitions breadcrumbs', () => {
            const items = fixture.debugElement.queryAll(By.css('.breadcrumb-item'))
            expect(items.length).toBe(2)
            expect(items[0].nativeElement.textContent).toContain('Admin')
            expect(items[1].nativeElement.textContent).toContain('Competitions')
        })
    })

    describe('with competitionID only', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, {
                useValue: mockRoute('comp1', null, null, 'admin/competitions/:competition-id/seasons')
            })
            fixture = TestBed.createComponent(BreadcrumbComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should render breadcrumbs up to Seasons', () => {
            const items = fixture.debugElement.queryAll(By.css('.breadcrumb-item'))
            expect(items.length).toBe(4)
            expect(items[0].nativeElement.textContent).toContain('Admin')
            expect(items[1].nativeElement.textContent).toContain('Competitions')
            expect(items[2].nativeElement.textContent).toContain('Competition comp1')
            expect(items[3].nativeElement.textContent).toContain('Seasons')
        })
    })

    describe('with competitionID and seasonID', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, {
                useValue: mockRoute(
                    'comp1',
                    'season1',
                    null,
                    'admin/competitions/:competition-id/seasons/:season-id/games'
                )
            })
            fixture = TestBed.createComponent(BreadcrumbComponent)
            component = fixture.componentInstance
            fixture.detectChanges()
        })

        it('should render breadcrumbs up to Games', () => {
            const items = fixture.debugElement.queryAll(By.css('.breadcrumb-item'))
            expect(items.length).toBe(6)
            expect(items[0].nativeElement.textContent).toContain('Admin')
            expect(items[1].nativeElement.textContent).toContain('Competitions')
            expect(items[2].nativeElement.textContent).toContain('Competition comp1')
            expect(items[3].nativeElement.textContent).toContain('Seasons')
            expect(items[4].nativeElement.textContent).toContain('season1')
            expect(items[5].nativeElement.textContent).toContain('Games')
        })
    })

    describe('with full route params', () => {
        beforeEach(() => {
            TestBed.overrideProvider(ActivatedRoute, {
                useValue: mockRoute(
                    'comp1',
                    'season1',
                    'game1',
                    'admin/competitions/:competition-id/seasons/:season-id/games/:game-id'
                )
            })

            fixture = TestBed.createComponent(BreadcrumbComponent)
            component = fixture.componentInstance
            router = TestBed.inject(Router)
            fixture.detectChanges()
        })

        it('should render all breadcrumbs including current game', () => {
            const items = fixture.debugElement.queryAll(By.css('.breadcrumb-item'))
            expect(items.length).toBe(7)
            expect(items[0].nativeElement.textContent).toContain('Admin')
            expect(items[1].nativeElement.textContent).toContain('Competitions')
            expect(items[2].nativeElement.textContent).toContain('Competition comp1')
            expect(items[3].nativeElement.textContent).toContain('Seasons')
            expect(items[4].nativeElement.textContent).toContain('season1')
            expect(items[5].nativeElement.textContent).toContain('Games')
            expect(items[6].nativeElement.textContent).toContain('game1')
        })

        it('should navigate correctly when breadcrumb links are clicked', () => {
            const routerSpy = spyOn(router, 'navigateByUrl')
            const links: HTMLElement[] = Array.from(
                fixture.nativeElement.querySelectorAll('.breadcrumb-item a')
            )

            links[0].click()
            expect(routerSpy.calls.all()[0].args[0].toString()).toBe('/admin')

            links[1].click()
            expect(routerSpy.calls.all()[1].args[0].toString()).toBe('/admin/competitions')

            links[2].click()
            expect(routerSpy.calls.all()[2].args[0].toString()).toBe('/admin/competitions/comp1')

            links[3].click()
            expect(routerSpy.calls.all()[3].args[0].toString()).toBe('/admin/competitions/comp1/seasons')

            links[4].click()
            expect(routerSpy.calls.all()[4].args[0].toString()).toBe(
                '/admin/competitions/comp1/seasons/season1'
            )

            links[5].click()
            expect(routerSpy.calls.all()[5].args[0].toString()).toBe(
                '/admin/competitions/comp1/seasons/season1/games'
            )

            links[6].click()
            expect(routerSpy.calls.all()[6].args[0].toString()).toBe(
                '/admin/competitions/comp1/seasons/season1/games/game1'
            )
        })
    })
})
