<template>
  <li>{{ item }}</li>
</template>

<script lang="ts">
export default {
  props: {
    item: {
      type: String,
      required: true,
      default: "item",
    },
  },
};
</script>

<style></style>
