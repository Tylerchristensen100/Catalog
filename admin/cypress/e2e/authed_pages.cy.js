/* eslint-disable no-undef */
describe('Log In', () => {
    it('passes', () => {
        cy.login()
    })
})



describe('Programs', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:5173/admin/programs')
        cy.intercept('/api/programs').as('getPrograms')

        cy.get('main ul').contains('p').click()

    })
})


describe('Courses', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:5173/admin/courses')
        cy.intercept('/api/courses').as('getCourses')
        cy.intercept('/api/programs').as('getPrograms')
        cy.wait('@getCourses')

        cy.get('main ul').contains('c')

    })
})


