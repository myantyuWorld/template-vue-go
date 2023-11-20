<script setup lang="ts">
import { defineProps } from "vue";
import BaseErrors from "@/components/parts/BaseErrors.vue";
import TitleLabel from "@/components/parts/TitleLabel.vue";

// fileやcolorはUIが違いすぎるので別枠で作成
// もしくはカラーピッカーなど別で作ったほうが見栄えがいいものは除外(場合によっては、time、dateも分ける？)
type InputAttr =
  | "text"
  | "number"
  | "email"
  | "url"
  | "password"
  | "tel"
  | "date"
  | "time"
  | "datetime-local";
defineProps<{
  id: string;
  name: string;
  type: InputAttr;
  value: string;
  placeholder: string;
  errors: object;
  labelTitle: string;
}>();
</script>
<template>
  <div>
    <TitleLabel :label="labelTitle" />
    <input
      :id="id"
      :name="name"
      :type="type"
      :value="value"
      :placeholder="placeholder"
      class="block w-full p-5 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-md focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
    />
    <BaseErrors :errors="errors" />
  </div>
</template>
