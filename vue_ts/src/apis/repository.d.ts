import * as repositoryFactory from "./repositoryFactory";

declare module "@vue/runtime-core" {
  export interface ComponentCustomProperties {
    $repository: repositoryFactory.Repositories;
  }
}
