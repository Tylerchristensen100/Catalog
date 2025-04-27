/* eslint-disable no-undef */
const baseURL = "https://capstone.freethegnomes.org/";
describe('Home page', () => {
    it('passes', () => {
        cy.visit(baseURL)
    })
})

describe('Programs page', () => {
    it('passes', () => {
        cy.visit(baseURL + 'programs')
    })
})

describe('Can view a program', () => {
    it('passes', () => {
        cy.visit(baseURL + 'programs')
        cy.get('main section#programs-list ul')
        .should('have.length.greaterThan', 0)
        cy.get('main section#programs-list ul li').contains('a').click()
    })
})


describe('Can view a programs details', () => {
    it('passes', () => {
        cy.visit(baseURL + 'programs/Nutrition,%20Minor')
        cy.get('main.program-page h1').contains('Nutrition, Minor')
        cy.get('main.program-page section p').contains("minor in nutrition")
    })
})



describe('Courses page', () => {
    it('passes', () => {
        cy.visit(baseURL + 'courses')
    })
})


describe('Can view a course', () => {
    it('passes', () => {
        cy.visit(baseURL + 'courses')
        cy.get('main section#courses-list ul')
        .should('have.length.greaterThan', 0)
    })
})

describe('Can view a course', () => {
    it('passes', () => {
        cy.visit(baseURL + 'courses/CS-1400')
        cy.get('main.course-page h1').contains('CS 1400')
    })
})



describe('Navigation Exists', () => {
    it('passes', () => {
        cy.visit(baseURL)
        cy.get('header#global-header nav').should('exist')
    })
})

describe('Nav > Programs', () => {
    it('passes', () => {
        cy.visit(baseURL)
        cy.get('header#global-header nav').should('exist')
        cy.get('header#global-header nav ul').contains('Programs').click()
    })
})

describe('Nav > Courses', () => {
    it('passes', () => {
        cy.visit(baseURL)
        cy.get('header#global-header nav').should('exist')
        cy.get('header#global-header nav ul').contains('Courses').click()
    })
})

describe('Nav > Colleges & Schools', () => {
    it('passes', () => {
        cy.visit(baseURL)
        cy.get('header#global-header nav').should('exist')
        cy.get('header#global-header nav ul').contains('Colleges & Schools').click()
    })
})
