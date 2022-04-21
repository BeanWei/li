import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import svgrPlugin from "@arco-plugins/vite-plugin-svgr";
import vitePluginForArco from "@arco-plugins/vite-react";
import { viteMockServe } from "vite-plugin-mock";
import proxy from "./config/proxy";

const { REACT_APP_ENV = "" } = process.env;

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: [
      { find: "@", replacement: "/src" },
      { find: /^~/, replacement: "" }, // https://github.com/vitejs/vite/issues/2185
    ],
  },
  server: {
    proxy: proxy[REACT_APP_ENV],
  },
  plugins: [
    react(),
    svgrPlugin({
      svgrOptions: {},
    }),
    vitePluginForArco({
      theme: "@arco-themes/react-arco-pro",
      modifyVars: {
        "arcoblue-6": "#165DFF",
      },
    }),
    viteMockServe({
      mockPath: "./mock",
      localEnabled: !!!REACT_APP_ENV,
    }),
  ],
  css: {
    preprocessorOptions: {
      less: {
        javascriptEnabled: true,
      },
    },
  },
});
