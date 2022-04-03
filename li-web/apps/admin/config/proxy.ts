export default {
  dev: {
    "/api": {
      target: "http://localhost:8299/api/",
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/api/, ""),
    },
    "/upload": {
      target: "http://localhost:8299/upload/",
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/upload/, ""),
    },
    "/locales": {
      target: "http://localhost:8299/locales/",
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/locales/, ""),
    },
  },
  test: {
    "/api/": {
      target: "http://localhost:8299",
      changeOrigin: true,
      rewrite: { "^/api": "" },
    },
  },
  pre: {
    "/api/": {
      target: "your pre url",
      changeOrigin: true,
      rewrite: { "^/api": "" },
    },
  },
};
