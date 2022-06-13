/// <reference types="cypress" />
import { mount } from "cypress/react"
import { BrowserRouter } from "react-router-dom"
import Link from "../../../src/components/layout/Link"

describe("<Link>", () => {
    beforeEach(() => {
        mount(<BrowserRouter><Link path="/categories" name="Category" /></BrowserRouter>)
    })

    it("exists", () => {
        cy.get(".nav-item").first().should("exist")
    })

    it("exists a", () => {
        cy.get(".nav-item a").first().should("exist")
    })

    it("exists a href", () => {
        cy.get(".nav-item a").first().should("have.prop", "href", "/categories")
    })

    it("exists a text", () => {
        cy.get(".nav-item a").first().should("have.text", "Category")
    })
})