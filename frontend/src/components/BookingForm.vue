<script setup lang="ts">
import { z } from 'zod';
import { zodResolver } from '@primevue/forms/resolvers/zod';
import { Form } from '@primevue/forms';
import InputText from 'primevue/inputtext';  // Add this
import Message from 'primevue/message';      // Add this
import Button from 'primevue/button';        // Add this
import axios from 'axios';
import { ref, computed, onMounted } from 'vue';
import { useToast } from 'primevue';

interface BackendResponse {
  message: string
}

interface Tour {
  id: number
  name: string
  created_at: string
  updated_at: string
  base_price: string
}

const toast = useToast()

// const emailSchema = z.object({
//   email: z.string().email({ message: 'Please enter a valid email address' }),
// });


const getAllTours = () => axios.get('/api/tours')
const tours = ref<Tour[]>([])
const toursError = ref('')

onMounted(async () => {
  toursError.value = ''
  try {
    const response = await getAllTours()
    tours.value = response.data

    const tourList = tours.value
      .map(tour => `• ${tour.name} - JPY${tour.base_price}`)
      .join('\n')

    toast.add({
      severity: 'success',
      summary: `Fetched ${tours.value.length} tours`,
      detail: tourList,
      life: 30000,
    })
  } catch (err) {
    toursError.value = 'Could not load tours. Make sure backend is running! Refresh to retry.'
  }
})

// const isSubmitDisabled = computed(() => !email.value)

const errorMessage = ref<string>('')

// type EmailFormData = z.infer<typeof emailSchema>;

// const resolver = zodResolver(emailSchema);


const resolver = ref(zodResolver(
  z.object({
    email: z.string().min(1, { message: 'Email is required via Zod.' })
  })
));

const initialValues = { email: '' };
const email = ref<string>('')

const onSubmit = async () => {
  try {
    const response = await axios.post<BackendResponse>('/api/test', {
      email: email.value
    })
    toast.add({
      severity: 'success',
      summary: `${response.data.message}`,
      detail: 'We\'ll contact you asap',
      life: 3000,
    })

  } catch (error: any) {
    if (error.response) {
      errorMessage.value = error.response.data.message
      toast.add({
        severity: 'error',
        summary: 'Submission failed',
        detail: 'Please try again later',
        life: 3000,
      })
    }
  }
};
</script>

<template>

  <div class="flex flex-col gap-1">
    <Form v-slot="$form" :initialValues :resolver @submit="onSubmit" :validateOnBlur="true"
      class="flex flex-col gap-4 w-full sm:w-56">
      <div class="flex flex-col gap-1">
        <InputText name="email" type="email" placeholder="Email" fluid />
        <Message v-if="$form.username?.invalid" severity="error" size="small" variant="simple">{{
          $form.email.error?.message }}</Message>
      </div>
      <Button type="submit" severity="secondary" label="Submit" />
    </Form>
  </div>
</template>
