/// <reference types="cypress" />

describe("header component test", () => {
    beforeEach(() => {
        cy.visit(Cypress.env.arguments("front_url"))
    })

    it("display navbar", () => {
        cy.get(".navbar").should("exist")
    })

    it("display header container", () => {
        cy.get("#header-container").should("exist")
    })

    it("display header content", () => {
        cy.get("#navbarSupportedContent").should("exist")
        cy.get("#navbarSupportedContent ul").should("exist")
        cy.get("#navbarSupportedContent ul").should("have.length", 5)
    })

    it("display header names not logged in", () => {
        cy.get("#navbarSupportedContent ul").children()[3].get("a").should("have.text", "Login")
        cy.get("#navbarSupportedContent ul").children()[3].get("a").should("have.property", "href", "/auth")
        cy.get("#navbarSupportedContent ul").children()[4].get("a").should("have.text", "Register")
        cy.get("#navbarSupportedContent ul").children()[4].get("a").should("have.property", "href", "/auth/register")
    })

    it("display header names logged in", () => {
        cy.window().its("sessionStorage").invoke("setItem", "userinfo", "testinfo")
        cy.get("#navbarSupportedContent ul").children()[3].get("a").should("have.text", "Profile")
        cy.get("#navbarSupportedContent ul").children()[3].get("a").should("have.property", "href", "/profile")
        cy.get("#navbarSupportedContent ul").children()[4].get("a").should("have.text", "Logout")
        cy.get("#navbarSupportedContent ul").children()[4].get("a").should("have.property", "href", "/auth/logout")
    })
})