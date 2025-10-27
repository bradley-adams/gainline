import { provideHttpClient } from '@angular/common/http'
import { provideHttpClientTesting } from '@angular/common/http/testing'
import { ComponentFixture, TestBed } from '@angular/core/testing'
import { By } from '@angular/platform-browser'
import { provideRouter, Router } from '@angular/router'
import { of, throwError } from 'rxjs'

import { TeamListComponent } from './team-list.component'
import { TeamDetailComponent } from '../team-detail/team-detail.component'
import { TeamsService } from '../../services/teams/teams.service'
import { NotificationService } from '../../services/notifications/notifications.service'
import { Team } from '../../types/api'

describe('TeamListComponent', () => {
    let component: TeamListComponent
    let fixture: ComponentFixture<TeamListComponent>
    let router: Router
    let teamsService: jasmine.SpyObj<TeamsService>
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

    beforeEach(async () => {
        teamsService = jasmine.createSpyObj('TeamsService', ['getTeams'])
        teamsService.getTeams.and.returnValue(of(mockTeams))

        notificationService = jasmine.createSpyObj('NotificationService', ['showErrorAndLog'])

        await TestBed.configureTestingModule({
            imports: [TeamListComponent],
            providers: [
                provideHttpClient(),
                provideHttpClientTesting(),
                provideRouter([
                    {
                        path: 'admin/teams/create',
                        component: TeamDetailComponent
                    },
                    {
                        path: 'admin/teams/:team-id',
                        component: TeamDetailComponent
                    }
                ]),
                { provide: TeamsService, useValue: teamsService },
                { provide: NotificationService, useValue: notificationService }
            ]
        }).compileComponents()

        fixture = TestBed.createComponent(TeamListComponent)
        component = fixture.componentInstance
        router = TestBed.inject(Router)
        fixture.detectChanges()
    })

    it('should create', () => {
        expect(component).toBeTruthy()
    })

    it('should load teams into the table', () => {
        const rows = fixture.nativeElement.querySelectorAll('tr')
        expect(rows.length).toBe(mockTeams.length + 1) // header + teams

        // first team row
        expect(rows[1].textContent).toContain(mockTeams[0].name)
        expect(rows[1].textContent).toContain(mockTeams[0].abbreviation)
        expect(rows[1].textContent).toContain(mockTeams[0].location)
        expect(rows[1].textContent).toContain('Jan 1, 2024, 1:00:00 PM')

        // second team row
        expect(rows[2].textContent).toContain(mockTeams[1].name)
        expect(rows[2].textContent).toContain(mockTeams[1].abbreviation)
        expect(rows[2].textContent).toContain(mockTeams[1].location)
        expect(rows[2].textContent).toContain('Jan 2, 2024, 1:00:00 PM')
    })

    it('should display "No teams found" row when dataSource is empty', () => {
        component.dataSource.data = []
        fixture.detectChanges()

        const noDataRow: HTMLElement = fixture.nativeElement.querySelector('tr.mat-row td.mat-cell')
        expect(noDataRow).toBeTruthy()
        expect(noDataRow.textContent).toContain('No teams found')
    })

    it('should show error when loading teams fails', () => {
        const mockError = new Error('Failed')
        teamsService.getTeams.and.returnValue(throwError(() => mockError))

        component.ngOnInit()
        expect(notificationService.showErrorAndLog).toHaveBeenCalledWith(
            'Load Error',
            'Failed to load teams',
            mockError
        )
    })

    it('should navigate to create team page when "Create Team" button is clicked', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')
        fixture.detectChanges()

        const button = fixture.debugElement.query(By.css('.actions button'))
        button.nativeElement.click()

        const call = routerSpy.calls.all()[0].args[0].toString()
        expect(call).toEqual('/admin/teams/create')
    })

    it('should display team table with correct headers and rows', () => {
        const tableRows = fixture.nativeElement.querySelectorAll('tr')
        expect(tableRows.length).toBe(mockTeams.length + 1)

        // Header row
        const headerRow = tableRows[0]
        expect(headerRow.cells[0].innerHTML).toBe('Team Name')
        expect(headerRow.cells[1].innerHTML).toBe('Team Abbreviation')
        expect(headerRow.cells[2].innerHTML).toBe('Team Location')
        expect(headerRow.cells[3].innerHTML).toBe('Created At')
        expect(headerRow.cells[4].innerHTML).toBe('Actions')

        expect(tableRows[1].cells[0].textContent).toBe('Team One')
        expect(tableRows[1].cells[1].textContent).toBe('T1')
        expect(tableRows[1].cells[2].textContent).toBe('City A')
        expect(tableRows[1].cells[3].textContent).toBe('Jan 1, 2024, 1:00:00 PM')
        expect(tableRows[1].cells[4].textContent).toBe('edit')

        expect(tableRows[2].cells[0].textContent).toBe('Team Two')
        expect(tableRows[2].cells[1].textContent).toBe('T2')
        expect(tableRows[2].cells[2].textContent).toBe('City B')
        expect(tableRows[2].cells[3].textContent).toBe('Jan 2, 2024, 1:00:00 PM')
        expect(tableRows[2].cells[4].textContent).toBe('edit')
    })

    it('should navigate correctly when edit button is clicked', () => {
        const routerSpy = spyOn(router, 'navigateByUrl')

        const editButton = fixture.debugElement.query(By.css('[data-testid="edit-button"]'))
        editButton.nativeElement.click()

        const call = routerSpy.calls.all()[0].args[0].toString()
        expect(call).toEqual('/admin/teams/team1')
    })
})
