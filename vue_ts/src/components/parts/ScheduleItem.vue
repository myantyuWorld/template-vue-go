<script setup lang="ts">
import { defineProps } from "vue";
import NormalLabel from "@/components/parts/NormalLabel.vue";

defineProps<{
  date?: string;
  description?: string;
  location?: string;
}>();
</script>

<template>
  <div>
    <div class="flex justify-between">
      <div class="w-6/12 p-3">
        <NormalLabel :label="date" />
      </div>
      <div class="w-3/12 p-3">
        <NormalLabel :label="description" />
      </div>
      <div class="w-3/12 p-3">
        <NormalLabel :label="location" />
      </div>
    </div>
  </div>
</template>
