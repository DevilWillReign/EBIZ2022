/// <reference types="cypress" />

describe("home component test", () => {
    beforeEach(() => {
        cy.visit(Cypress.env("front_url"))
    })

    it("display product carousel", () => {
        cy.get("#productCarousel").should("exist")
    })

    it("display product carousel indicators", () => {
        cy.get("#productCarousel .carousel-indicators").should("exist")
        cy.get("#productCarousel .carousel-indicators button").should("have.length", 4)
    })

    it("display product carousel inner", () => {
        cy.get("#productCarousel .carousel-inner").should("exist")
        cy.get("#productCarousel .carousel-inner .carousel-item").should("have.length", 4)
    })

    it("display product carousel inner item", () => {
        cy.get("#productCarousel .carousel-inner").should("exist")
        cy.get("#productCarousel .carousel-inner .carousel-item").first().should("exist")
        cy.get("#productCarousel .carousel-inner .carousel-item img").first().should("have.prop", "alt", "morrison")
    })

    it("display product carousel prev button", () => {
        cy.get("#productCarousel .carousel-control-prev").should("exist")
        cy.get("#productCarousel .carousel-control-prev span").should("have.length", 2)
    })

    it("display product carousel next button", () => {
        cy.get("#productCarousel .carousel-control-next").should("exist")
        cy.get("#productCarousel .carousel-control-next span").should("have.length", 2)
    })
})