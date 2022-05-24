/// <reference types="cypress" />

describe("about component test", () => {
    beforeEach(() => {
        cy.visit(Cypress.env.arguments("front_url") + "/about")
    })

    it("display about element", () => {
        cy.get("#about").should("exist")
    })

    it("display about accordion element", () => {
        cy.get("#about").should("exist")
        cy.get("#about .accordion-item").should("have.length", 5)
    })

    it("display about accordion first header element", () => {
        cy.get("#about .accordion-item").first().should("exist")
        cy.get("#about .accordion-item").first().get(".accordion-header").should("exist")
        cy.get("#about .accordion-item").first().get(".accordion-header button").should("exist")
        cy.get("#about .accordion-item").first().get(".accordion-header button").should("have.text", "Accordion Item #1")
    })

    it("display about accordion first collapse element", () => {
        cy.get("#about .accordion-item").first().should("exist")
        cy.get("#about .accordion-item").first().get(".accordion-collapse").should("exist")
        cy.get("#about .accordion-item").first().get(".accordion-collapse .accordion-body").should("exist")
        cy.get("#about .accordion-item").first().get(".accordion-collapse .accordion-body").should("have.text",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit.\
        Etiam efficitur metus sed nisl viverra, et condimentum arcu varius.\
        Maecenas sodales elementum justo. Nulla eleifend convallis ipsum eget aliquam.\
        Cras sem ipsum, consequat in odio at, elementum vulputate dolor.\
        Quisque eu mauris enim.\
        Curabitur porttitor, tellus blandit euismod feugiat, eros quam ultrices lorem, sit amet tristique massa nisl sed diam.\
        Vestibulum sollicitudin ligula sed tellus dictum, nec blandit leo rhoncus.\
        Aliquam lacinia, turpis eget scelerisque pulvinar, felis leo consectetur lacus, vel maximus ipsum urna nec metus.\
        Praesent molestie non diam sit amet accumsan.")
    })
})