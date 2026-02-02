<script setup lang="ts">
import LoginUser from '@/components/LoginUser.vue';
import UserInfo from '@/components/UserInfo.vue';
import store from '@/store';

</script>

<template>
  <main>
    <h1 class="green">Tour Control</h1>
    <UserInfo v-if="store.state.user" :user="store.state.user" />
    <LoginUser v-else />
  </main>
</template>
