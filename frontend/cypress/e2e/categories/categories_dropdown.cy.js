/// <reference types="cypress" />

describe("categories component test", () => {
    beforeEach(() => {
        cy.intercept(Cypress.env("api_url") + "/categories", {elements: [
            {id: 1, name: "Name1", description: "Description1"},
            {id: 2, name: "Name2", description: "Description2"},
            {id: 3, name: "Name3", description: "Description3"},
            {id: 4, name: "Name4", description: "Description4"}
        ]})
        cy.visit(Cypress.env("front_url") + "/categories")
    })

    it("display categories dropdown container", () => {
        cy.get("#dropdown-Categories").should("exist")
    })

    it("display categories link", () => {
        cy.get("#dropdown-Categories a").should("exist")
        cy.get("#dropdown-Categories a").should("have.text", "Categories")
    })

    it("display categories dropdown", () => {
        cy.get("#dropdown-Categories button").should("exist")
        cy.get("#dropdown-Categories .dropdown-menu").should("exist")
    })

    it("display 4 items", () => {
        cy.get("#dropdown-Categories .dropdown-menu").should("exist")
        cy.get("#dropdown-Categories .dropdown-menu li").should("have.length", 4)
    })

    it("display first item", () => {
        cy.get("#dropdown-Categories .dropdown-menu li").first().should("exist")
        cy.get("#dropdown-Categories .dropdown-menu li").first().should("have.text", "Name1")
    })

    it("should go to categories", () => {
        cy.get("#dropdown-Categories a").should("exist")
        cy.get("#dropdown-Categories a").click()
        cy.location("href").should("contain", "categories")
    })

    it("should go to category", () => {
        cy.intercept(Cypress.env("api_url") + "/categories/1/extended",
        {id: 1, name: "Name1", description: "Description1",
            products: [{id: 1, name: "Name1", code: "Code1", price: 1.05, availability: 40, description: "Description1"}]
        })
        cy.get("#dropdown-Categories .dropdown-menu li").should("have.length", 4)
        cy.get("#dropdown-Categories .dropdown-menu li a").first().click({ force: true })
        cy.location("href").should("contain", "categories/1")
    })
})