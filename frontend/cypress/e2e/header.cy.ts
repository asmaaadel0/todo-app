describe("Test Header", () => {
  it('should display the header with "Todo App" text', () => {
    cy.visit("/");
    cy.get(".header h1").should("have.text", "Todo App");
  });
});
