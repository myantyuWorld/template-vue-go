module.exports = {
  preset: "@vue/cli-plugin-unit-jest/presets/typescript",
  transformIgnorePatterns: ["/node_modules"],
  // transform: {
  //   "node_modules/three/examples/.+.(j|t)sx?$": "ts-jest",
  // },
  moduleNameMapper: {
    "^axios$": "axios/dist/node/axios.cjs", // https://chaika.hatenablog.com/entry/2023/08/18/083000
  },
  testEnvironment: "node", // or jest-environment-jsdom
};
