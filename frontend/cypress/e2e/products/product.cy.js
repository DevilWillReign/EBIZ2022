/// <reference types="cypress" />

describe("products component test", () => {
    beforeEach(() => {
        cy.intercept(Cypress.env("api_url") + "/products/1",
            {id: 1, name: "Name1", code: "Code1", price: 1.05, availability: 40, description: "Description1"}
        )
        cy.visit(Cypress.env("front_url") + "/products/1")
    })

    it("display product", () => {
        cy.get("#product").should("exist")
        cy.get("#product li").should("have.length", 6)
    })

    it("display name", () => {
        cy.get("#product-name").should("exist")
        cy.get("#product-name").should("have.text", "Product name: Name1")
    })

    it("display code", () => {
        cy.get("#product-code").should("exist")
        cy.get("#product-code").should("have.text", "Product code: Code1")
    })

    it("display price", () => {
        cy.get("#product-price").should("exist")
        cy.get("#product-price").should("have.text", "Product price: 1.05")
    })

    it("display availability", () => {
        cy.get("#product-availability").should("exist")
        cy.get("#product-availability").should("have.text", "Product availability: 40")
    })

    it("display description", () => {
        cy.get("#product-description").should("exist")
        cy.get("#product-description").should("have.text", "Product description: Description1")
    })

    it("display quantity", () => {
        cy.get("#product-quantity").should("exist")
        cy.get("#product-quantity-selector").should("exist")
        cy.get("#product-quantity-selector").should("have.value", 1)
    })

    it("display quantity change", () => {
        cy.get("#product-quantity-selector").should("exist")
        cy.get("#product-quantity-selector").should("have.value", 1)
        cy.get("#product-quantity-selector").select(25).should("have.value", 26)
    })

    it("display quantity options", () => {
        cy.get("#product-quantity-selector option").should("have.length", 40)
    })

    it("add product", () => {
        const cart = JSON.stringify([{id: 1, name: "Name1", code: "Code1", price: 1.05, availability: 40, description: "Description1", quantity: 26}])
        cy.get("#product-quantity-selector").select(25).should("have.value", 26)
        cy.get("#add-product").click()
        cy.window().its("localStorage").invoke("getItem", "cart").should("exist")
        cy.window().its("localStorage").invoke("getItem", "cart").should("eq", cart)
    })
})