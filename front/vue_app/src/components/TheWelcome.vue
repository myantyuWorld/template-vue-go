<script setup lang="ts">
import WelcomeItem from "./WelcomeItem.vue";
</script>

<template>
  <div class="m-12">
    <div class="w-screen flex justify-center items-center">
      <h1
        class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white"
      >
        Welcome Sorich!
      </h1>
    </div>
    <div class="w-screen flex justify-center items-center">
      <p
        class="mb-6 text-lg font-normal text-gray-500 lg:text-xl sm:px-16 xl:px-48 dark:text-gray-400"
      >
        株式会社SORICHの研修にようこそ！このページは、自由に編集してもらってもいいです！
      </p>
    </div>
    <hr class="m-8" />
    <h1
      class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-3xl lg:text-4xl dark:text-white"
    >
      フロントエンド側で使用する技術(front/README.mdにも記載)
    </h1>
    <WelcomeItem
      title="Vite"
      description="次世代フロントエンドツール | 従来のビルドツール(Vue CLI等)に比べて、高速で動作するビルドツール"
      url="https://ja.vitejs.dev/"
    />
    <WelcomeItem
      title="vue.js"
      description="今回のToDoアプリ開発に使用するフロントエンドフレームワークです | Web ユーザーインターフェース構築のための、親しみやすく、パフォーマンスと汎用性の高いフレームワーク。"
      url="https://ja.vuejs.org/"
    />
    <WelcomeItem
      title="vue-router"
      description="Vue.jsを利用したSPA構築で、ルーティング制御をするための公式プラグイン"
      url="https://router.vuejs.org/guide/"
    />
    <WelcomeItem
      title="axios"
      description="ブラウザや node.js で動く Promise ベースの HTTP クライアントである｡ REST-API を実行したいときなど､これを使うと実装が簡単にできる"
      url="https://github.com/axios/axios"
    />
    <WelcomeItem
      title="tailwindcss"
      description="TailwindCSSは、ユーティリティクラス(utility class)を自由に組み合わせて活用し、よりオリジナリティの高いデザインのWebサイト及びWebアプリケーションを作成可能なCSSフレームワーク。"
      url="https://tailwindcss.com/"
    />
    <h1
      class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-3xl lg:text-4xl dark:text-white"
    >
      バックエンド側で使用する技術(front/README.mdにも記載)
    </h1>
    <WelcomeItem
      title="echo"
      description="高性能、拡張性、ミニマリストの Go Web フレームワーク"
      url="https://echo.labstack.com/"
    />
    <WelcomeItem
      title="gorm"
      description="デベロッパーフレンドリーを目指した、Go言語のORMライブラリ"
      url="https://gorm.io/ja_JP/docs/index.html"
    />

    <h1
      class="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-3xl lg:text-4xl dark:text-white"
    >
      今回作成するToDoアプリのイメージはこちら
    </h1>
    <WelcomeItem
      title="Figma"
      description="ブラウザ上で簡単にデザインができるツールです。インターフェースのデザインを、場所を選ばずにブラウザさえ起動できれば使えるため、利便性が高くチーム体制での作業にも向いています。エンジニアは基本的にFigmaを使ってデザインするわけでなく、デザイナーの方のデザインをシステムに反映する際に確認しながら使います"
      url="https://www.figma.com/file/rk3KO2peVvQvLWGlpxtySc/TODO-application?type=design&node-id=0%3A1&mode=design&t=i9eiDF8XGt7ZqIHx-1"
    />
  </div>
</template>
