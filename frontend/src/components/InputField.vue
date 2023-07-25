<template>
  <div>
    <form @submit.prevent="submitForm">
      <input type="title" name="title" id="title" v-model="title" />
    </form>
    <div v-if="error" class="error">Please enter the title.</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  data() {
    return {
      title: "" as string,
      error: false as boolean,
    };
  },
  emits: ["add"],
  methods: {
    submitForm() {
      if (this.title == "") {
        this.error = true;
        return;
      }
      this.error = false;
      this.$emit("add", this.title);
      this.title = "";
    },
  },
});
</script>

<style scoped>
form {
  max-width: 420px;
  margin: 30px auto;
  background: white;
  text-align: left;
  padding: 40px;
  border-radius: 10px;
}

input {
  display: block;
  padding: 10px 6px;
  box-sizing: border-box;
  width: 100%;
  border: none;
  border-bottom: 1px solid #ddd;
  color: #555;
  font-size: 2rem;
}
input:focus {
  outline: none;
  border-bottom: 3px solid #d2eed3;
}

.error {
  color: red;
  margin-top: 30px;
  font-size: 1rem;
  font-weight: bold;
}
</style>
