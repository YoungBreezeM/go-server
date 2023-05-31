import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { createProxyMiddleware } from 'http-proxy-middleware';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],

})
