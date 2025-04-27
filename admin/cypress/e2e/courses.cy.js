/* eslint-disable no-undef */

describe('Course details', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:5173/admin/courses')
        cy.intercept('/api/courses').as('getCourses')
        cy.wait('@getCourses')

        cy.get('button[data-accordion-label="Principles of Corporate Tax"]').click();

        cy.get('p').contains('Course Code')
        cy.get('p').contains('Major')
        cy.get('p').contains('Credit Hours')
        cy.get('p').contains('Covers accounting')
    })
})







