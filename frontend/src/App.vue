<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router'
import { onMounted, onBeforeUnmount } from 'vue'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'
import store from '@/store'
import NavigationBar from './components/NavigationBar.vue'
import { startSSE, stopSSE } from '@/sse'


const router = useRouter()
const toast = useToast()

const logout = () => {
  stopSSE()
  store.clearUser()
  router.push('/login')
}

onMounted(() => {
  if (!store.state.accessToken) return
  startSSE((msg) => {
    const currentRoute = router.currentRoute.value

    if (!currentRoute.meta?.requiresAuth) return

    toast.add({
      severity: 'info',
      summary: 'From server:',
      detail: msg,
      // life: 5000
    })
    // window.location.reload()
  })
})

onBeforeUnmount(() => {
  stopSSE()
})


</script>

<template>
  <Toast />
  <NavigationBar v-if="store.state.accessToken" :user="store.state.user" :on-logout="logout" />
  <RouterView />
</template>
