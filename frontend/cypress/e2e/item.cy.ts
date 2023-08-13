describe("Todo Item", () => {
  beforeEach(() => {
    cy.visit("/");
  });

  it("should display the task title when not in edit mode", () => {
    cy.get("#label-0").should("have.text", "My Task");
  });

  it("should toggle edit mode when the task title is double-clicked", () => {
    cy.get("#label-1").dblclick();
    cy.get(".item-box form").should("be.visible");
  });

  it("should update the task title when edited and submitted", () => {
    cy.get("#label-1").dblclick();
    cy.get(".item-box input").type("{selectall}New Task{enter}");
    cy.get("#label-1").should("have.text", "New Task");
  });

  it("should update the task title when edited and submitted", () => {
    cy.get("#label-1").dblclick();
    cy.get(".item-box input").type("{selectall}My Task{enter}");
    cy.get("#label-1").should("have.text", "My Task");
  });

  it("should toggle the task completed status when the completed icon is clicked", () => {
    cy.get("#update-0").should("not.have.class", "green");
    cy.get("#update-1").click();
    cy.get("#update-1").should("have.class", "green");
  });

  it("should toggle the task completed status when the completed icon is clicked", () => {
    cy.get("#update-1").click();
    cy.get("#update-1").should("not.have.class", "green");
  });

  it('should emit a "delete" event when the delete icon is clicked', () => {
    cy.get("#delete-2").click();
  });
});
