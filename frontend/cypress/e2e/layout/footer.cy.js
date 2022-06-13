/// <reference types="cypress" />

describe("footer component test", () => {
    beforeEach(() => {
        cy.visit(Cypress.env("front_url"))
    })

    it("display footer", () => {
        cy.get("#footer").should("exist")
    })

    it("display links list", () => {
        cy.get("#footer-links").should("exist")
    })

    it("display links header", () => {
        cy.get("#footer-links h5").should("exist")
        cy.get("#footer-links h5").should("have.text", "Links")
    })

    it("display links links", () => {
        cy.get("#footer-links ul").should("exist")
        cy.get("#footer-links ul li").should("have.length", 1)
    })

    it("display links link about", () => {
        cy.get("#footer-links ul li a").should("exist")
        cy.get("#footer-links ul li a").should("have.text", "About")
    })

    it("links link about click", () => {
        cy.get("#footer-links ul li a").click()
        cy.location("href").should("contain", "about")
    })

    it("display copyright", () => {
        cy.get("#footer-copyright").should("exist")
        cy.get("#footer-copyright").should("have.text", "© 2022 Copyright")
    })
})