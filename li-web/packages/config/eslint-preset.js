module.exports = {
  root: true,
  extends: ["prettier"],
  settings: {
      next: {
      rootDir: ["apps/*/", "packages/*/"],
      },
  },
  rules: {
      "react/jsx-key": "off",
  },
};
