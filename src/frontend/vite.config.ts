import { defineConfig } from 'vite'
import { compilerOptions } from './tsconfig.json'
import react from '@vitejs/plugin-react'
import { svgBuilder } from './src/plugins/svgBuilder'
import { BACKENDHOST } from './src/consts'

export default defineConfig({
  plugins: [
    react({
      jsxImportSource: '@emotion/react',
    }),
    svgBuilder('./assets/icons/'),
  ],
  resolve: {
    alias: Object.fromEntries(
      Object.entries(compilerOptions.paths).map(([key, value]) => [
        key.replace('*', ''),
        `/${value[0].replace('*', '')}`,
      ]),
    ),
  },
  server: {
    proxy: {
      '/api': {
        target: `${BACKENDHOST}`,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
      // '/test': {
      //   target: 'https://yapi.hustunique.com/mock',
      //   changeOrigin: true,
      //   rewrite: (path) => path.replace(/^\/test/, ''),
      // },
    },
  },
})
