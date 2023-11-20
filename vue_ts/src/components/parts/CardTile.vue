<script setup lang="ts">
import { defineProps } from "vue";
import NormalLabel from "@/components/parts/NormalLabel.vue";
import TitleLabel from "@/components/parts/TitleLabel.vue";

defineProps<{
  title?: string;
  description?: string;
  date?: string;
  size?: string;
}>();
</script>

<template>
  <div class="p-3 lg:px-6 bg-blue rounded-lg">
    <div
      class="block p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
      :class="size"
    >
      <TitleLabel :label="title" />
      <NormalLabel :label="description" />
      <div class="flex justify-end">
        <TitleLabel label="最終更新日" v-if="date != null" />
        <NormalLabel :label="date" v-if="date != null" />
      </div>
    </div>
  </div>
</template>
