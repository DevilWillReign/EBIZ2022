/// <reference types="cypress" />

describe("profile component test", () => {
    beforeEach(() => {
        cy.intercept(Cypress.env("api_url") + "/user/me", {
            name: "Name1",
            email: "Name1@example.com"
        })
        cy.window().its("localStorage").invoke("setItem", "userinfo", "true")

        cy.visit(Cypress.env("front_url") + "/profile")
    })

    afterEach(() => {
        cy.window().its("localStorage").invoke("removeItem", "userinfo")
    })

    it("display user", () => {
        cy.get("#user-info").should("exist")
    })

    it("display user name", () => {
        cy.get("#user-name").should("exist")
        cy.get("#user-name").should("have.text", "Name: Name1")
    })

    it("display user email", () => {
        cy.get("#user-email").should("exist")
        cy.get("#user-email").should("have.text", "Email: Name1@example.com")
    })

    it("display user cart button", () => {
        cy.get("#user-cart").should("exist")
        cy.get("#user-cart").should("have.text", "Cart")
    })

    it("display user payments button", () => {
        cy.get("#user-payments").should("exist")
        cy.get("#user-payments").should("have.text", "Payments")
    })

    it("display user transactions button", () => {
        cy.get("#user-transactions").should("exist")
        cy.get("#user-transactions").should("have.text", "Transactions")
    })
})