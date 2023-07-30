<template>
  <div>
    <the-header></the-header>
    <input-field @add="addTask"></input-field>
    <div class="box" v-if="tasks.length != 0 || showTasks.length != 0">
      <ul>
        <task-item
          v-for="(task, index) in showTasks"
          :index="index"
          :key="task.id"
          :task="task"
          @delete="deleteTask"
          @update="UpdateTask"
        ></task-item>
      </ul>
      <the-footer
        :count="showTasks.length"
        @change-filter="filterChanged"
      ></the-footer>
    </div>
    <div v-else>No Tasks yet.</div>
    <div v-if="error != ''" class="error">{{ error }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

import TheHeader from "./components/TheHeader.vue";
import InputField from "./components/InputField.vue";
import TaskItem from "./components/TaskItem.vue";
import TheFooter from "./components/TheFooter.vue";

declare interface Tasks {
  id: number;
  title: string;
  completed: boolean;
}
export default defineComponent({
  name: "App",

  components: {
    TheHeader,
    InputField,
    TaskItem,
    TheFooter,
  },
  data() {
    return {
      tasks: [] as Tasks[],
      showTasks: [] as Tasks[],
      completedTasks: [] as Tasks[],
      activeTasks: [] as Tasks[],
      filter: "all" as string,
      error: "" as string,
      baseurl: "http://localhost:3000" as string,
      editInput: false,
    };
  },
  beforeMount() {
    this.getTasks();
  },
  watch: {
    filter() {
      this.filterChanged(this.filter);
    },
  },
  methods: {
    async getTasks() {
      try {
        const response = await fetch(this.baseurl + "/tasks");
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }

        this.tasks = await response.json();
        if (this.tasks == null) {
          this.tasks = [];
        }
        this.updateFilter();
        this.error = "";
      } catch (error) {
        if (error instanceof Error) {
          this.error = "Error fetching tasks:" + error.message;
        } else {
          console.log("Unexpected error", error);
        }
      }
    },

    async addTask(title: string) {
      const newTask = {
        title: title,
        completed: false,
      };
      const response = await fetch(this.baseurl + "/tasks", {
        method: "POST",
        body: JSON.stringify(newTask),
      });
      this.error = "";
      if (response.status != 201) {
        this.error = "Error adding task";
      }

      this.getTasks();
    },

    async deleteTask(id: number) {
      try {
        const response = await fetch(this.baseurl + `/tasks/${id}`, {
          method: "DELETE",
        });

        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        this.error = "";
      } catch (error) {
        if (error instanceof Error) {
          this.error = "Error deleting task:" + error.message;
        } else {
          console.log("Unexpected error", error);
        }
      }
      this.getTasks();
    },

    async UpdateTask(id: number, title: string, completed: boolean) {
      const updatedTask = { id: id, title: title, completed: completed };
      try {
        const response = await fetch(this.baseurl + `/tasks`, {
          headers: {
            "Content-Type": "application/json",
          },
          method: "PUT",
          body: JSON.stringify(updatedTask),
        });

        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        this.getTasks();
        this.error = "";
      } catch (error) {
        if (error instanceof Error) {
          this.error = "Error updating task:" + error.message;
        } else {
          console.log("Unexpected error", error);
        }
      }
    },

    filterChanged(filter: string) {
      this.filter = filter;
      this.updateFilter();
    },

    updateFilter() {
      this.updateTasks();
      if (this.filter == "all") {
        this.showTasks = this.tasks;
      }
      if (this.filter == "active") {
        this.showTasks = this.activeTasks;
      }
      if (this.filter == "completed") {
        this.showTasks = this.completedTasks;
      }
    },

    updateTasks() {
      this.completedTasks = this.tasks.filter((t) => {
        return t.completed == true;
      });
      this.activeTasks = this.tasks.filter((t) => {
        return t.completed == false;
      });
    },
  },
});
</script>

<style>
.box {
  max-width: 30rem;
  margin: auto;
}
ul {
  padding: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.error {
  color: red;
  margin-top: 30px;
  font-size: 1rem;
  font-weight: bold;
}
</style>
