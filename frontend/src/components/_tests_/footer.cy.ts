import TheFooter from "../TheFooter.vue";

describe("The Footer", () => {
  beforeEach(() => cy.mount(TheFooter));

  it("should have counter", () => {
    cy.get(".todo-count").should("exist");
  });

  it("should have all button", () => {
    cy.get("#all").should("exist");
  });

  it("should have active button", () => {
    cy.get("#active").should("exist");
  });

  it("should have completed button", () => {
    cy.get("#completed").should("exist");
  });
});
