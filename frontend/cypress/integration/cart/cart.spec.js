/// <reference types="cypress" />

describe("cart component test", () => {
    const cart = JSON.stringify([
        {id: 1, name: "Product1", code: "Prd1", quantity: 20},
        {id: 2, name: "Product2", code: "Prd2", quantity: 14},
        {id: 3, name: "Product3", code: "Prd3", quantity: 2}
    ])

    beforeEach(() => {
        cy.visit(Cypress.env("front_url") + "/profile/cart")
    })

    afterEach(() => {
        cy.window().its("localStorage").invoke("removeItem", "cart")
    })

    it("display empty cart", () => {
        cy.get("#cart-empty").should("exist")
        cy.get("#cart-empty").should("have.text", "Cart empty")
    })

    it("display pay button", () => {
        cy.window().its("localStorage").invoke("setItem", "cart", cart)
        cy.window().reload()
        cy.get("#cart-empty").should("not.exist")
        cy.get("#cart-pay").should("exist")
    })

    it("display elements", () => {
        cy.window().its("localStorage").invoke("setItem", "cart", cart)
        cy.window().reload()
        cy.get("#cart-list").should("exist")
        cy.get("#cart-list li").should("have.length", 3)
    })

    it("display first element product name", () => {
        cy.window().its("localStorage").invoke("setItem", "cart", cart)
        cy.window().reload()
        cy.get("#cart-list li").first().should("have.text", "Product1Prd120")
    })

    it("display last element product name", () => {
        cy.window().its("localStorage").invoke("setItem", "cart", cart)
        cy.window().reload()
        cy.get("#cart-list li").last().should("have.text", "Product3Prd32")
    })
})