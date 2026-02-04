<script setup lang="ts">
import { ref, computed } from 'vue';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext'
import FloatLabel from 'primevue/floatlabel';
import Button from 'primevue/button';
import axios from 'axios';
import type { User } from '@/types/user';
import store from '@/store';
import { useToast } from 'primevue/usetoast';

interface LoginResponse {
  user: User
  access_token: string
  refresh_token: string
}
//TODO: add expiration time

const username = ref<string>('')
const password = ref<string>('')
const isLoginDisabled = computed(() => !username.value || !password.value)

const errorMessage = ref<string>('')
const toast = useToast()

const handleLogin = async () => {
  try {
    const response = await axios.post<LoginResponse>('/api/login', {
      email: username.value,
      password: password.value,
    })

    store.setUser(response.data.user, response.data.access_token, response.data.refresh_token)
    toast.add({
      severity: 'success',
      summary: `Hello, ${response.data.user.email}`,
      detail: 'Login successful',
      life: 3000,
    })
  } catch (error: any) {
    if (error.response && error.response.status === 404) {
      errorMessage.value = error.response.data.message
    } else {
      errorMessage.value = 'an error occurred'
    }

    toast.add({
      severity: 'error',
      summary: 'Login failed',
      detail: errorMessage.value,
      life: 3000,
    })
  }
}

</script>

<template>
  <div class="flex flex-column row-gap-5">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-user"></i>
      </InputGroupAddon>
      <FloatLabel>
        <InputText id="username" v-model="username" />
        <label for="username">Email</label>
      </FloatLabel>
    </InputGroup>

    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-lock"></i>
      </InputGroupAddon>
      <FloatLabel>
        <InputText type="password" id="password" v-model="password" />
        <label for="password">Password</label>
      </FloatLabel>
    </InputGroup>

    <Button label="Login" :disabled="isLoginDisabled" @click="handleLogin" />
  </div>
</template>
