<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  errors: object;
}>();
</script>
<template>
  <div v-for="error of errors" :key="error.$uid">
    <div class="text-red-700 font-bold">{{ error.$message }}</div>
  </div>
</template>
