import { defineConfig } from "vite";
import preact from "@preact/preset-vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [preact()],
  server: {
    port: 3000,
    host: "0.0.0.0",
    cors: true,
  },
  base: "/",
  build: {
    outDir: "./public/assets/",
    rollupOptions: {
      input: {
        main: "./assets/main.tsx",
      },
      output: {
        entryFileNames: "[name].js",
        assetFileNames: "[name][extname]",
        chunkFileNames: "[name][extname]",
      },
    },
  },
});
