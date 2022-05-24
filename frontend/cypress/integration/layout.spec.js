/// <reference types="cypress" />

describe("layout component test", () => {
    beforeEach(() => {
        cy.visit(Cypress.env.arguments("front_url"))
    })

    it("check header element", () => {
        cy.get(".navbar").should("exist")
    })

    it("check container element", () => {
        cy.get("#container").should("exist")
    })

    it("check footer element", () => {
        cy.get("footer").should("exist")
    })
})