<template>
  <div class="item-box">
    <div>
      <form v-if="edit" @submit.prevent="updateTask">
        <input v-model="inputField" />
      </form>
      <label :id="'label-' + index" @dblclick="toggleEdit()" v-else>{{
        task.title
      }}</label>
    </div>
    <div>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
        fill="currentColor"
        class="bi bi-trash3-fill"
        viewBox="0 0 16 16"
        @click="deleteTask()"
        :id="'delete-' + index"
      >
        <path
          d="M11 1.5v1h3.5a.5.5 0 0 1 0 1h-.538l-.853 10.66A2 2 0 0 1 11.115 16h-6.23a2 2 0 0 1-1.994-1.84L2.038 3.5H1.5a.5.5 0 0 1 0-1H5v-1A1.5 1.5 0 0 1 6.5 0h3A1.5 1.5 0 0 1 11 1.5Zm-5 0v1h4v-1a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5ZM4.5 5.029l.5 8.5a.5.5 0 1 0 .998-.06l-.5-8.5a.5.5 0 1 0-.998.06Zm6.53-.528a.5.5 0 0 0-.528.47l-.5 8.5a.5.5 0 0 0 .998.058l.5-8.5a.5.5 0 0 0-.47-.528ZM8 4.5a.5.5 0 0 0-.5.5v8.5a.5.5 0 0 0 1 0V5a.5.5 0 0 0-.5-.5Z"
        />
      </svg>

      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
        fill="currentColor"
        class="bi bi-check2-square"
        viewBox="0 0 16 16"
        :class="task.completed ? 'green' : ''"
        @click="toggleComplete"
        :id="'update-' + index"
      >
        <path
          d="M3 14.5A1.5 1.5 0 0 1 1.5 13V3A1.5 1.5 0 0 1 3 1.5h8a.5.5 0 0 1 0 1H3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V8a.5.5 0 0 1 1 0v5a1.5 1.5 0 0 1-1.5 1.5H3z"
        />
        <path
          d="m8.354 10.354 7-7a.5.5 0 0 0-.708-.708L8 9.293 5.354 6.646a.5.5 0 1 0-.708.708l3 3a.5.5 0 0 0 .708 0z"
        />
      </svg>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    task: {
      type: Object,
      required: true,
      default: () => ({
        title: "task",
        completed: false,
      }),
    },
    index: {
      type: Number,
      required: true,
      default: 0,
    },
  },
  data() {
    return {
      inputField: this.task.title as string,
      inputComplete: this.task.completed as boolean,
      edit: false as boolean,
    };
  },
  methods: {
    deleteTask() {
      this.$emit("delete", this.task.id);
    },
    updateTask() {
      this.$emit("update", this.task.id, this.inputField, this.inputComplete);
      this.edit = false;
    },
    toggleComplete() {
      this.inputComplete = !this.task.completed;
      this.updateTask();
    },
    toggleEdit() {
      this.edit = !this.edit;
    },
  },
});
</script>

<style scoped>
.item-box {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
  padding: 0.5rem;
  background-color: white;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
label {
  word-break: break-all;
  padding: 1rem;
  line-height: 3;
  transition: color 0.4s;
}
input {
  position: relative;
  margin: 0;
  font-size: 1rem;
  font-family: inherit;
  font-weight: inherit;
  line-height: 1.4em;
  border: 0;
  color: inherit;
  padding: 6px;
  border: 1px solid #999;
  box-shadow: inset 0 -1px 5px 0 rgba(0, 0, 0, 0.2);
  box-sizing: border-box;
}
svg {
  padding: 0px;
  margin-left: 3rem;
  color: gray;
  font-weight: bolder;
}

svg:hover {
  color: black;
}

li {
  display: inline;
}
.green {
  color: green;
}
</style>
