import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "path";
import { config } from "dotenv";

config({ path: resolve(__dirname, "../.env") });

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src"),
    },
  },

  define: {
    "process.env": process.env,
  },
});
