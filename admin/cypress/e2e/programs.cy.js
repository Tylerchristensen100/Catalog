/* eslint-disable no-undef */

describe('Programs details', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:5173/admin/programs')
        cy.intercept('/api/programs').as('getPrograms')
        cy.wait('@getPrograms')

        cy.get('button[data-accordion-label="Accounting, A.S."]').click();

        cy.get('p').contains('Program Type')
        cy.get('p').contains('Program CIP')
        cy.get('p').contains('Online')
        cy.get('p').contains('School')
        cy.get('p').contains('Grad Level')
    })
})


describe('Edit Program autofills fields', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:5173/admin/programs')
        cy.intercept('/api/programs').as('getPrograms')
        cy.wait('@getPrograms')


        cy.get('button[data-accordion-icon="Accounting, A.S."]').click();

        cy.get('input[name="name"]').should('exist')
        cy.get('input[name="school"]').should('exist')
        cy.get('input[name="program_level"]').should('exist')
        cy.get('input[name="cip"]').should('exist')
        cy.get('input[name="major_code"]').should('exist')
        cy.get('input[name="online"]').should('exist')
        cy.get('textarea[name="description"]').should('exist')
    })
})




describe('Creates new programs', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:5173/admin/programs')
        cy.intercept('/api/programs').as('getPrograms')
        cy.wait('@getPrograms')


        cy.get('button').contains("New").click();

        cy.get('input[name="name"]').should('exist')
        cy.get('input[name="school"]').should('exist')
        cy.get('input[name="program_level"]').should('exist')
        cy.get('input[name="cip"]').should('exist')
        cy.get('input[name="major_code"]').should('exist')
        cy.get('input[name="online"]').should('exist')
        cy.get('textarea[name="description"]').should('exist')

        cy.get('input[name="name"]').type('Test Program')
        cy.get('input[name="school"]').type('Woodbury School of Business', { force: true })
        cy.get('input[name="program_level"]').type('Certificate', { force: true })
        cy.get('input[name="cip"]').type('55.1123')
        cy.get('input[name="major_code"]').type('TT')
        cy.get('input[name="online"]').click({ force: true })
        cy.get('textarea[name="description"]').type('Test Description')

        cy.intercept('POST', 'api/admin/program', {
            statusCode: 201,
            body: {
                name: 'Test Program',
                school: 'Woodbury School of Business',
                program_level: 'Certificate',
                cip: '55.1123',
                major_code: 'TT',
                online: true,
                description: 'Test Description',
            },
        }).as('createProgram');

        cy.get('button').contains("Submit").click();

        cy.wait('@createProgram').then((interception) => {
            expect(interception.response.statusCode).to.equal(201)
            expect(interception.response.body.name).to.equal('Test Program')
            expect(interception.response.body.school).to.equal('Test School')
            expect(interception.response.body.program_level).to.equal('Test Level')
            expect(interception.response.body.cip).to.equal('55.1123')
            expect(interception.response.body.major_code).to.equal('TT')
            expect(interception.response.body.online).to.equal(true)
            expect(interception.response.body.description).to.equal('Test Description')

        });
        cy.visit('http://localhost:5173/admin/programs')
    })
})