/* eslint-disable no-undef */
describe('Home page', () => {
  it('passes', () => {
    cy.visit('http://localhost:5173/admin/')
  })
})

describe('Programs page', () => {
  it('passes', () => {
    cy.visit('http://localhost:5173/admin/programs')
  })
})


describe('Courses page', () => {
  it('passes', () => {
    cy.visit('http://localhost:5173/admin/courses')
  })
})



