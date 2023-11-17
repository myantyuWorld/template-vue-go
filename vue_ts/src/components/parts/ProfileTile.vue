<script setup lang="ts">
import { defineProps } from "vue";
import CardTile from "@/components/parts/CardTile.vue";

defineProps<{
  name?: string;
  age?: string;
  sex?: string;
  birthday?: string;
}>();
</script>

<template>
  <div class="p-4 m-6 lg:px-6 bg-white rounded-lg flex justify-around">
    <div class="w-5/12">
      <img src="@/assets/chai.png" class="w-10/12 rounded-full" />
    </div>
    <div class="w-7/12">
      <div class="flex justify-between">
        <div class="w-6/12">
          <CardTile title="名前" :description="name" size="max-w-lg" />
        </div>
        <div class="w-6/12">
          <CardTile title="年齢" :description="age" size="max-w-lg" />
        </div>
      </div>
      <CardTile title="性別" :description="sex" size="max-w-lg" />
      <CardTile title="誕生日" :description="birthday" size="max-w-lg" />
    </div>
  </div>
</template>
