<script setup lang="ts">
import type { User } from '@/types/user';
import Card from 'primevue/card';
import type { PropType } from 'vue';

const props = defineProps({
  user: {
    type: Object as PropType<User>,
    required: true
  }
})
</script>


<template>
  <Card>
    <template #title>User Profile</template>
    <template #content>
      <div class="flex flex-column row-gap-2">
        <div>
          <i class="pi pi-user"></i>
          <span> {{ user.id }}</span>
        </div>


        <div>
          <i class="pi pi-envelope"></i>
          <span> {{ user.email }}</span>
        </div>
      </div>
    </template>
  </Card>
</template>
