import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';

// https://vite.dev/config/
export default defineConfig({
  plugins: [tailwindcss(), svelte()],
  outDir: "dist",
  server: {
    port: 5173,
    cors: true,
  }
});
