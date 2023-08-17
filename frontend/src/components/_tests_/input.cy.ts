import InputField from "../InputField.vue";

describe("Input field", () => {
  beforeEach(() =>
    cy.mount(InputField, {
      data() {
        return {
          error: false,
        };
      },
    })
  );

  it("should have input field", () => {
    cy.get("#title").should("exist");
  });

  it("should have haven't error", () => {
    cy.get(".error").should("not.exist");
  });

  it("should have have error", () => {
    cy.mount(InputField, {
      data() {
        return {
          error: true,
        };
      },
    });
    cy.get(".error").should("exist");
  });
});
