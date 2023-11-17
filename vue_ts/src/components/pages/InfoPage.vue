<!-- Vite + Vue 3 + TypeScript + Tailwind CSS の簡易環境構築 | https://zenn.dev/showy1/articles/c5d1b5d33552be -->
<script setup lang="ts">
import { inject, onMounted, ref } from "vue";
import { key } from "@/provider";

import { PetSummary } from "@/apis/petRepository";
import CardTile from "@/components/parts/CardTile.vue";
import ProfileTile from "@/components/parts/ProfileTile.vue";
import SchedulesTile from "@/components/parts/SchedulesTile.vue";
import SpinnerTile from "@/components/parts/Spinner.vue";

// 文字列だと一致するチェックが緩い（＝取れないかもしれない) *** is possibily undefinedの理由 | const repository = inject<Repositories>("repository");
// Injectをする場所がsetupの箇所です（何かの関数の中ではダメです）
const repository = inject(key);
const state = ref<PetSummary>({
  pet: {
    name: "",
    age: "",
    sex: "",
    nowWeight: "",
    targetWeight: "",
    birthDay: "",
  },
  memo: {
    title: "",
    date: "",
  },
  schedules: [],
});
const isLoading = ref(true);
const fetchPetSummary = async () => {
  // TODO: ローディングアニメーションの制御も、いちいち使う側でやりたくない
  if (!repository) {
    return;
  }
  const { data } = await repository.pet.getPetSummary(2);
  await new Promise((resolve) => setTimeout(resolve, 1000));
  isLoading.value = false;
  state.value = data;
};
onMounted(() => {
  fetchPetSummary();
});
</script>
<template>
  <div>
    <SpinnerTile :is-loading="isLoading" />
    <div
      class="p-3 bg-emerald-500 h-screen"
      :class="isLoading ? 'opacity-20' : ''"
    >
      <ProfileTile
        :name="state.pet.name"
        :age="state.pet.age"
        :sex="state.pet.sex"
        :birthday="state.pet.birthDay"
      />
      <div class="flex justify-between">
        <div class="w-6/12">
          <CardTile
            title="体重"
            :description="state.pet.nowWeight"
            size="max-w-lg"
          />
        </div>
        <div class="w-6/12">
          <CardTile
            title="目標体重"
            :description="state.pet.targetWeight"
            size="max-w-lg"
          />
        </div>
      </div>
      <CardTile
        title="気になるメモ"
        :description="state.memo.title"
        :date="state.memo.date"
      />
      <SchedulesTile :schedules="state.schedules" />
    </div>
  </div>
</template>
