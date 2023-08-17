import { createApp } from "vue";
import App from "./App.vue";

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $API_BASE_URL: string;
  }
}

const app = createApp(App);

app.config.globalProperties.$API_BASE_URL = process.env.VUE_APP_API_BASE_URL;

app.mount("#app");
