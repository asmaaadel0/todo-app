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
          @update="updateTask"
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

declare interface Task {
  id: number;
  title: string;
  completed: boolean;
}
enum FilterStatus {
  ALL = "all",
  Active = "active",
  Completed = "completed",
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
      tasks: [] as Task[],
      showTasks: [] as Task[],
      completedTasks: [] as Task[],
      activeTasks: [] as Task[],
      filter: FilterStatus.ALL as string,
      error: "" as string,
      baseurl: this.$API_BASE_URL,
      editInput: false,
      deletedTask: [] as Task[],
      updatedTask: [] as Task[],
    };
  },
  beforeMount() {
    this.filterChanged(this.filter);
  },
  watch: {
    filter() {
      this.filterChanged(this.filter);
    },
  },
  methods: {
    async getTasks() {
      this.error = "";
      try {
        const response = await fetch(this.baseurl + "/tasks");
        if (response.status != 200) {
          this.error = "Error geting tasks";
          throw new Error("Network response was not ok");
        }

        this.tasks = await response.json();
        if (this.tasks == null) {
          this.tasks = [];
        }
        this.updateFilter();
      } catch (error) {
        if (error instanceof Error) {
          this.error = "Error fetching tasks:" + error.message;
        } else {
          console.log("Unexpected error", error);
        }
      }
    },

    async addTask(title: string) {
      this.error = "";
      const newTask = {
        title: title,
        completed: false,
      };
      const response = await fetch(this.baseurl + "/tasks", {
        method: "POST",
        body: JSON.stringify(newTask),
      });
      if (response.status != 201) {
        this.error = "Error adding task";
      }
      let id: number = await response.json();

      this.tasks.push({ id: id, ...newTask });
      this.activeTasks.push({ id: id, ...newTask });
    },

    async deleteTask(id: number) {
      this.error = "";
      try {
        const response = await fetch(this.baseurl + `/tasks/${id}`, {
          method: "DELETE",
        });

        if (response.status != 200) {
          this.error = "Error deleting task";
          throw new Error("Network response was not ok");
        }
      } catch (error) {
        if (error instanceof Error) {
          this.error = "Error deleting task:" + error.message;
        } else {
          console.log("Unexpected error", error);
        }
      }
      this.deletedTask = this.tasks.filter((t: Task) => {
        return t.id == id;
      });

      let index = this.tasks.indexOf(this.deletedTask[0]);

      if (index !== -1) {
        this.tasks.splice(index, 1);
      }
      index = this.completedTasks.indexOf(this.deletedTask[0]);

      if (index !== -1) {
        this.completedTasks.splice(index, 1);
      }

      index = this.activeTasks.indexOf(this.deletedTask[0]);

      if (index !== -1) {
        this.activeTasks.splice(index, 1);
      }
    },

    async updateTask(id: number, title: string, completed: boolean) {
      this.error = "";
      const updatedTask = { id: id, title: title, completed: completed };
      try {
        const response = await fetch(this.baseurl + `/tasks`, {
          headers: {
            "Content-Type": "application/json",
          },
          method: "PUT",
          body: JSON.stringify(updatedTask),
        });

        if (response.status != 201) {
          this.error = "Error updating task";
          throw new Error("Network response was not ok");
        }
        this.updatedTask = this.tasks.filter((t: Task) => {
          return t.id == id;
        });

        let index = this.tasks.indexOf(this.updatedTask[0]);

        if (index !== -1) {
          this.tasks[index] = { id: id, title: title, completed: completed };
        }

        index = this.completedTasks.indexOf(this.updatedTask[0]);

        if (index !== -1) {
          this.completedTasks[index] = {
            id: id,
            title: title,
            completed: completed,
          };
        }

        index = this.activeTasks.indexOf(this.updatedTask[0]);

        if (index !== -1) {
          this.activeTasks[index] = {
            id: id,
            title: title,
            completed: completed,
          };
        }
      } catch (error) {
        if (error instanceof Error) {
          this.error = "Error updating task:" + error.message;
        } else {
          console.log("Unexpected error", error);
        }
      }
    },

    filterChanged(filter: string) {
      this.getTasks();
      this.filter = filter;
      this.updateFilter();
    },

    updateFilter() {
      this.updateTasks();
      if (this.filter == FilterStatus.ALL) {
        this.showTasks = this.tasks;
      }
      if (this.filter == FilterStatus.Active) {
        this.showTasks = this.activeTasks;
      }
      if (this.filter == FilterStatus.Completed) {
        this.showTasks = this.completedTasks;
      }
    },

    updateTasks() {
      this.completedTasks = this.tasks.filter((t: Task) => {
        return t.completed == true;
      });
      this.activeTasks = this.tasks.filter((t: Task) => {
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
