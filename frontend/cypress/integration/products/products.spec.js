/// <reference types="cypress" />

describe("products component test", () => {
    beforeEach(() => {
        cy.intercept(Cypress.env("api_url") + "/products", [
            {id: 1, name: "Name1", code: "Code1"},
            {id: 2, name: "Name2", code: "Code2"},
            {id: 3, name: "Name3", code: "Code3"},
            {id: 4, name: "Name4", code: "Code4"}
        ])
        cy.visit(Cypress.env("front_url") + "/products")
    })

    it("display product list", () => {
        cy.get("#product-list").should("exist")
    })

    it("display 4 items", () => {
        cy.get("#product-list li").should("have.length", 4)
    })

    it("display first product info check", () => {
        cy.get("#product-list li").should("have.length", 4)
        cy.get("#product-list li").first().should("have.text", "Code1 Name1")
    })

    it("display last product info check", () => {
        cy.get("#product-list li").should("have.length", 4)
        cy.get("#product-list li").last().should("have.text", "Code4 Name4")
    })

    it("should go to product", () => {
        cy.intercept(Cypress.env("api_url") + "/products/1",
            {id: 1, name: "Name1", code: "Code1", price: 1.05, availability: 40, description: "Description1"}
        )
        cy.get("#product-list li").should("have.length", 4)
        cy.get("#product-list li").first().should("have.text", "Code1 Name1")
        cy.get("#product-list li a").first().click()
        cy.location("href").should("contain", "products/1")
    })
})