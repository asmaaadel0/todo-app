<template>
  <the-header></the-header>
  <input-field></input-field>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import InputField from "./components/InputField.vue";
import TheHeader from "./components/TheHeader.vue";

@Options({
  components: {
    TheHeader,
    InputField,
  },
})
export default class App extends Vue {}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
