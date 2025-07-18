import tailwindcss from '@tailwindcss/vite'

export default defineNuxtConfig({
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
  modules: ["@nuxt/ui"],
  vite: {
    plugins: [tailwindcss()]
  },
  css: ["assets/css/main.css"],
  colorMode: {
    preference: 'light'
  }
})
