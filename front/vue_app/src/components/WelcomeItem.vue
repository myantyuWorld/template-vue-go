<script setup lang="ts">
defineProps<{
  title: string;
  description: string;
  url: string;
}>();

</script>
<template>
  <div class="m-8 items-center">
    <a
      :href="url"
      target="_blank"
      class="block max-w-full p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
    >
      <h1
        class="mb-2 text-3xl font-bold tracking-tight text-gray-900 dark:text-white"
      >
        {{ title }}
      </h1>
      <p class="font-normal text-gray-700 dark:text-gray-400">
        {{ description }}
      </p>
    </a>
  </div>
</template>