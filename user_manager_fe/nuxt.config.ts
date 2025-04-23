import { defineNuxtConfig } from "nuxt/config";

export default defineNuxtConfig({
  // Configuración del entorno
  typescript: { strict: true },

  // Importación de estilos globales
  css: ['~/assets/scss/main.scss'],

  // Configuración de la app
  app: {
    head: {
      title: 'User Management App',
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Aplicación de gestión de usuarios' },
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/perfil-del-usuario.png' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Ubuntu&display=swap'},
      ]
    }
  },

  components: true,
  compatibilityDate: '2025-04-23'
})