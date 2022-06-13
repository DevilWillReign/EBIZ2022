/// <reference types="cypress" />

describe("payments component test", () => {
    beforeEach(() => {
        cy.intercept(Cypress.env("api_url") + "/user/payments", {elements: [
            { id: 1, paymenttype: 0, transactionid: 1 },
            { id: 2, paymenttype: 1, transactionid: 2 },
            { id: 3, paymenttype: 0, transactionid: 3 },
            { id: 4, paymenttype: 2, transactionid: 4 },
        ]})
        cy.window().its("localStorage").invoke("setItem", "userinfo", "true")

        cy.visit(Cypress.env("front_url") + "/profile/payments")
    })

    afterEach(() => {
        cy.window().its("localStorage").invoke("removeItem", "userinfo")
    })

    it("display payments", () => {
        cy.get("#payments-list").should("exist")
    })

    it("display payments items", () => {
        cy.get("#payments-list").should("exist")
        cy.get("#payments-list li").should("have.length", 4)
    })

    it("display payments item 1", () => {
        cy.get("#payments-item-1").should("exist")
    })

    it("display payments item text", () => {
        cy.get("#payments-item-1").should("exist")
        cy.get("#payments-item-1-text").should("exist")
        cy.get("#payments-item-1-text").should("have.text", "1 0")
    })

    it("display payments item transaction", () => {
        cy.get("#payments-item-1").should("exist")
        cy.get("#payments-item-1-transaction").should("exist")
        cy.get("#payments-item-1-transaction").should("have.text", "Go to transaction")
        cy.get("#payments-item-1-transaction").should("have.prop", "href", Cypress.env("front_url") + "/transactions/1")
    })
})