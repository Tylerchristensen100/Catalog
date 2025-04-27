/* eslint-disable no-undef */
import { mount } from 'cypress/react'
import DataProvider from "../../src/global/contexts/DataContext";

// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add('login', (email, password) => { console.log(email, password) })
//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })

Cypress.Commands.add('mount', (component, options) => {
  // Wrap any parent components needed
  return mount(<DataProvider>{component}</DataProvider>, options);
})

Cypress.Commands.add('login', () => {

  cy.visit('http://localhost:3100/admin/')
  cy.setCookie('access_token', 'LXKeSDprQJgpy8wFdmNzGncrVJO9pKPVgBIduv-3YKaDI09TAfDW9uQS8GIBgPubHHE0vLFzOtZR4lbqyPFfnKaP9_RGXwqGbc-67Xju', {
    domain: 'localhost:3100',
    path: '/',
    secure: true,
    expiry: new Date().getTime() + (60 * 60 * 10), // Optional: Set expiration (in milliseconds)
  })


  cy.contains('logout')
})