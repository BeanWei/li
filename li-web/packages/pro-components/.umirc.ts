export default {
  chainWebpack: (memo, { env, webpack, createCSSRule }) => {
    memo.module
      .rule("ts-in-node_modules")
      .include.add(require("path").join(__dirname, "../../packages/"));
  },
};
