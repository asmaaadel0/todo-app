describe("Test Header", () => {
  it('should display the header with "Todo App" text', () => {
    cy.visit("http://localhost:8080/");
    cy.get(".header h1").should("have.text", "Todo App");
  });
});
