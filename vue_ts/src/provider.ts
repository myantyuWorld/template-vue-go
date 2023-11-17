import { App, InjectionKey } from "vue";
import repositoryFactory from "./apis/repositoryFactory";

// ジェネリクスに、使いたい型を定義する
// export interface InjectionKey<T> extends Symbol {
// }
// Symbolの名前、は文字列でOK
export const key: InjectionKey<typeof repositoryFactory> = Symbol("repository");

export const register = (app: App) => app.provide(key, repositoryFactory);
