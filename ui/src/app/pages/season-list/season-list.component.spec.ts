import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ComponentFixture, TestBed } from '@angular/core/testing'
import { By } from '@angular/platform-browser'
import { ActivatedRoute, convertToParamMap, provideRouter, Router } from '@angular/router'
import { of, throwError } from 'rxjs'

import { provideHttpClient } from '@angular/common/http'
import { NotificationService } from '../../services/notifications/notifications.service'
import { SeasonsService } from '../../services/seasons/seasons.service'
import { Season, Stage, StageType, Team } from '../../types/api'
import { SeasonDetailComponent } from '../season-detail/season-detail.component'
import { SeasonListComponent } from './season-list.component'

describe('SeasonListComponent', () => {
    let component: SeasonListComponent
    let fixture: ComponentFixture<SeasonListComponent>
    let router: Router

    let seasonsService: jasmine.SpyObj<SeasonsService>
    let notificationService: jasmine.SpyObj<NotificationService>

    const mockTeams: Team[] = [
        {
            id: 'team1',
            abbreviation: 'T1',
            location: 'City A',
            name: 'Team One',
            created_at: new Date('2024-01-01T00:00:00Z'),
            updated_at: new Date('2024-01-01T00:00:00Z')
        },
        {
            id: 'team2',
            abbreviation: 'T2',
            location: 'City B',
            name: 'Team Two',
            created_at: new Date('2024-01-02T00:00:00Z'),
            updated_at: new Date('2024-01-02T00:00:00Z')
        },
        {
            id: 'team3',
            abbreviation: 'T3',
            location: 'City C',
            name: 'Team Three',
            created_at: new Date('2024-01-03T00:00:00Z'),
            updated_at: new Date('2024-01-03T00:00:00Z')
        },
        {
            id: 'team4',
            abbreviation: 'T4',
            location: 'City D',
            name: 'Team Four',
            created_at: new Date('2024-01-04T00:00:00Z'),
            updated_at: new Date('2024-01-04T00:00:00Z')
        }
    ]

    const MockEmptyStages: Stage[] = []
    const MockStages: Stage[] = [
        {
            id: 'stage-round-1',
            name: 'Round 1',
            stage_type: StageType.StageTypeRegular,
            order_index: 1,
            created_at: new Date('2025-01-01T00:00:00Z'),
            updated_at: new Date('2025-01-01T00:00:00Z')
        },
        {
            id: 'stage-round-2',
            name: 'Round 2',
            stage_type: StageType.StageTypeRegular,
            order_index: 2,
            created_at: new Date('2025-01-01T00:00:00Z'),
            updated_at: new Date('2025-01-01T00:00:00Z')
        },
        {
            id: 'stage-round-3',
            name: 'Final',
            stage_type: StageType.StageTypeFinals,
            order_index: 3,
            created_at: new Date('2025-01-01T00:00:00Z'),
            updated_at: new Date('2025-01-01T00:00:00Z')
        }
    ]

    const mockSeasons: Season[] = [
        {
            id: 'season1',
            competition_id: 'comp1',
            start_date: new Date('2025-01-01T00:00:00Z'),
            end_date: new Date('2025-12-31T23:59:59Z'),
            stages: MockStages,
            teams: mockTeams,
            created_at: new Date('2024-12-01T00:00:00Z'),
            updated_at: new Date('2024-12-01T00:00:00Z')
        },
        {
            id: 'season2',
            competition_id: 'comp1',
            start_date: new Date('2024-01-01T00:00:00Z'),
            end_date: new Date('2024-12-31T23:59:59Z'),
            stages: MockEmptyStages,
            teams: mockTeams,
            created_at: new Date('2023-12-01T00:00:00Z'),
            updated_at: new Date('2023-12-01T00:00:00Z')
        }
    ]

    beforeEach(async () => {
        seasonsService = jasmine.createSpyObj('SeasonsService', ['getSeasons', 'deleteSeason'])
        seasonsService.getSeasons.and.returnValue(of(mockSeasons))
        seasonsService.deleteSeason.and.returnValue(of(undefined))

        notificationService = jasmine.createSpyObj('NotificationService', [
            'showConfirm',
            'showErrorAndLog',
            'showWarnAndLog',
            'showSnackbar'
        ])

        await TestBed.configureTestingModule({
            imports: [SeasonListComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/:competition-id/seasons/create',
                        component: SeasonDetailComponent
                    },
                    {
                        path: 'admin/competitions/:competition-id/seasons/:season-id',
                        component: SeasonDetailComponent
                    }
                ]),
                { provide: SeasonsService, useValue: seasonsService },
                {
                    provide: ActivatedRoute,
                    useValue: { snapshot: { paramMap: convertToParamMap({ 'competition-id': 'comp1' }) } }
                },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(SeasonListComponent)
        component = fixture.componentInstance
        router = TestBed.inject(Router)
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should load seasons into the table on init', () => {
        const rows = fixture.nativeElement.querySelectorAll('tr')
        expect(rows.length).toBe(3)
        expect(rows[1].textContent).toContain('season1')
        expect(rows[2].textContent).toContain('season2')
    })

    it('should display "No seasons found" row when dataSource is empty', () => {
        component.dataSource.data = []
        const noDataRow: HTMLElement = fixture.nativeElement.querySelector('tr.mat-row td.mat-cell')
        expect(noDataRow).toBeTruthy()
        expect(noDataRow.textContent).toContain('No seasons found')
    })

    it('should show error when getSeasons fails', () => {
        const mockError = new Error('Failed')
        seasonsService.getSeasons.and.returnValue(throwError(() => mockError))
        component.ngOnInit()
        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Load Error',
            'Failed to load seasons',
            mockError
        )
    })

    it('should navigate to create season page when button is clicked', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')
        const button = fixture.debugElement.query(By.css('.actions button'))
        button.nativeElement.click()
        expect(routerSpy).toHaveBeenCalled()
    })

    it('should display season table with correct headers and data', () => {
        const tableRows = fixture.nativeElement.querySelectorAll('tr')
        expect(tableRows.length).toBe(mockSeasons.length + 1)

        const headerRow = tableRows[0]
        expect(headerRow.cells[0].innerHTML).toBe('Season ID')
        expect(headerRow.cells[1].innerHTML).toBe('Starts')
        expect(headerRow.cells[2].innerHTML).toBe('Ends')
        expect(headerRow.cells[3].innerHTML).toBe('Stages')
        expect(headerRow.cells[4].innerHTML).toBe('Actions')

        expect(tableRows[1].cells[0].textContent).toBe('season1')
        expect(tableRows[2].cells[0].textContent).toBe('season2')

        expect(tableRows[1].cells[1].textContent).toBe('Jan 1, 2025, 1:00:00 PM')
        expect(tableRows[2].cells[1].textContent).toBe('Jan 1, 2024, 1:00:00 PM')

        expect(tableRows[1].cells[2].textContent).toBe('Jan 1, 2026, 12:59:59 PM')
        expect(tableRows[2].cells[2].textContent).toBe('Jan 1, 2025, 12:59:59 PM')

        expect(Number(tableRows[1].cells[3].textContent)).toBe(3)
        expect(tableRows[2].cells[3].textContent).toBe('No stages')

        expect(tableRows[1].cells[4].textContent).toContain('edit')
        expect(tableRows[1].cells[4].textContent).toContain('calendar_today')
        expect(tableRows[1].cells[4].textContent).toContain('delete')

        expect(tableRows[2].cells[4].textContent).toContain('edit')
        expect(tableRows[2].cells[4].textContent).toContain('calendar_today')
        expect(tableRows[2].cells[4].textContent).toContain('delete')
    })

    it('should navigate correctly when edit and games buttons are clicked', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')
        const tableRows = fixture.nativeElement.querySelectorAll('tr')

        const editButton = tableRows[1].querySelector('[data-testid="edit-button"]')
        const seasonsButton = tableRows[1].querySelector('[data-testid="games-button"]')

        editButton.click()
        const call = routerSpy.calls.all()[0].args[0].toString()
        expect(call).toEqual('/admin/competitions/comp1/seasons/season1')

        seasonsButton.click()
        const call2 = routerSpy.calls.all()[1].args[0].toString()
        expect(call2).toEqual('/admin/competitions/comp1/seasons/season1/games')
    })

    it('should call deleteSeason when confirmed', () => {
        notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
        component.confirmDelete(mockSeasons[0])
        expect(seasonsService.deleteSeason).toHaveBeenCalledWith('comp1', 'season1')
        expect(notificationService.showSnackbar).toHaveBeenCalledWith('Season deleted successfully')
    })

    it('should not call deleteSeason when cancelled', () => {
        notificationService.showConfirm.and.returnValue({ afterClosed: () => of(false) } as any)
        component.confirmDelete(mockSeasons[0])
        expect(seasonsService.deleteSeason).not.toHaveBeenCalled()
        expect(notificationService.showSnackbar).not.toHaveBeenCalled()
    })

    it('should show error if deleteSeason fails', () => {
        const mockError = new Error('Failed')
        seasonsService.deleteSeason.and.returnValue(throwError(() => mockError))
        notificationService.showConfirm.and.returnValue({ afterClosed: () => of(true) } as any)
        component.confirmDelete(mockSeasons[0])
        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Delete Error',
            'Failed to delete season',
            mockError
        )
    })
})
