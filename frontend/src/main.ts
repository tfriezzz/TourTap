import './assets/main.css'

import { createApp } from 'vue';
import PrimeVue from 'primevue/config';
import App from './App.vue';
import Aura from '@primeuix/themes/aura';
import router from './router';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';

import Dropdown from 'primevue/dropdown';
import Calendar from 'primevue/calendar';
import InputNumber from 'primevue/inputnumber';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Message from 'primevue/message';


const app = createApp(App);
app.use(PrimeVue, {
  theme: {
    preset: Aura
  }
});

app.component('Dropdown', Dropdown);
app.component('Calendar', Calendar);
app.component('InputNumber', InputNumber);
app.component('InputText', InputText);
app.component('Button', Button);
app.component('Message', Message);

app.use(router);
app.use(ToastService)

app.mount('#app');
