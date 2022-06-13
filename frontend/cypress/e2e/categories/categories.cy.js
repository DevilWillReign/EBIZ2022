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

    it("display category list", () => {
        cy.get("#category-list").should("exist")
    })

    it("display 4 items", () => {
        cy.get("#category-list .card").should("have.length", 4)
    })

    it("display first category img", () => {
        cy.get("#category-list .card").should("have.length", 4)
        cy.get("#category-list #category-Name1 img").should("exist")
    })

    it("display first category body", () => {
        cy.get("#category-list .card").should("have.length", 4)
        cy.get("#category-list #category-Name1-body").should("exist")
    })

    it("display first category name", () => {
        cy.get("#category-list .card").should("have.length", 4)
        cy.get("#category-list #category-Name1-name").should("exist")
        cy.get("#category-list #category-Name1-name").should("have.text", "Name1")
    })

    it("display first category description", () => {
        cy.get("#category-list .card").should("have.length", 4)
        cy.get("#category-list #category-Name1-description").should("exist")
        cy.get("#category-list #category-Name1-description").should("have.text", "Description1")
    })

    it("display first category link", () => {
        cy.get("#category-list .card").should("have.length", 4)
        cy.get("#category-list #category-Name1-body a").should("exist")
        cy.get("#category-list #category-Name1-body a").should("have.text", "Go to category")
    })

    it("should go to category", () => {
        cy.intercept(Cypress.env("api_url") + "/categories/1",
        {id: 1, name: "Name1", description: "Description1",
            products: [{id: 1, name: "Name1", code: "Code1", price: 1.05, availability: 40, description: "Description1"}]
        })
        cy.get("#category-list .card").should("have.length", 4)
        cy.get("#category-list #category-Name1-body a").click()
        cy.location("href").should("contain", "categories/1")
    })
})