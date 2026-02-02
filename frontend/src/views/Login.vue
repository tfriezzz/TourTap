<script setup lang="ts">
import LoginUser from '@/components/LoginUser.vue';
import UserInfo from '@/components/UserInfo.vue';
import store from '@/store';
import type { User } from '@/types/user';
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const onLogout = (user: User) => {
  store.clearUser()
  toast.add({
    severity: 'success',
    summary: `Goodbye, ${user.email}`,
    detail: 'Logout successful',
    life: 3000,
  })
}

</script>

<template>
  <main>
    <h1 class="green">Tour Control</h1>
    <UserInfo v-if="store.state.user" :user="store.state.user" @logout="onLogout" />
    <LoginUser v-else />
  </main>
</template>
