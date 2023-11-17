<script setup lang="ts">
import { defineProps } from "vue";
import ScheduleItem from "@/components/parts/ScheduleItem.vue";
import TitleLabel from "@/components/parts/TitleLabel.vue";

defineProps<{
  schedules: object;
}>();
</script>

<template>
  <div class="p-2 lg:px-6 bg-blue rounded-lg">
    <div
      class="block p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
    >
      <TitleLabel label="次の予定" />
      <ScheduleItem
        v-for="item of schedules"
        :date="item.date"
        :description="item.title"
        :location="item.location"
        :key="item.title"
      />
    </div>
  </div>
</template>
