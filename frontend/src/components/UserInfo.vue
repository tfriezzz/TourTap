<script setup lang="ts">
import type { User } from '@/types/user';
import Card from 'primevue/card';
import Divider from 'primevue/divider';
// import Button from 'primevue/button'
import type { PropType } from 'vue';

const props = defineProps({
  user: {
    type: Object as PropType<User>,
    required: true
  }
})

const emit = defineEmits<{
  (e: 'logout', user: User): void
}>()

const onLogout = () => emit('logout', props.user)

</script>


<template>
  <h1 class="flex justify-content-around flex-wrap">User Menu</h1>
  <Card>
    <template #title>User Profile</template>
    <template #content>
      <div class="flex flex-column row-gap-2">
        <div>
          <i class="pi pi-user"></i>
          <span class="m-2"> {{ user.name }}</span>
        </div>


        <div>
          <i class="pi pi-envelope"></i>
          <span class="m-2">
            <a :href="`mailto:${user.email}`">{{ user.email }} </a>
          </span>
        </div>
      </div>
      <Divider />
      <Button label="Logout" icon="pi pi-sign-out" @click="onLogout" />
    </template>
  </Card>
</template>
