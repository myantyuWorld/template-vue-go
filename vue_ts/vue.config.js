module.exports = {
  devServer: {
    client: {
      // Docker環境でVUE CLIのホットリロードが効かない | https://qiita.com/double_lemon/items/f858698871399040ef59
      // "webpack-dev-server": "^4.7.3",// webpack4系からは、hostオプションが消えているよう
      webSocketURL: "ws://0.0.0.0:3001/ws",
    },
  },
};
