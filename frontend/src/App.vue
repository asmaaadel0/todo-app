<template>
  <div>
    <the-header></the-header>
    <input-field @add="addTask"></input-field>
    <div class="box" v-if="tasks.length != 0">
      <ul>
        <task-list
          v-for="task in tasks"
          :key="task.id"
          :task="task"
          @delete="deleteTask"
          @update="UpdateTask"
        ></task-list>
      </ul>
      <the-footer
        :count="tasks.length"
        @change-filter="filterChanged"
      ></the-footer>
    </div>
    <div v-else>No Tasks yet.</div>
    <div v-if="error != ''" class="error">{{ error }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import InputField from "./components/InputField.vue";
import TheHeader from "./components/TheHeader.vue";
import TaskList from "./components/TaskList.vue";
import TheFooter from "./components/TheFooter.vue";

export default defineComponent({
  name: "App",
  components: {
    TheHeader,
    InputField,
    TaskList,
    TheFooter,
  },
  data() {
    return {
      tasks: [{}],
      showTasks: [{}],
      completedTasks: [{}],
      activeTasks: [{}],
      filter: "all",
      error: "",
      baseurl: "http://localhost:3000",
    };
  },
  beforeMount() {
    this.getTasks();
    this.showTasks = this.tasks;
    this.updateTasks();
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
      } catch (error) {
        this.error = "Error fetching tasks:" + error.message;
        console.error("Error fetching data:", error.message);
      }
    },

    async addTask(title: string) {
      const response = await fetch(this.baseurl + "/tasks", {
        method: "POST",
        body: JSON.stringify({
          title: title,
          isChecked: false,
        }),
      });
      if (response.status != 200) {
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
        this.getTasks();
      } catch (error) {
        this.error = "Error deleting task:" + error.message;
      }
    },

    async UpdateTask(id: number, title: string, isChecked: boolean) {
      const updatedTask = { title: title, isChecked: isChecked };
      try {
        const response = await fetch(this.baseurl + `/tasks/${id}`, {
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
      } catch (error) {
        this.error = "Error updating task:" + error.message;
      }
    },

    filterChanged(filter: string) {
      this.filter = filter;
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
        return t.isChecked == true;
      });
      this.activeTasks = this.tasks.filter((t) => {
        return t.isChecked == false;
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
