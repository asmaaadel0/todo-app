import TheHeader from "../TheHeader.vue";

describe("The Header", () => {
  beforeEach(() => cy.mount(TheHeader));

  it("should have header title", () => {
    cy.get("#header").should("exist");
  });
});
