<template>
  <the-header></the-header>
  <input-field @add="addTask"></input-field>
  <div class="box">
    <ul>
      <task-list
        v-for="task in showTasks"
        :key="task.id"
        :task="task"
      ></task-list>
    </ul>
    <the-footer
      :count="showTasks.length"
      @change-filter="filterChanged"
    ></the-footer>
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
      tasks: [
        {
          id: 1,
          title: "hii",
          isChecked: false,
        },
        {
          id: 2,
          title: "hello",
          isChecked: true,
        },
      ],
      showTasks: [{}],
      completedTasks: [{}],
      activeTasks: [{}],
      filter: "all",
    };
  },
  beforeMount() {
    this.showTasks = this.tasks;
    this.updateTasks();
  },
  methods: {
    addTask(title: string) {
      this.tasks.push({ id: Math.random(), title: title, isChecked: false });
      this.updateTasks();
    },
    updateTasks() {
      this.completedTasks = this.tasks.filter((t) => {
        return t.isChecked == true;
      });
      this.activeTasks = this.tasks.filter((t) => {
        return t.isChecked == false;
      });
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
</style>
