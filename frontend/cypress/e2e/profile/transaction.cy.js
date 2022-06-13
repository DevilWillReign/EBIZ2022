/// <reference types="cypress" />

describe("transaction component test", () => {
    const transaciton_data = new Date()

    beforeEach(() => {
        cy.intercept(Cypress.env("api_url") + "/user/transactions/1", {
            id: 1,
            createdat: transaciton_data,
            total: 1000,
            payment: {id: 1, paymenttype: 0},
            quantifiedproducts: [
                {id: 1, name: "Name1", quantity: 1, price: 250},
                {id: 2, name: "Name2", quantity: 2, price: 125},
                {id: 3, name: "Name3", quantity: 5, price: 50},
                {id: 4, name: "Name4", quantity: 10, price: 25},
            ]
        })
        cy.window().its("localStorage").invoke("setItem", "userinfo", "true")

        cy.visit(Cypress.env("front_url") + "/profile/transactions/1")
    })

    afterEach(() => {
        cy.window().its("localStorage").invoke("removeItem", "userinfo")
    })

    it("display transaction", () => {
        cy.get("#transaction-info").should("exist")
    })

    it("display transaction data", () => {
        cy.get("#transaction-date").should("exist")
        cy.get("#transaction-date").should("have.text", "Date: " + transaciton_data)
    })

    it("display transaction total", () => {
        cy.get("#transaction-total").should("exist")
        cy.get("#transaction-total").should("have.text", "Total: 1000")
    })

    it("display transaction payment", () => {
        cy.get("#transaction-payment").should("exist")
        cy.get("#transaction-payment").should("have.text", "Paid")
    })

    it("display transaction products label", () => {
        cy.get("#transaction-products-label").should("exist")
        cy.get("#transaction-products-label").should("have.text", "Products:")
    })

    it("display transaction products", () => {
        cy.get("#transaction-products").should("exist")
        cy.get("#transaction-products li").should("have.length", 4)
    })

    it("display transaction first product", () => {
        cy.get("#transaction-product-1").should("exist")
        cy.get("#transaction-product-1 a").should("have.text", "Name: Name1, Quantity: 1, Price: 250")
    })
})