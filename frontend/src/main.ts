import './assets/main.css'

import { createApp } from 'vue';
import PrimeVue from 'primevue/config';
import App from './App.vue';
import Aura from '@primeuix/themes/aura';
import router from './router';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';

import Select from 'primevue/dropdown';
import DatePicker from 'primevue/datepicker';
import InputNumber from 'primevue/inputnumber';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Message from 'primevue/message';
import ProgressSpinner from 'primevue/progressspinner';


const app = createApp(App);
app.use(PrimeVue, {
  theme: {
    preset: Aura
  }
});

app.component('Select', Select);
app.component('DatePicker', DatePicker);
app.component('InputNumber', InputNumber);
app.component('InputText', InputText);
app.component('Button', Button);
app.component('Message', Message);
app.component('ProgressSpinner', ProgressSpinner);

app.use(router);
app.use(ToastService)

app.mount('#app');
