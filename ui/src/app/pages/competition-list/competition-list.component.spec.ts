import { ComponentFixture, TestBed } from '@angular/core/testing'
import { By } from '@angular/platform-browser'
import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { provideRouter, Router } from '@angular/router'
import { of, throwError } from 'rxjs'

import { CompetitionListComponent } from './competition-list.component'
import { CompetitionDetailComponent } from '../competition-detail/competition-detail.component'
import { CompetitionsService } from '../../services/competitions/competitions.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { Competition } from '../../types/api'

describe('CompetitionListComponent', () => {
    let component: CompetitionListComponent
    let fixture: ComponentFixture<CompetitionListComponent>
    let router: Router

    let competitionsService: jasmine.SpyObj<CompetitionsService>
    let notificationService: jasmine.SpyObj<NotificationService>

    const mockCompetitions: Competition[] = [
        {
            id: 'comp1',
            name: 'Competition 1',
            created_at: new Date('2023-01-01T00:00:00Z'),
            updated_at: new Date('2023-01-02T00:00:00Z')
        },
        {
            id: 'comp2',
            name: 'Competition 2',
            created_at: new Date('2023-01-03T00:00:00Z'),
            updated_at: new Date('2023-01-04T00:00:00Z')
        }
    ]

    beforeEach(async () => {
        competitionsService = jasmine.createSpyObj('CompetitionsService', ['getCompetitions'])
        competitionsService.getCompetitions.and.returnValue(of(mockCompetitions))

        notificationService = jasmine.createSpyObj('NotificationService', ['showErrorAndLog'])

        await TestBed.configureTestingModule({
            imports: [CompetitionListComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/competitions/create',
                        component: CompetitionDetailComponent
                    },
                    {
                        path: 'admin/competitions/:competition-id',
                        component: CompetitionDetailComponent
                    }
                ]),
                { provide: CompetitionsService, useValue: competitionsService },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(CompetitionListComponent)
        component = fixture.componentInstance
        router = TestBed.inject(Router)
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should load competitions into the table', () => {
        const rows = fixture.nativeElement.querySelectorAll('tr')
        expect(rows.length).toBe(3)

        expect(rows[1].textContent).toContain('Competition 1')
        expect(rows[2].textContent).toContain('Competition 2')
    })

    it('should display "No competitions found" when no data is available', () => {
        component.dataSource.data = []

        const noDataRow: HTMLElement = fixture.nativeElement.querySelector('tr.mat-row td.mat-cell')
        expect(noDataRow).toBeTruthy()
        expect(noDataRow.textContent).toContain('No competitions found')
    })

    it('should show error when competitions fail to load', () => {
        const mockError = new Error('Failed')
        competitionsService.getCompetitions.and.returnValue(throwError(() => mockError))

        component.ngOnInit()

        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Load Error',
            'Failed to load competitions',
            mockError
        )
    })

    it('should have a "Create Competition" button with correct link', () => {
        const button = fixture.debugElement.query(By.css('.actions button'))
        expect(button.attributes['ng-reflect-router-link']).toContain('/admin/competitions/create')
    })

    it('should display table correctly', () => {
        const tableRows = fixture.nativeElement.querySelectorAll('tr')
        expect(tableRows.length).toBe(mockCompetitions.length + 1)

        const headerRow = tableRows[0]
        expect(headerRow.cells[0].textContent).toContain('Competition Name')

        expect(headerRow.cells[1].innerHTML).toBe('Created At')
        expect(headerRow.cells[2].innerHTML).toBe('Actions')

        expect(tableRows[1].cells[0].textContent).toBe('Competition 1')
        expect(tableRows[2].cells[0].textContent).toBe('Competition 2')

        expect(tableRows[1].cells[1].textContent).toBe('Jan 1, 2023, 1:00:00 PM')
        expect(tableRows[2].cells[1].textContent).toBe('Jan 3, 2023, 1:00:00 PM')

        expect(tableRows[1].cells[2].textContent).toBe('editcalendar_today')
        expect(tableRows[2].cells[2].textContent).toBe('editcalendar_today')
    })

    it('should navigate correctly when edit and seasons buttons are clicked', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')

        const tableRows = fixture.nativeElement.querySelectorAll('tr')
        const editButton = tableRows[1].querySelector('[data-testid="edit-button"]')
        const seasonsButton = tableRows[1].querySelector('[data-testid="seasons-button"]')

        editButton.click()
        const editCall = routerSpy.calls.all()[0].args[0].toString()
        expect(editCall).toEqual('/admin/competitions/comp1')

        seasonsButton.click()
        const seasonsCall = routerSpy.calls.all()[1].args[0].toString()
        expect(seasonsCall).toEqual('/admin/competitions/comp1/seasons')
    })
})
