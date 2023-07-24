<template>
  <the-header></the-header>
  <input-field @add="addItem"></input-field>
  <ul v-for="item in items" :key="item">
    <item-list :item="item"></item-list>
  </ul>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import InputField from "./components/InputField.vue";
import TheHeader from "./components/TheHeader.vue";
import ItemList from "./components/ItemList.vue";

export default defineComponent({
  name: "App",
  components: {
    TheHeader,
    InputField,
    ItemList,
  },
  data() {
    return {
      items: ["test", "fix"],
    };
  },
  methods: {
    addItem(title: string) {
      this.items.push(title);
    },
  },
});
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
