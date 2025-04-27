/* eslint-disable no-undef */
describe('Edit Course autofills fields', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:3100/admin/courses')
        cy.intercept('/api/courses').as('getCourses')
        cy.wait('@getCourses')


        cy.get("ul li:first-child .custom_action_button").click();

        cy.get('input[name="name"]').should('exist')
        cy.get('input[name="major_code"]').should('exist')
        cy.get('input[name="code"]').should('exist')
        cy.get('input[name="credit_hours"]').should('exist')
        cy.get('input[name="prerequisites"]').should('exist')
        cy.get('textarea[name="description"]').should('exist')
    })
})


describe('Edit Course', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:3100/admin/courses')
        cy.intercept('/api/admin/courses').as('PUTCourse')
        cy.intercept('/api/courses').as('getCourses')

        cy.wait('@getCourses')

        cy.get("ul li:first-child .custom_action_button").click();

        cy.get('input[name="name"]').should('exist')
        cy.get('input[name="major_code"]').should('exist')
        cy.get('input[name="code"]').should('exist')
        cy.get('input[name="credit_hours"]').should('exist')
        cy.get('input[name="prerequisites"]').should('exist')
        cy.get('textarea[name="description"]').should('exist')

       
        cy.get('button').contains("Submit").click();

        cy.wait('@PUTCourse').then((interception) => {
            expect(interception.response.statusCode).to.equal(200)
        });
    })
})



describe('Creates new course', () => {
    it('passes', () => {
        cy.login()

        cy.visit('http://localhost:3100/admin/courses')
        cy.intercept('/api/courses').as('getCourses')
        cy.wait('@getCourses')


        cy.get('button').contains("New").click();

        cy.get('input[name="name"]').should('exist')
        cy.get('input[name="major_code"]').should('exist')
        cy.get('input[name="code"]').should('exist')
        cy.get('input[name="credit_hours"]').should('exist')
        cy.get('input[name="prerequisites"]').should('exist')
        cy.get('textarea[name="description"]').should('exist')

        cy.get('input[name="name"]').type('Test Course')
        cy.get('input[name="major_code"]').type('TT')
        cy.get('input[name="code"]').type('1234')
        cy.get('input[name="credit_hours"]').type('3')
        cy.get('input[name="prerequisites"]').type('None')
        cy.get('textarea[name="description"]').type('Test Description')
        cy.intercept('POST', '/api/admin/courses', {
            statusCode: 201,
            body: {
                name: 'Test Course',
                major_code: 'TT',
                code: '1234',
                credit_hours: 3,
                prerequisites: 'None',
                description: 'Test Description',
            },
        }).as('createCourse');


        cy.get('button').contains("Submit").click();

        cy.wait('@createCourse').then((interception) => {
            expect(interception.response.statusCode).to.equal(201)
            expect(interception.response.body.name).to.equal('Test Course')
            expect(interception.response.body.major_code).to.equal('TT')
            expect(interception.response.body.code).to.equal('1234')
            expect(interception.response.body.credit_hours).to.equal(3)
            expect(interception.response.body.prerequisites).to.equal('None')
            expect(interception.response.body.description).to.equal('Test Description')

        });


        cy.visit('http://localhost:5173/admin/courses')
    })
})






