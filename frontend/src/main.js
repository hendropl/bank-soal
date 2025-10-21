import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './style.css' // <-- Pastikan baris ini ada

createApp(App).use(router).mount('#app')