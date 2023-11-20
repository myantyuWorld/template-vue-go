<script setup lang="ts">
import { inject, reactive, ref } from "vue";
import { useVuelidate } from "@vuelidate/core";
import { required } from "@/utils/i18n-validators.ts";
import { key } from "@/provider";

import SubmitButton from "@/components/parts/SubmitButton.vue";
import BaseInput from "@/components/parts/BaseInput.vue";
import HorizontalLine from "@/components/parts/HorizontalLine.vue";
import SpinnerTile from "@/components/parts/Spinner.vue";
import { Schedule } from "@/apis/petRepository";
import ErrorNotifier from "../parts/ErrorNotifier.vue";

const repository = inject(key);
// Vue3+Vuelidateでexternal validationsを試す | https://zenn.dev/kakkoyakakko/articles/ddac0fb3c4c642
const formData = reactive({
  title: "",
  date: "",
  location: "",
});
let isLoading = ref(false);
let isError = ref(false);
const rules = {
  title: { required },
  date: { required },
  location: { required },
};
const v$ = useVuelidate(rules, formData);

const clickEvent = async () => {
  isError.value = false;
  console.log("submit", formData);
  v$.value.$validate();
  if (!v$.value.$invalid) {
    console.log("バリデーションパス、リクエスト送信");
    // TODO: ローディングアニメーションの制御も、いちいち使う側でやりたくない | おそらくこの処理をうまくレイヤー化できれば、HTTPリクエスト・レスポンスのテスト化が可能？
    isLoading.value = true;
    const schedule: Schedule = {
      title: formData.title,
      date: formData.date,
      location: formData.location,
    };
    if (!repository) {
      return;
    }
    const { data } = await repository?.schedule.create(1, schedule);
    await new Promise((resolve) => setTimeout(resolve, 1000));
    isLoading.value = false;
    if (data === undefined) {
      isError.value = true;
    }
  }
};
</script>
<template>
  <div>
    <SpinnerTile :is-loading="isLoading" />
    <div
      class="p-8 bg-emerald-500 h-screen"
      :class="isLoading ? 'opacity-20' : ''"
    >
      <div class="p-4 bg-white rounded-lg grid shadow-xl">
        <div class="text-2xl font-bold">
          <h1>予定を追加</h1>
        </div>
        <ErrorNotifier :visibility="isError" />
        <BaseInput
          id="title"
          name="title"
          type="text"
          placeholder="例）トリミング"
          v-model:value="formData.title"
          :errors="v$.title.$errors"
          label-title="タイトル"
        />
        <HorizontalLine />
        <BaseInput
          id="date"
          name="date"
          type="datetime-local"
          placeholder=""
          v-model:value="formData.date"
          :errors="v$.date.$errors"
          label-title="日時"
        />
        <HorizontalLine />
        <BaseInput
          id="location"
          name="location"
          type="text"
          placeholder="例）ぺテモ立川店"
          v-model:value="formData.location"
          :errors="v$.location.$errors"
          label-title="場所"
        />
        <HorizontalLine />
        <div class="p-8 text-center">
          <SubmitButton label="作成する" @click="clickEvent" />
        </div>
      </div>
    </div>
  </div>
</template>
