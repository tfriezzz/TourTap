<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { computed } from 'vue'
import Menubar from 'primevue/menubar'
import Button from 'primevue/button'
import store from '@/store'

const router = useRouter()
const route = useRoute()

const showNavbar = computed(() => {
  return !route.meta.hideNavbar
})

const logout = () => {
  store.clearUser()
  router.push('/login')
}

const menuItems = computed(() => [
  {
    label: 'User Menu',
    icon: 'pi pi-user',
    command: () => router.push('/login')
  },
  {
    label: 'Bookings',
    icon: 'pi pi-calendar',
    command: () => router.push('/')
  },
  {
    label: 'Pending',
    icon: 'pi pi-clock',
    command: () => router.push('/pending')
  },
])
</script>

<template>
  <Menubar v-if="showNavbar" :model="menuItems" class="shadow-2 mb-4">
    <template #end>
      <div class="flex align-items-center gap-3">
        <span class="font-semibold">
          {{ store.state.user?.name }}
        </span>
        <Button label="Logout" icon="pi pi-sign-out" outlined @click="logout" />
      </div>
    </template>
  </Menubar>
</template>
