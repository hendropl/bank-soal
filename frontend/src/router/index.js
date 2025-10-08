import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/Auth/LoginView.vue'
import RegisterView from '../views/Auth/RegisterView.vue'
import ForgotPassword from '../views/Auth/ForgotPassword.vue'
import HomePage from '../views/Auth/HomePage.vue'
import ExamPage from '../views/Auth/ExamPage.vue'
import UpdateProfile from '../views/Auth/UpdateProfile.vue'

const routes = [
  {
    path: '/',
    name: 'HomePage',
    component: HomePage
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  {
    path: '/forgot-password',
    name: 'ForgotPassword',
    component: ForgotPassword
  },
  {
    path: '/ujian',
    name: 'ExamPage',
    component: ExamPage
  },
  {
    path: '/update-profile',
    name: 'UpdateProfile',
    component: UpdateProfile
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router