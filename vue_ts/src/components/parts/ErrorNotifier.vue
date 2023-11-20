<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  visibility: boolean;
}>();
</script>
<template>
  <div
    class="p-4 m-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400"
    role="alert"
    v-if="visibility"
  >
    処理に失敗しました。
  </div>
</template>
