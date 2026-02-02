import './assets/main.css'

import { createApp } from 'vue';
import PrimeVue from 'primevue/config';
import App from './App.vue';
import Aura from '@primeuix/themes/aura';
import router from './router';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';

const app = createApp(App);
app.use(PrimeVue, {
  theme: {
    preset: Aura
  }
});
app.use(router);
app.use(ToastService)

app.mount('#app');
