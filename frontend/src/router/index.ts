import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Booking from '../views/Booking.vue'
import Login from '../views/Login.vue'
import Pending from '@/views/Pending.vue'
import store from '@/store'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: { requiresAuth: true }
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
      meta: { requiresAuth: true }
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const isAuthenticated = !!store.state.accessToken

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  }

  else if (to.meta.publicOnly && isAuthenticated) {
    next('/')
  }

  else {
    next()
  }
})

export default router
