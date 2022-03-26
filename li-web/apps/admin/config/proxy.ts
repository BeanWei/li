export default {
  dev: {
    "/api": {
      target: "http://localhost:8299/api/",
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/api/, ""),
    },
    "/storage": {
      target: "http://localhost:8299/storage/",
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/storage/, ""),
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
