<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router'
import Toast from 'primevue/toast'
import store from '@/store'
import NavigationBar from './components/NavigationBar.vue'

const router = useRouter()

const logout = () => {
  store.clearUser()
  router.push('/login')
}



</script>

<template>
  <Toast />
  <NavigationBar v-if="store.state.accessToken" :user="store.state.user" :on-logout="logout" />
  <RouterView />
</template>
