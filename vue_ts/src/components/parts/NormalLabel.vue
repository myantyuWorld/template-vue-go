<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  label?: string;
}>();
</script>

<template>
  <div class="p-2">
    <p class="font-normal text-2xl text-gray-700 dark:text-gray-400">
      {{ label }}
    </p>
  </div>
</template>
