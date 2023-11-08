// vite.config.js
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: '0.0.0.0',
    port: 8081, // your dev server will run on port 8081
    proxy: {
      // Proxying API requests to backend server running on port 8080
      '/api': config(),
    },
  },
});

function config() {
  if (process.env.mode === 'dockerdev') {
    return {
      target: 'http://se_org_backend:8080',
      changeOrigin: true,
      secure: false,
    }
  }
  return {
    target: 'http://localhost:8080',
    changeOrigin: true,
    secure: false,
  }
}