/// <reference types="cypress" />

describe("products component test", () => {
    beforeEach(() => {
        cy.intercept(Cypress.env("api_url") + "/categories/1/extended",
        {id: 1, name: "Name1", description: "Description1",
            products: [
                {id: 1, name: "Name1", code: "Code1", price: 1.05, availability: 40, description: "Description1"},
                {id: 2, name: "Name2", code: "Code2", price: 100.05, availability: 20, description: "Description2"}
            ]
        })
        cy.visit(Cypress.env("front_url") + "/categories/1")
    })

    it("display category", () => {
        cy.get("#category").should("exist")
        cy.get("#category li").should("have.length", 3)
    })

    it("display name", () => {
        cy.get("#category-name").should("exist")
        cy.get("#category-name").should("have.text", "Category name: Name1")
    })

    it("display description", () => {
        cy.get("#category-description").should("exist")
        cy.get("#category-description").should("have.text", "Category description: Description1")
    })

    it("display products list", () => {
        cy.get("#category-products").should("exist")
        cy.get("#category-products ul").should("exist")
    })

    it("display products", () => {
        cy.get("#category-products ul").should("exist")
        cy.get("#category-products ul li").should("have.length", 2)
    })

    it("display first product", () => {
        cy.get("#category-products ul li").first().should("exist")
        cy.get("#category-products ul li").first().should("have.text", "name: Name1, price: 1.05, avaliability: 40")
    })

    it("display back button", () => {
        cy.get("#category-back").should("exist")
        cy.get("#category-back").should("have.text", "Back to categories list")
    })

    it("should go to categories", () => {
        cy.intercept(Cypress.env("api_url") + "/categories", {elements: [
            {id: 1, name: "Name1", description: "Description1"},
            {id: 2, name: "Name2", description: "Description2"},
            {id: 3, name: "Name3", description: "Description3"},
            {id: 4, name: "Name4", description: "Description4"}
        ]})
        cy.get("#category-back").should("exist")
        cy.get("#category-back").click()
        cy.location("href").should("not.contain", "categories/1")
    })
})