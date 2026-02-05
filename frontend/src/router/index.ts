import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Booking from '../views/Booking.vue'
import Login from '../views/Login.vue'
import Pending from '@/views/Pending.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/booking',
      name: 'booking',
      component: Booking,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/pending',
      name: 'pending',
      component: Pending,
    },
  ],
})

export default router
