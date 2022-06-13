/// <reference types="cypress" />
import { mount } from "cypress/react"
import { BrowserRouter } from "react-router-dom"
import Dropdown from "../../../src/components/layout/Dropdown"

describe("<Link>", () => {
    const categories = [
        {id: 1, name: "Name1", description: "Description1"},
        {id: 2, name: "Name2", description: "Description2"},
        {id: 3, name: "Name3", description: "Description3"},
        {id: 4, name: "Name4", description: "Description4"}
    ]

    beforeEach(() => {
        mount(<BrowserRouter><Dropdown path="/categories" name="Category" elements={categories} /></BrowserRouter>)
    })

    it("exists categories dropdown container", () => {
        cy.get("#dropdown-Categories").should("exist")
    })

    it("exists categories link", () => {
        cy.get("#dropdown-Categories a").should("exist")
        cy.get("#dropdown-Categories a").should("have.text", "Categories")
    })

    it("exists categories dropdown", () => {
        cy.get("#dropdown-Categories button").should("exist")
        cy.get("#dropdown-Categories .dropdown-menu").should("exist")
    })

    it("exists 4 items", () => {
        cy.get("#dropdown-Categories .dropdown-menu").should("exist")
        cy.get("#dropdown-Categories .dropdown-menu li").should("have.length", 4)
    })

    it("exists first item", () => {
        cy.get("#dropdown-Categories .dropdown-menu li").first().should("exist")
        cy.get("#dropdown-Categories .dropdown-menu li").first().should("have.text", "Name1")
    })

    it("exists go to categories", () => {
        cy.get("#dropdown-Categories a").should("exist")
        cy.get("#dropdown-Categories a").click()
        cy.location("href").should("contain", "categories")
    })
})