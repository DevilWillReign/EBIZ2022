const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
    excludeSpecPattern: ["**/1-getting-started/*.cy.js", "**/2-advanced-examples/*.cy.js"]
  },
});
