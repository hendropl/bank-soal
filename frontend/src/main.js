import { createApp } from 'vue'
import App from './App.vue'
import router from './router' // tambahkan ini
import './style.css'

const app = createApp(App)
app.use(router) // gunakan router
app.mount('#app')
