/// <reference types="cypress" />

describe("header component test", () => {
    beforeEach(() => {
        cy.visit(Cypress.env("front_url"))
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
        cy.get("#navbarSupportedContent ul li").should("have.length", 6)
    })

    it("display header names not logged in", () => {
        cy.get("#navbarSupportedContent ul a").eq(4).should("have.text", "Login")
        cy.get("#navbarSupportedContent ul a").eq(4).should("have.prop", "href", Cypress.env("front_url") + "/auth")
        cy.get("#navbarSupportedContent ul a").eq(5).should("have.text", "Register")
        cy.get("#navbarSupportedContent ul a").eq(5).should("have.prop", "href", Cypress.env("front_url") + "/auth/register")
    })

    it("display header names logged in", () => {
        cy.window().its("localStorage").invoke("setItem", "userinfo", "testinfo")
        cy.window().reload()
        cy.get("#navbarSupportedContent ul a").eq(4).should("have.text", "Profile")
        cy.get("#navbarSupportedContent ul a").eq(4).should("have.prop", "href", Cypress.env("front_url") + "/profile")
        cy.get("#navbarSupportedContent ul a").eq(5).should("have.text", "Logout")
        cy.get("#navbarSupportedContent ul a").eq(5).should("have.prop", "href", Cypress.env("front_url") + "/auth/logout")
    })
})