<script setup lang="ts">
import FloatingButton from "@/components/parts/FloatingButton.vue";
import { InjectionKey, provide } from "vue";
import repositoryFactory from "./apis/repositoryFactory";
import { key } from "./main";

// TODO : Vue3を用いた「リアルタイムバリデーション」の実装方法 | https://blog.labol.co.jp/entry/real-time-validation-implementation-method
// TODO : VueのComposition APIでフォームのバリデーション実装 | https://tech.mof-mof.co.jp/blog/validation-vue-composition-api/
// TODO : api request
// TODO : api処理の共通化
/***
 * |
 * | [参考]
 * |
 * _/_/_/_/_/_/_/_/_/_/_/_/_/_/
 * (全体的)
 * _/_/_/_/_/_/_/_/_/_/_/_/_/_/
 * Vue.js+TypeScriptでフロントエンドアプリを構築時のTips | https://rinoguchi.net/2019/05/vue-typescript.html
 * |
 * |
 * _/_/_/_/_/_/_/_/_/_/_/_/_/_/
 * (コンポーネントの抽象化)
 * _/_/_/_/_/_/_/_/_/_/_/_/_/_/
 * 汎用的なフォームコンポーネントを作成したい | https://qiita.com/taka_1156/items/c65700208072b1f30337
 * 使用例 | https://qiita.com/taka_1156/items/c65700208072b1f30337#%E4%BD%BF%E7%94%A8%E6%96%B9%E6%B3%95
 * バリデーションが必要 | https://vuelidate-next.netlify.app/
 * |
 * |
 * _/_/_/_/_/_/_/_/_/_/_/_/_/_/
 * (Ajax処理の抽象化)
 * _/_/_/_/_/_/_/_/_/_/_/_/_/_/
 * `ajax.js`、案件レビューで抽象化できているので、いい感じにtypescriptに変更したら使えるかも、アプローチとしては、以下の記事に似ている
 * JS(javascript)でajaxの通信部分を同期的にして共通化する | https://qiita.com/TetsuTaka/items/489c04b3146b75645f7e
 * 【Nuxt.js】axiosのAPI共通処理をラップし、Repositoryパターンを実現する(TypeScript) | https://sagatto.com/20201229_nuxt_axios_ts_repository
 * 【Nuxt.js】axiosのリクエスト送信/レスポンス受信の共通化を行う | https://qiita.com/shin4488/items/757fbc5ec8fc06c54c8c
 * The Ultimate AJAX Guide For Vue.js Apps | https://vuejsdevelopers.com/2017/08/28/vue-js-ajax-recipes/
 * |
 * |
 * _/_/_/_/_/_/_/_/_/_/_/_/_/_/
 * (アーキテクチャの導入が必要？{今後を考えるとその方がいいし、勉強になる})
 * _/_/_/_/_/_/_/_/_/_/_/_/_/_/
 * Vue 3 + TypeScript ではじめるリポジトリパターン | https://qiita.com/blasg5884/items/394184cbf2f3ea760d4a
 * 【Vue.js】Web API通信のデザインパターン (個人的ベストプラクティス) | https://qiita.com/07JP27/items/0923cbe3b6435c19d761
 * |
 * |
 */
// Symbol --- ユニークのキーを作成できる型
// provide("repository", repositoryFactory);
provide(key, repositoryFactory);
</script>
<template>
  <div class="">
    <!-- <nav>
      <router-link to="/info">Info</router-link>
      <router-link to="/schedule">AddSchedule</router-link>
    </nav> -->
    <router-view />
    <FloatingButton />
  </div>
</template>
