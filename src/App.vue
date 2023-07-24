<template>
  <the-header></the-header>
  <input-field @add="addItem"></input-field>
  <div class="box">
    <ul>
      <item-list
        v-for="item in items"
        :key="item.title"
        :item="item"
      ></item-list>
    </ul>
    <the-footer :count="items.length"></the-footer>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import InputField from "./components/InputField.vue";
import TheHeader from "./components/TheHeader.vue";
import ItemList from "./components/ItemList.vue";
import TheFooter from "./components/TheFooter.vue";

export default defineComponent({
  name: "App",
  components: {
    TheHeader,
    InputField,
    ItemList,
    TheFooter,
  },
  data() {
    return {
      items: [
        {
          title: "hii",
          isChecked: false,
        },
        {
          title: "hello",
          isChecked: true,
        },
      ],
    };
  },
  methods: {
    addItem(title: string) {
      this.items.push({ title: title, isChecked: false });
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
