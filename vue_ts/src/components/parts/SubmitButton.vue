<script setup lang="ts">
import { defineProps } from "vue";

defineProps<{
  label?: string;
}>();
// https://qiita.com/mml/items/2be6b8df94d932e9a047
</script>

<template>
  <button
    type="submit"
    class="w-4/5 p-3 text-lg font-bold bg-orange-500 text-white rounded-lg"
    @click="$emit('clickEvent')"
  >
    {{ label }}
  </button>
</template>
