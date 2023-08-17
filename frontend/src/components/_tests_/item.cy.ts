import TaskItem from "../TaskItem.vue";
import type { Task } from "@/App.vue";

describe("task item", () => {
  const task: Task = {
    id: 1,
    title: "Test",
    completed: false,
  };
  const index = 0;
  beforeEach(() =>
    cy.mount(TaskItem, {
      propsData: {
        task,
        index,
      },
    })
  );

  it("should have label for task", () => {
    cy.get("#label-0").should("exist");
  });

  it("should have delete button for task", () => {
    cy.get("#delete-0").should("exist");
  });

  it("should have update button for task", () => {
    cy.get("#update-0").should("exist");
  });
});
