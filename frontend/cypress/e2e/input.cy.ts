describe("Test input", () => {
  beforeEach(() => {
    cy.visit("http://localhost:8080/");
  });
  it("should display an error message when title is empty", () => {
    cy.get('input[name="title"]').type("{enter}");
    cy.contains("Please enter the title.").should("be.visible");
  });

  it('should emit "add" event when form is submitted with a title', () => {
    cy.get('input[name="title"]').type("My Task{enter}");
    cy.get('input[name="title"]').type("My Task{enter}");
    cy.get('input[name="title"]').type("My Task{enter}");
  });
});
