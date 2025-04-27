import { defineConfig } from "cypress";
import customViteConfig from "./vite.config.js";

export default defineConfig({
  component: {
    devServer: {
      framework: "react",
      bundler: "vite",
      viteConfig: customViteConfig,

    },
    supportFile: './cypress/support/e2e.js',
  },

  e2e: {
    supportFile: './cypress/support/e2e.js',
  },
});
