describe("Test Footer", () => {
  beforeEach(() => {
    cy.visit("/");
  });

  it("add task", () => {
    cy.get('input[name="title"]').type("My Task{enter}");
  });

  it("should display the correct count of items left", () => {
    cy.get(".todo-count").should("contain.text", "items left"); // Assuming the default count is 0
  });

  it("should change the filter when All button is clicked", () => {
    cy.get('button:contains("All")').click();
    cy.get("button.clicked").should("contain.text", "All");
  });

  it("should change the filter when Active button is clicked", () => {
    cy.get('button:contains("Active")').click();
    cy.get("button.clicked").should("contain.text", "Active");
  });

  it("should change the filter when Completed button is clicked", () => {
    cy.get('button:contains("Completed")').click();
    cy.get("button.clicked").should("contain.text", "Completed");
  });

  it('should emit the "change-filter" event when a button is clicked', () => {
    cy.get('button:contains("Active")').click();
  });
});
